import React from "react";
import { types, messages } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";
import { AddressPopup, NamePopup } from "@components";

const columnHelper = createColumnHelper<types.Monitor>();

export const tableColumns: CustomColumnDef<types.Monitor, any>[] = [
  columnHelper.accessor("address", {
    header: () => "Address",
    cell: (info) => <Formatter type="address-address-only" value={info.renderValue()} />,
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("name", {
    header: () => "Name",
    cell: (info) => {
      const { address, name } = info.row.original;
      return <Formatter type="address-name-only" value={info.renderValue()} />;
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
  columnHelper.accessor("deleted", {
    header: () => "Deleted",
    cell: (info) => <Formatter type="check" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
];
