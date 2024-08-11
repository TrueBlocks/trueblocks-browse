import React from "react";
import { types } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";
import { useToEther, useDateTime } from "@hooks";

const columnHelper = createColumnHelper<types.Transaction>();

// Find: NewViews
export const tableColumns: CustomColumnDef<types.Transaction, any>[] = [
  columnHelper.accessor((row) => `${row.blockNumber}.${row.transactionIndex}`, {
    id: "blockTx",
    header: () => "Id",
    cell: (info) => info.getValue(),
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("timestamp", {
    id: "Timestamp",
    cell: (info) => useDateTime(info.getValue()),
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("from", {
    header: () => "From",
    cell: (info) => <Formatter type="address" value={info.renderValue()} />,
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("to", {
    header: () => "To",
    cell: (info) => <Formatter type="address" value={info.renderValue()} />,
    meta: { className: "wide cell" },
  }),
  // columnHelper.accessor("date", {
  //   header: () => "Date",
  //   cell: (info) => info.renderValue(),
  //   meta: { className: "medium cell" },
  // }),
  // columnHelper.accessor("logCount", {
  //   header: () => "nEvents",
  //   cell: (info) => (info.renderValue() === 0 ? "-" : info.renderValue()),
  //   meta: { className: "medium cell" },
  // }),
  columnHelper.accessor("value", {
    header: () => "Ether",
    cell: (info) => useToEther(info.renderValue() as bigint),
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("hasToken", {
    header: () => "hasToken",
    cell: (info) => <Formatter type="check" value={info.renderValue()} />,
    meta: { className: "small center cell" },
  }),
  columnHelper.accessor("isError", {
    header: () => "isError",
    cell: (info) => <Formatter type="error" value={info.renderValue()} />,
    meta: { className: "small center cell" },
  }),
];
