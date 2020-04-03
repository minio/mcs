// This file is part of MinIO Console Server
// Copyright (c) 2019 MinIO, Inc.
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

import React, { useState, useEffect } from "react";
import { createStyles, Theme, withStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import Typography from "@material-ui/core/Typography";
import TextField from "@material-ui/core/TextField";
import InputAdornment from "@material-ui/core/InputAdornment";
import SearchIcon from "@material-ui/icons/Search";
import {Button, IconButton, LinearProgress, TableFooter, TablePagination} from "@material-ui/core";
import Paper from "@material-ui/core/Paper";
import Table from "@material-ui/core/Table";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import TableBody from "@material-ui/core/TableBody";
import Checkbox from "@material-ui/core/Checkbox";
import ViewIcon from "@material-ui/icons/Visibility";
import DeleteIcon from "@material-ui/icons/Delete";
import {CreateIcon} from "../../../icons";
import api from "../../../common/api";
import {MinTablePaginationActions} from "../../../common/MinTablePaginationActions";
import {GroupsList} from "./types";
import {groupsSort, usersSort} from "../../../utils/sortFunctions";
import {UsersList} from "../Users/types";
import AddGroup from "../Groups/AddGroup";

interface IGroupsProps {
    classes: any;
    openGroupModal: any;
}

const styles = (theme: Theme) =>
    createStyles({
        seeMore: {
            marginTop: theme.spacing(3)
        },
        paper: {
            // padding: theme.spacing(2),
            display: "flex",
            overflow: "auto",
            flexDirection: "column"
        },
        addSideBar: {
            width: "320px",
            padding: "20px"
        },
        errorBlock: {
            color: "red"
        },
        tableToolbar: {
            paddingLeft: theme.spacing(2),
            paddingRight: theme.spacing(0)
        },
        wrapCell: {
            maxWidth: "200px",
            whiteSpace: "normal",
            wordWrap: "break-word"
        },
        minTableHeader: {
            color: "#393939",
            "& tr": {
                "& th": {
                    fontWeight:'bold'
                }
            }
        },
        actionsTray: {
            textAlign: "right",
            "& button": {
                marginLeft: 10,
            }
        },
        searchField: {
            background: "#FFFFFF",
            padding: 12,
            borderRadius: 5,
            boxShadow: "0px 3px 6px #00000012",
        }
    });

const Groups = ({
    classes,
    }: IGroupsProps) => {

    const [addGroupOpen, setGroupOpen] = useState<boolean>(false);
    const [selectedGroup, setSelectedGroup] = useState<any>(null);
    const [deleteOpen, setDeleteOpen] = useState<boolean>(false);
    const [loading, isLoading] = useState<boolean>(false);
    const [records, setRecords] = useState<any[]>([]);
    const [totalRecords, setTotalRecords] = useState<number>(0);
    const [rowsPerPage, setRowsPerPage] = useState<number>(10);
    const [page, setPage] = useState<number>(0);
    const [error, setError] = useState<string>("");

    const handleChangePage = (event: unknown, newPage: number) => {
        setPage(newPage);
    };

    const handleChangeRowsPerPage = (
        event: React.ChangeEvent<HTMLInputElement>
    ) => {
        const rPP = parseInt(event.target.value, 10);
        setPage(0);
        setRowsPerPage(rPP);
    };

    useEffect(() => {
        isLoading(true);
    }, []);

    useEffect(() => {
        isLoading(true);
    }, [page, rowsPerPage]);

    useEffect(() => {
        if(loading) {
            fetchRecords();
        }
    }, [loading]);

    const fetchRecords = () => {
        const offset = page * rowsPerPage;
        api
            .invoke("GET", `/api/v1/groups?offset=${offset}&limit=${rowsPerPage}`)
            .then((res: GroupsList) => {
                setRecords(res.groups.sort(groupsSort));
                setTotalRecords(res.total);
                setError("");
                isLoading(false);

                // if we get 0 results, and page > 0 , go down 1 page
                if ((!res.groups || res.groups.length === 0) && page > 0) {
                    const newPage = page - 1;
                    setPage(newPage);
                }
            })
            .catch(err => {
                setError(err);
                isLoading(false);
            });
    };



    const closeAddModalAndRefresh = () => {
        setGroupOpen(false);
        isLoading(true);
    };

    const closeDeleteModalAndRefresh = (refresh: boolean) => {
        setDeleteOpen(false);

        if (refresh) {
            isLoading(true);
        }
    };

    return (<React.Fragment>
        <AddGroup
            open={addGroupOpen}
            selectedGroup={selectedGroup}
            closeModalAndRefresh={closeAddModalAndRefresh}
        />
        <Grid container>
            <Grid item xs={12}>
                <Typography variant="h6">Groups</Typography>
            </Grid>
            <Grid item xs={12}>
                <br />
            </Grid>
            <Grid item xs={12} className={classes.actionsTray}>
                <TextField
                    placeholder="Search Groups"
                    className={classes.searchField}
                    id="search-resource"
                    label=""
                    InputProps={{
                        disableUnderline: true,
                        startAdornment: (
                            <InputAdornment position="start">
                                <SearchIcon />
                            </InputAdornment>
                        ),
                    }}
                />
                <Button
                    variant="contained"
                    color="primary"
                    startIcon={ <CreateIcon /> }
                    onClick={() => {
                        setSelectedGroup(null);
                        setGroupOpen(true);
                    }}
                >
                    Create Group
                </Button>
            </Grid>

            <Grid item xs={12}>
                <br />
            </Grid>
            <Grid item xs={12}>
                <Paper className={classes.paper}>
                    {loading && <LinearProgress />}
                    {records != null && records.length > 0 ? (
                        <Table size="medium">
                            <TableHead className={classes.minTableHeader}>
                                <TableRow>
                                    <TableCell>Select</TableCell>
                                    <TableCell>Name</TableCell>
                                    <TableCell align="right">Actions</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {records.map(group => (
                                    <TableRow key={`user-${group}`}>
                                        <TableCell padding="checkbox">
                                            <Checkbox
                                                value="secondary"
                                                color="primary"
                                                inputProps={{ 'aria-label': 'secondary checkbox' }}
                                            />
                                        </TableCell>
                                        <TableCell className={classes.wrapCell}>
                                            {group}
                                        </TableCell>
                                        <TableCell align="right">
                                            <IconButton
                                                aria-label="view"
                                                onClick={() => {
                                                    setGroupOpen(true);
                                                    setSelectedGroup(group);
                                                }}
                                            >
                                                <ViewIcon />
                                            </IconButton>
                                            <IconButton
                                                aria-label="delete"
                                                onClick={() => {
                                                    setDeleteOpen(true);
                                                    setSelectedGroup(group);
                                                }}
                                            >
                                                <DeleteIcon />
                                            </IconButton>
                                        </TableCell>
                                    </TableRow>
                                ))}
                            </TableBody>
                            <TableFooter>
                                <TableRow>
                                    <TablePagination
                                        rowsPerPageOptions={[5, 10, 25]}
                                        colSpan={3}
                                        count={totalRecords}
                                        rowsPerPage={rowsPerPage}
                                        page={page}
                                        SelectProps={{
                                            inputProps: { "aria-label": "rows per page" },
                                            native: true
                                        }}
                                        onChangePage={handleChangePage}
                                        onChangeRowsPerPage={handleChangeRowsPerPage}
                                        ActionsComponent={MinTablePaginationActions}
                                    />
                                </TableRow>
                            </TableFooter>
                        </Table>
                    ) : (
                        <div>No Groups Available</div>
                    )}
                </Paper>
            </Grid>
        </Grid>
    </React.Fragment>)
};

export default withStyles(styles)(Groups);
