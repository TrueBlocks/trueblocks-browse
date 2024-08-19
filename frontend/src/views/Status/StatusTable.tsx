import React from "react";
import { types } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";

const columnHelper = createColumnHelper<types.CacheItem>();

export const tableColumns: CustomColumnDef<types.CacheItem, any>[] = [
  columnHelper.accessor("type", {
    header: () => "Type",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
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
    cell: (info) => <Formatter type="date" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("path", {
    header: () => "Path",
    cell: (info) => <Formatter type="path" value={info.renderValue()} />,
    meta: { className: "wide cell" },
  }),
];
