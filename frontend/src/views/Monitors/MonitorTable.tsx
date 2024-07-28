import React from "react";
import { useLocation } from "wouter";
import { types } from "@gocode/models";
import { createColumnHelper, ColumnDef } from "@tanstack/react-table";
import { CustomMeta } from "@components";

// Find: NewViews
type CustomColumnDef<TData, TValue> = ColumnDef<TData, TValue> & {
  meta?: CustomMeta;
};

const monitorColumnHelper = createColumnHelper<types.MonitorEx>();

export const monitorColumns: CustomColumnDef<types.MonitorEx, any>[] = [
  monitorColumnHelper.accessor("address", {
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
  monitorColumnHelper.accessor("name", {
    header: () => "Name",
    cell: (info) => info.renderValue(),
    meta: { className: "wide cell" },
  }),
  monitorColumnHelper.accessor("nRecords", {
    header: () => "Record Count",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell center" },
  }),
  monitorColumnHelper.accessor("fileSize", {
    header: () => "File Size",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell center" },
  }),
  monitorColumnHelper.accessor("lastScanned", {
    header: () => "Last Scanned",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell center" },
  }),
  monitorColumnHelper.accessor("deleted", {
    header: () => "Deleted",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell" },
  }),
];
