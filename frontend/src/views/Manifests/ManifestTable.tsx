import React from "react";
import { types } from "@gocode/models";
import { createColumnHelper, ColumnDef } from "@tanstack/react-table";
import { CustomMeta } from "@components";

type CustomColumnDef<TData, TValue> = ColumnDef<TData, TValue> & {
  meta?: CustomMeta;
};

const chunkColumnHelper = createColumnHelper<types.ChunkRecord>();

export const chunkColumns: CustomColumnDef<types.ChunkRecord, any>[] = [
  chunkColumnHelper.accessor("range", {
    header: () => "Range",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell" },
  }),
  chunkColumnHelper.accessor("bloomHash", {
    header: () => "BloomHash",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  chunkColumnHelper.accessor("bloomSize", {
    header: () => "BloomSize",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  chunkColumnHelper.accessor("indexHash", {
    header: () => "IndexHash",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  chunkColumnHelper.accessor("indexSize", {
    header: () => "IndexSize",
    cell: (info) => info.renderValue(),
    meta: { className: "wide cell" },
  }),
];
