import React from "react";
import { types } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";

const abiColumnHelper = createColumnHelper<types.Abi>();

// Find: NewViews
export const abiColumns: CustomColumnDef<types.Abi, any>[] = [
  abiColumnHelper.accessor("name", {
    header: () => "Name",
    cell: (info) => {
      const { address, name } = info.row.original;
      return address && address.toString() !== "0x0" ? <Formatter type="address" value={address} /> : name;
    },
    meta: { className: "wide cell" },
  }),
  abiColumnHelper.accessor("lastModDate", {
    header: () => "lastModDate",
    cell: (info) => info.renderValue(),
    meta: { className: "large cell" },
  }),
  abiColumnHelper.accessor("fileSize", {
    header: () => "fileSize",
    cell: (info) => <Formatter type="bytes" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  abiColumnHelper.accessor("isKnown", {
    header: () => "isKnown",
    cell: (info) => <Formatter type="check" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  abiColumnHelper.accessor("nFunctions", {
    header: () => "nFunctions",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  abiColumnHelper.accessor("nEvents", {
    header: () => "nEvents",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
];
