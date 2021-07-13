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

import request from "superagent";
import get from "lodash/get";
import { clearSession } from "../utils";
import { baseUrl } from "../../history";

export class API {
  invoke(method: string, url: string, data?: object) {
    return request(method, `${baseUrl}${url}`)
      .send(data)
      .then((res) => res.body)
      .catch((err) => {
        // if we get unauthorized, kick out the user
        if (err.status === 401) {
          clearSession();
          // Refresh the whole page to ensure cache is clear
          // and we dont end on an infinite loop
          window.location.href = "/login";
          return;
        }
        return this.onError(err);
      });
  }

  onError(err: any) {
    if (err.status) {
      const errMessage = get(
        err.response,
        "body.message",
        err.status.toString()
      );

      const throwMessage =
        errMessage.charAt(0).toUpperCase() + errMessage.slice(1);

      return Promise.reject(throwMessage);
    } else {
      clearSession();
      window.location.href = "/login";
    }
  }
}

const api = new API();
export default api;
