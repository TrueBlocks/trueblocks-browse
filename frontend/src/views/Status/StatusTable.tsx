import React from "react";
import { types } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef } from "@components";

const cacheColumnHelper = createColumnHelper<types.CacheItem>();

export const cacheColumns: CustomColumnDef<types.CacheItem, any>[] = [
  cacheColumnHelper.accessor("type", {
    header: () => "Type",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell" },
  }),
  cacheColumnHelper.accessor("nFolders", {
    header: () => "nFolders",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  cacheColumnHelper.accessor("nFiles", {
    header: () => "nFiles",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  cacheColumnHelper.accessor("sizeInBytes", {
    header: () => "SizeInBytes",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  cacheColumnHelper.accessor("lastCached", {
    header: () => "LastCached",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell" },
  }),
  cacheColumnHelper.accessor("path", {
    header: () => "Path",
    cell: (info) => info.renderValue(),
    meta: { className: "wide cell" },
  }),
];
