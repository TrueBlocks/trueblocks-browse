import React from "react";
import { types } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";

const columnHelper = createColumnHelper<types.Monitor>();

export const tableColumns: CustomColumnDef<types.Monitor, any>[] = [
  columnHelper.accessor("address", {
    header: () => "Name/Address",
    cell: (info) => {
      return <Formatter type="address-editor" value={info.renderValue()} />;
    },
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("nRecords", {
    header: () => "Record Count",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("fileSize", {
    header: () => "File Size",
    cell: (info) => <Formatter type="bytes" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("lastScanned", {
    header: () => "Last Scanned",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("isEmpty", {
    header: () => "isEmpty",
    cell: (info) => <Formatter type="check" value={info.renderValue()} />,
    meta: { className: "small centered cell" },
  }),
  columnHelper.accessor("isStaged", {
    header: () => "isStaged",
    cell: (info) => <Formatter type="check" value={info.renderValue()} />,
    meta: { className: "small centered cell" },
  }),
  columnHelper.accessor("deleted", {
    header: () => "Deleted",
    cell: (info) => <Formatter type="check" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
];
