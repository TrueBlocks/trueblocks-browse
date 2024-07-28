import React from "react";
import { IconCircleCheck } from "@tabler/icons-react";
import { types } from "@gocode/models";
import { createColumnHelper, ColumnDef } from "@tanstack/react-table";
import { CustomMeta } from "@components";
import { NameTags } from "./NameTag";

type CustomColumnDef<TData, TValue> = ColumnDef<TData, TValue> & {
  meta?: CustomMeta;
};

const nameColumnHelper = createColumnHelper<types.NameEx>();

export const nameColumns: CustomColumnDef<types.NameEx, any>[] = [
  nameColumnHelper.accessor("type", {
    header: () => "Type",
    cell: (row) => <NameTags name={row.row.original} />,
    meta: { className: "small cell" },
  }),
  nameColumnHelper.accessor("tags", {
    header: () => "Tags",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell" },
  }),
  nameColumnHelper.accessor("address", {
    header: () => "Address",
    cell: (info) => info.renderValue(),
    meta: { className: "wide cell" },
  }),
  nameColumnHelper.accessor("name", {
    header: () => "Name",
    cell: (info) => info.renderValue(),
    meta: { className: "wide cell" },
  }),
  nameColumnHelper.accessor("symbol", {
    header: () => "Symbol",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  nameColumnHelper.accessor("decimals", {
    header: () => "Decimals",
    cell: (info) => (info.getValue() === 0 ? "-" : info.getValue()),
    meta: { className: "small center cell" },
  }),
  nameColumnHelper.accessor("isContract", {
    header: () => "isContract",
    cell: (info) => (info.getValue() ? <IconCircleCheck size={20} color="white" fill="green" /> : ""),
    meta: { className: "small center cell" },
  }),
  nameColumnHelper.accessor("isErc20", {
    header: () => "isErc20",
    cell: (info) => (info.getValue() ? <IconCircleCheck size={20} color="white" fill="green" /> : ""),
    meta: { className: "small center cell" },
  }),
  nameColumnHelper.accessor("isErc721", {
    header: () => "isErc721",
    cell: (info) => (info.getValue() ? <IconCircleCheck size={20} color="white" fill="green" /> : ""),
    meta: { className: "small center cell" },
  }),
];
