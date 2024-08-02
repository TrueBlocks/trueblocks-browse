import React from "react";
import { types } from "@gocode/models";
import { createColumnHelper, ColumnDef } from "@tanstack/react-table";
import { CustomMeta } from "@components";

type CustomColumnDef<TData, TValue> = ColumnDef<TData, TValue> & {
  meta?: CustomMeta;
};

const indexColumnHelper = createColumnHelper<types.ChunkStats>();

// Find: NewViews
export const indexColumns: CustomColumnDef<types.ChunkStats, any>[] = [
  indexColumnHelper.accessor("range", {
    header: () => "range",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell" },
  }),
  indexColumnHelper.accessor("rangeEnd", {
    header: () => "rangeEnd",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell" },
  }),
  indexColumnHelper.accessor("addrsPerBlock", {
    header: () => "addrsPerBlock",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  indexColumnHelper.accessor("appsPerAddr", {
    header: () => "appsPerAddr",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  indexColumnHelper.accessor("appsPerBlock", {
    header: () => "appsPerBlock",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  indexColumnHelper.accessor("bloomSz", {
    header: () => "bloomSz",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  indexColumnHelper.accessor("chunkSz", {
    header: () => "chunkSz",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  indexColumnHelper.accessor("nAddrs", {
    header: () => "nAddrs",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  indexColumnHelper.accessor("nApps", {
    header: () => "nApps",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  indexColumnHelper.accessor("nBlocks", {
    header: () => "nBlocks",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  indexColumnHelper.accessor("nBlooms", {
    header: () => "nBlooms",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  indexColumnHelper.accessor("ratio", {
    header: () => "addrsPerBlock",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  // indexColumnHelper.accessor("recWid", {
  //   header: () => "recWid",
  //   cell: (info) => info.renderValue(),
  //   meta: { className: "small cell" },
  // }),
];
