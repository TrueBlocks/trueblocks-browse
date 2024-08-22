import React from "react";
import { types } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter, AddressPopup } from "@components";
import { useToEther, useDateTime } from "@hooks";

const columnHelper = createColumnHelper<types.Transaction>();

export const tableColumns: CustomColumnDef<types.Transaction, any>[] = [
  columnHelper.accessor((row) => `${row.blockNumber}.${row.transactionIndex}`, {
    id: "blockTx",
    header: () => "Id",
    cell: (info) => <Formatter type="appearance" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("timestamp", {
    id: "Timestamp",
    cell: (info) => <Formatter type="timestamp" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("from", {
    header: () => "From",
    cell: (info) => <Formatter type="address-and-name" value={info.renderValue()} />,
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("to", {
    header: () => "To",
    cell: (info) => <Formatter type="address-and-name" value={info.renderValue()} />,
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("value", {
    header: () => "Ether",
    cell: (info) => <Formatter type="ether" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("isError", {
    header: () => "isError",
    cell: (info) => <Formatter type="error" value={info.renderValue()} />,
    meta: { className: "small center cell" },
  }),
  columnHelper.accessor("input", {
    header: () => "Function",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium center cell" },
  }),
];
