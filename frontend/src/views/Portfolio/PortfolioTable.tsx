import React from "react";
import { types } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";

const columnHelper = createColumnHelper<types.HistoryContainer>();

export const tableColumns: CustomColumnDef<types.HistoryContainer, any>[] = [
  columnHelper.accessor("address", {
    header: () => "Address",
    cell: (info) => <Formatter type="address-editor" value={info.renderValue()} />,
    meta: { className: "large cell" },
  }),
  columnHelper.accessor("balance", {
    header: () => "Balance",
    cell: (info) => <Formatter type="ether" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("nItems", {
    header: () => "nItems",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("nLogs", {
    header: () => "nLogs",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("nTokens", {
    header: () => "nTokens",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("nErrors", {
    header: () => "nErrors",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
];
