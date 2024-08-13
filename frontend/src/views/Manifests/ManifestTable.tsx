import React from "react";
import { types } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";

const columnHelper = createColumnHelper<types.ChunkRecord>();

export const tableColumns: CustomColumnDef<types.ChunkRecord, any>[] = [
  columnHelper.accessor("range", {
    header: () => "Range",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("bloomHash", {
    header: () => "BloomHash",
    cell: (info) => <Formatter type="hash" value={info.renderValue()} />,
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("indexHash", {
    header: () => "IndexHash",
    cell: (info) => <Formatter type="hash" value={info.renderValue()} />,
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("bloomSize", {
    header: () => "BloomSize",
    cell: (info) => <Formatter type="bytes" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("indexSize", {
    header: () => "IndexSize",
    cell: (info) => <Formatter type="bytes" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
];
