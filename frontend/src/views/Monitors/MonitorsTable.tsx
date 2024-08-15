import React from "react";
import { types, messages } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";
import { AddressPopup, NamePopup } from "@components";

const columnHelper = createColumnHelper<types.Monitor>();

export const tableColumns: CustomColumnDef<types.Monitor, any>[] = [
  columnHelper.accessor("address", {
    header: () => "Address",
    cell: (info) => info.renderValue(),
    meta: {
      className: "wide cell",
      editor: (getValue: () => any) => <AddressPopup address={getValue} />,
    },
  }),
  columnHelper.accessor("name", {
    header: () => "Name",
    cell: (info) => info.renderValue(),
    meta: {
      className: "wide cell",
      editor: (getValue: () => any) => (
        <NamePopup name={getValue} onSubmit={(newValue: string) => console.log(newValue)} />
      ),
    },
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
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell" },
  }),
];
