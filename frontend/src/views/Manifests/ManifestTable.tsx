import React from "react";
import { types } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";

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
    cell: (info) => <Formatter type="bytes" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  chunkColumnHelper.accessor("indexHash", {
    header: () => "IndexHash",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  chunkColumnHelper.accessor("indexSize", {
    header: () => "IndexSize",
    cell: (info) => <Formatter type="bytes" value={info.renderValue()} />,
    meta: { className: "wide cell" },
  }),
];
