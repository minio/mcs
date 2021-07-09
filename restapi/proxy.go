// This file is part of MinIO Console Server
// Copyright (c) 2021 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package restapi

import (
	"bytes"
	"crypto/sha1"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	url2 "net/url"
	"strings"

	v2 "github.com/minio/operator/pkg/apis/minio.min.io/v2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/minio/console/cluster"

	"github.com/minio/console/pkg/auth"
)

func serveProxy(responseWriter http.ResponseWriter, req *http.Request) {
	log.Println("Proxy", req.URL.Path)
	urlParts := strings.Split(req.URL.Path, "/")

	if len(urlParts) < 5 {
		log.Println(len(urlParts))
		return
	}
	namespace := urlParts[3]
	tenantName := urlParts[4]

	// validate the tenantName

	token, err := auth.GetTokenFromRequest(req)
	if err != nil {
		log.Println(err)
		responseWriter.WriteHeader(401)
		return
	}
	claims, err := auth.SessionTokenAuthenticate(token)
	if err != nil {
		log.Println("Unable to validate the session token %s: %v", token, err)
		responseWriter.WriteHeader(401)

		return
	}

	//STSSessionToken := currToken.Value
	STSSessionToken := claims.STSSessionToken

	opClientClientSet, err := cluster.OperatorClient(STSSessionToken)
	if err != nil {
		log.Println(err)
		responseWriter.WriteHeader(404)
		return
	}
	opClient := operatorClient{
		client: opClientClientSet,
	}
	tenant, err := opClient.TenantGet(req.Context(), namespace, tenantName, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
		return
	}

	nsTenant := fmt.Sprintf("%s/%s", namespace, tenantName)

	tenantSchema := "https"
	if !tenant.TLS() {
		tenantSchema = "http"
	}

	tenantUrl := fmt.Sprintf("%s://%s.%s.svc.%s:9443", tenantSchema, tenant.ConsoleCIServiceName(), tenant.Namespace, v2.GetClusterDomain())
	// for development
	//tenantUrl = "http://localhost:9091"
	tenantUrl = "https://localhost:9443"

	h := sha1.New()
	h.Write([]byte(nsTenant))
	log.Printf("Proxying request for %s/%s", namespace, tenantName)
	tenantCookieName := fmt.Sprintf("token-%x", string(h.Sum(nil)))
	log.Println("tenantCookieName", tenantCookieName)
	tenantCookie, err := req.Cookie(tenantCookieName)
	if err != nil {
		log.Println("no cookie", err)

		// login to tenantName
		loginUrl := fmt.Sprintf("%s/api/v1/login", tenantUrl)

		// get the tenant credentials
		clientSet, err := cluster.K8sClient(STSSessionToken)
		if err != nil {
			log.Println(err)
			return
		}

		currentSecret, err := clientSet.CoreV1().Secrets(namespace).Get(req.Context(), tenant.Spec.CredsSecret.Name, metav1.GetOptions{})
		if err != nil {
			log.Println(err)
			responseWriter.WriteHeader(500)
			return
		}

		data := map[string]string{
			"accessKey": string(currentSecret.Data["accesskey"]),
			"secretKey": string(currentSecret.Data["secretkey"]),
		}
		payload, _ := json.Marshal(data)

		loginReq, err := http.NewRequest(http.MethodPost, loginUrl, bytes.NewReader(payload))
		if err != nil {
			log.Println(err)
			return
		}
		loginReq.Header.Add("Content-Type", "application/json")

		if err != nil {
			log.Println(err)
			return
		}

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}

		loginResp, err := client.Do(loginReq)
		for _, c := range loginResp.Cookies() {
			fmt.Println("resp cookie:", c.Name)
			if c.Name == "token" {
				tenantCookie = c
				c := &http.Cookie{
					Name:     tenantCookieName,
					Value:    c.Value,
					Path:     c.Path,
					Expires:  c.Expires,
					HttpOnly: c.HttpOnly,
				}

				http.SetCookie(responseWriter, c)
				break
			}
		}
		defer loginResp.Body.Close()
		b, _ := io.ReadAll(loginResp.Body)
		fmt.Println(string(b))

	}

	origin, _ := url2.Parse(tenantUrl)
	targetUrl, _ := url2.Parse(tenantUrl)

	targetUrl.Scheme = "http"
	if tenant.TLS() {
		targetUrl.Scheme = "https"
	}
	targetUrl.Host = origin.Host
	targetUrl.Path = strings.Replace(req.URL.Path, fmt.Sprintf("/api/proxy/%s/%s", namespace, tenantName), "", -1)

	fmt.Println("targetUrl", targetUrl)

	proxiedCookie := &http.Cookie{
		Name:     "token",
		Value:    tenantCookie.Value,
		Path:     tenantCookie.Path,
		Expires:  tenantCookie.Expires,
		HttpOnly: tenantCookie.HttpOnly,
	}

	proxyCookieJar, _ := cookiejar.New(nil)
	proxyCookieJar.SetCookies(targetUrl, []*http.Cookie{proxiedCookie})

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr,
		Jar: proxyCookieJar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}

	proxRequest, err := http.NewRequest(req.Method, targetUrl.String(), req.Body)

	for _, v := range req.Header.Values("Content-Type") {
		proxRequest.Header.Add("Content-Type", v)
	}

	resp, err := client.Do(proxRequest)

	for hk, hv := range resp.Header {
		fmt.Println(hk, hv)
		// allow iframing
		if hk != "X-Frame-Options" {
			for _, v := range hv {
				responseWriter.Header().Add(hk, v)
			}
		}
	}
	// Allow iframes
	responseWriter.Header().Set("X-Frame-Options", "SAMEORIGIN")
	responseWriter.Header().Set("X-XSS-Protection", "0")

	io.Copy(responseWriter, resp.Body)

}
