import React from "react";
import { useLocation } from "wouter";
import { types } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";

const columnHelper = createColumnHelper<types.Monitor>();

// Find: NewViews
export const tableColumns: CustomColumnDef<types.Monitor, any>[] = [
  columnHelper.accessor("address", {
    header: () => "Address",
    cell: ({ getValue }) => {
      const [_, setLocation] = useLocation();
      const address = getValue();
      return (
        <a
          href="#"
          onClick={(e) => {
            e.preventDefault();
            setLocation(`/history/${address}`);
          }}
        >
          {address}
        </a>
      );
    },
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("name", {
    header: () => "Name",
    cell: (info) => info.renderValue(),
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
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell" },
  }),
];
