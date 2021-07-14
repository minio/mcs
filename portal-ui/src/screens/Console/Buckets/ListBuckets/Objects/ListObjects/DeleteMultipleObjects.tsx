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

import React, { useState } from "react";
import { connect } from "react-redux";
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  LinearProgress,
} from "@material-ui/core";
import { setErrorSnackMessage } from "../../../../../../actions";
import api from "../../../../../../common/api";

interface IDeleteObjectProps {
  closeDeleteModalAndRefresh: (refresh: boolean) => void;
  deleteOpen: boolean;
  selectedObjects: string[];
  selectedBucket: string;
  setErrorSnackMessage: typeof setErrorSnackMessage;
}

const DeleteObject = ({
  closeDeleteModalAndRefresh,
  deleteOpen,
  selectedBucket,
  selectedObjects,
  setErrorSnackMessage,
}: IDeleteObjectProps) => {
  const [deleteLoading, setDeleteLoading] = useState<boolean>(false);

  const removeRecord = (i: number) => {
    if (deleteLoading) {
      return;
    }
    let recursive = false;
    if (selectedObjects[i].endsWith("/")) {
      recursive = true;
    }
    setDeleteLoading(true);
    api
      .invoke(
        "DELETE",
        `/api/v1/buckets/${selectedBucket}/objects?path=${selectedObjects[i]}&recursive=${recursive}`
      )
      .then(() => {
        setDeleteLoading(false);
        closeDeleteModalAndRefresh(true);
      })
      .catch((err) => {
        setDeleteLoading(false);
        setErrorSnackMessage(err);
      });
  };

  return (
    <Dialog
      open={deleteOpen}
      onClose={() => {
        closeDeleteModalAndRefresh(false);
      }}
      aria-labelledby="alert-dialog-title"
      aria-describedby="alert-dialog-description"
    >
      <DialogTitle id="alert-dialog-title">Delete</DialogTitle>
      <DialogContent>
        {deleteLoading && <LinearProgress />}
        <DialogContentText id="alert-dialog-description">
          Are you sure you want to delete the selected objects?{" "}
        </DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button
          onClick={() => {
            closeDeleteModalAndRefresh(false);
          }}
          color="primary"
          disabled={deleteLoading}
        >
          Cancel
        </Button>
        <Button
          onClick={() => {
            for (var i = 0; i < selectedObjects.length; i++) {
              removeRecord(i);
            }
          }}
          color="secondary"
          disabled={deleteLoading}
        >
          Delete
        </Button>
      </DialogActions>
    </Dialog>
  );
};

const mapDispatchToProps = {
  setErrorSnackMessage,
};

const connector = connect(null, mapDispatchToProps);

export default connector(DeleteObject);
