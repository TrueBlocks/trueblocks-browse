import React from "react";
import { types } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";

const columnHelper = createColumnHelper<types.CacheItem>();

// Find: NewViews
export const tableColumns: CustomColumnDef<types.CacheItem, any>[] = [
  columnHelper.accessor("type", {
    header: () => "Type",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("nFolders", {
    header: () => "nFolders",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("nFiles", {
    header: () => "nFiles",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("sizeInBytes", {
    header: () => "SizeInBytes",
    cell: (info) => <Formatter type="bytes" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("lastCached", {
    header: () => "LastCached",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("path", {
    header: () => "Path",
    cell: (info) => info.renderValue(),
    meta: { className: "wide cell" },
  }),
];
