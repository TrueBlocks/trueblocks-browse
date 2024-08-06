import React from "react";
import { IconCircleCheck } from "@tabler/icons-react";
import { types } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef } from "@components";

const transactionColumnHelper = createColumnHelper<types.Transaction>();

export const transactionColumns: CustomColumnDef<types.Transaction, any>[] = [
  transactionColumnHelper.accessor((row) => `${row.blockNumber}.${row.transactionIndex}`, {
    id: "blockTx",
    header: () => "Id",
    cell: (info) => info.getValue(),
    meta: { className: "small cell" },
  }),
  transactionColumnHelper.accessor("from", {
    header: () => "From",
    cell: (info) => info.renderValue(),
    meta: { className: "wide cell" },
  }),
  transactionColumnHelper.accessor("to", {
    header: () => "To",
    cell: (info) => info.renderValue(),
    meta: { className: "wide cell" },
  }),
  // transactionColumnHelper.accessor("date", {
  //   header: () => "Date",
  //   cell: (info) => info.renderValue(),
  //   meta: { className: "medium cell" },
  // }),
  // transactionColumnHelper.accessor("fromName", {
  //   header: () => "From",
  //   cell: (info) => info.renderValue(),
  //   meta: { className: "wide cell" },
  // }),
  // transactionColumnHelper.accessor("toName", {
  //   header: () => "To",
  //   cell: (info) => info.renderValue(),
  //   meta: { className: "wide cell" },
  // }),
  // transactionColumnHelper.accessor("logCount", {
  //   header: () => "nEvents",
  //   cell: (info) => (info.renderValue() === 0 ? "-" : info.renderValue()),
  //   meta: { className: "medium cell" },
  // }),
  // transactionColumnHelper.accessor("ether", {
  //   header: () => "Ether",
  //   cell: (info) => info.renderValue(),
  //   meta: { className: "medium cell" },
  // }),
  transactionColumnHelper.accessor("hasToken", {
    header: () => "hasToken",
    cell: (info) => (info.getValue() ? <IconCircleCheck size={20} color="white" fill="green" /> : ""),
    meta: { className: "small center cell" },
  }),
  transactionColumnHelper.accessor("isError", {
    header: () => "isError",
    cell: (info) => (info.getValue() ? <IconCircleCheck size={20} color="green" fill="red" /> : ""),
    meta: { className: "small center cell" },
  }),
];
