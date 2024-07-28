import React from "react";
import { IconCircleCheck } from "@tabler/icons-react";
import { types } from "@gocode/models";
import { createColumnHelper, ColumnDef } from "@tanstack/react-table";
import { CustomMeta } from "@components";

// Find: NewViews
type CustomColumnDef<TData, TValue> = ColumnDef<TData, TValue> & {
  meta?: CustomMeta;
};

const txColumnHelper = createColumnHelper<types.TransactionEx>();

export const txColumns: CustomColumnDef<types.TransactionEx, any>[] = [
  txColumnHelper.accessor((row) => `${row.blockNumber}.${row.transactionIndex}`, {
    id: "blockTx",
    header: () => "Id",
    cell: (info) => info.getValue(),
    meta: { className: "small cell" },
  }),
  txColumnHelper.accessor("date", {
    header: () => "Date",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell" },
  }),
  txColumnHelper.accessor("fromName", {
    header: () => "From",
    cell: (info) => info.renderValue(),
    meta: { className: "wide cell" },
  }),
  txColumnHelper.accessor("toName", {
    header: () => "To",
    cell: (info) => info.renderValue(),
    meta: { className: "wide cell" },
  }),
  txColumnHelper.accessor("logCount", {
    header: () => "nEvents",
    cell: (info) => (info.renderValue() === 0 ? "-" : info.renderValue()),
    meta: { className: "medium cell" },
  }),
  txColumnHelper.accessor("ether", {
    header: () => "Ether",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell" },
  }),
  txColumnHelper.accessor("hasToken", {
    header: () => "hasToken",
    cell: (info) => (info.getValue() ? <IconCircleCheck size={20} color="white" fill="green" /> : ""),
    meta: { className: "small center cell" },
  }),
  txColumnHelper.accessor("isError", {
    header: () => "isError",
    cell: (info) => (info.getValue() ? <IconCircleCheck size={20} color="green" fill="red" /> : ""),
    meta: { className: "small center cell" },
  }),
];
