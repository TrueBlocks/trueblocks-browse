import React from "react";
import { types } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter, FormatterProps } from "@components";

const columnHelper = createColumnHelper<types.Abi>();

export const tableColumns: CustomColumnDef<types.Abi, any>[] = [
  columnHelper.accessor("address", {
    header: () => "Name/Address",
    cell: (info) => {
      const { address, name } = info.row.original;
      return <Formatter type="address-and-name" value={address} value2={name} />;
    },
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("lastModDate", {
    header: () => "lastModDate",
    cell: (info) => <Formatter type="date" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("fileSize", {
    header: () => "fileSize",
    cell: (info) => <Formatter type="bytes" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("isKnown", {
    header: () => "isKnown",
    cell: (info) => <Formatter type="check" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("isEmpty", {
    header: () => "isEmpty",
    cell: (info) => <Formatter type="error" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("nFunctions", {
    header: () => "nFunctions",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("nEvents", {
    header: () => "nEvents",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("nFunctions", {
    header: () => "hasContructor",
    cell: (info) => <Formatter type="check" value={info.renderValue() > 0} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("nEvents", {
    header: () => "hasFallback",
    cell: (info) => <Formatter type="check" value={info.renderValue() > 0} />,
    meta: { className: "medium cell" },
  }),
];

// function AbiFormatter({ className, value, size }: Omit<FormatterProps, "type">) {
//   // 42 byte files in the abis folder means we could not find the abi
//   return <Formatter size={size} className={className} type="error" value={value === 42} />;
// }
