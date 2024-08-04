import React from "react";
import { types } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";

const indexColumnHelper = createColumnHelper<types.ChunkStats>();

// Find: NewViews
export const indexColumns: CustomColumnDef<types.ChunkStats, any>[] = [
  indexColumnHelper.accessor("range", {
    header: () => "range",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell" },
  }),
  indexColumnHelper.accessor("nBlocks", {
    header: () => "nBlocks",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  indexColumnHelper.accessor("nAddrs", {
    header: () => "nAddrs",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  indexColumnHelper.accessor("nApps", {
    header: () => "nApps",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  indexColumnHelper.accessor("chunkSz", {
    header: () => "chunkSz",
    cell: (info) => <Formatter type="bytes" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  indexColumnHelper.accessor("nBlooms", {
    header: () => "nBlooms",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  indexColumnHelper.accessor("bloomSz", {
    header: () => "bloomSz",
    cell: (info) => <Formatter type="bytes" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  indexColumnHelper.accessor("addrsPerBlock", {
    header: () => "addrsPerBlock",
    cell: (info) => <Formatter type="float" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  indexColumnHelper.accessor("appsPerAddr", {
    header: () => "appsPerAddr",
    cell: (info) => <Formatter type="float" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  indexColumnHelper.accessor("appsPerBlock", {
    header: () => "appsPerBlock",
    cell: (info) => <Formatter type="float" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
];
