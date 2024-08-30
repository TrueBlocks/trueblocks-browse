import React from "react";
import { types } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter, ExploreButton, ExportButton, ViewButton } from "@components";
import { Group } from "@mantine/core";

const columnHelper = createColumnHelper<types.HistoryContainer>();

export const tableColumns: CustomColumnDef<types.HistoryContainer, any>[] = [
  columnHelper.accessor("address", {
    header: () => "Address",
    cell: (info) => <Formatter type="address-editor" value={info.renderValue()} />,
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("balance", {
    header: () => "Balance",
    cell: (info) => <Formatter type="ether" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("nItems", {
    header: () => "nItems",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("nLogs", {
    header: () => "nLogs",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("nErrors", {
    header: () => "nErrors",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("address", {
    header: () => " ",
    cell: (info) => (
      <Group wrap={"nowrap"}>
        <ExploreButton size="xs" noText endpoint="address" value={info.renderValue()} />
        <ViewButton size="xs" noText value={info.renderValue()} />
        <ExportButton size="xs" noText value={info.renderValue()} />
      </Group>
    ),
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("nItems", {
    header: () => "Items",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
];
