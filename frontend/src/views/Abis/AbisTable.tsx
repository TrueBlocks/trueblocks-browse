import React from "react";
import { types } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef } from "@components";

const abiColumnHelper = createColumnHelper<types.AbiFile>();

// Find: NewViews
export const abiColumns: CustomColumnDef<types.AbiFile, any>[] = [
  abiColumnHelper.accessor("name", {
    header: () => "Name",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell" },
  }),
  abiColumnHelper.accessor("lastModDate", {
    header: () => "lastModDate",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  abiColumnHelper.accessor("fileSize", {
    header: () => "fileSize",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  abiColumnHelper.accessor("path", {
    header: () => "path",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell" },
  }),
];
