import React from "react";
import { IconCircleCheck } from "@tabler/icons-react";
import { app } from "@gocode/models";
import { createColumnHelper, ColumnDef } from "@tanstack/react-table";
import { CustomMeta } from "../CustomMeta";
import { NameTags } from "./NameTag";

type CustomColumnDef<TData, TValue> = ColumnDef<TData, TValue> & {
  meta?: CustomMeta;
};

const nameColumnHelper = createColumnHelper<app.NameEx>();

export const nameColumns: CustomColumnDef<app.NameEx, any>[] = [
  nameColumnHelper.accessor("type", {
    header: () => "Type",
    cell: (row) => <NameTags name={row.row.original} />,
    meta: { className: "small" },
  }),
  nameColumnHelper.accessor("tags", {
    header: () => "Tags",
    cell: (info) => info.renderValue(),
    meta: { className: "medium-small" },
  }),
  nameColumnHelper.accessor("address", {
    header: () => "Address",
    cell: (info) => info.renderValue(),
    meta: { className: "wide" },
  }),
  nameColumnHelper.accessor("name", {
    header: () => "Name",
    cell: (info) => info.renderValue(),
    meta: { className: "wide" },
  }),
  nameColumnHelper.accessor("symbol", {
    header: () => "Symbol",
    cell: (info) => info.renderValue(),
    meta: { className: "small" },
  }),
  nameColumnHelper.accessor("decimals", {
    header: () => "Decimals",
    cell: (info) => info.renderValue(),
    meta: { className: "small" },
  }),
  nameColumnHelper.accessor("isContract", {
    header: () => "isContract",
    cell: (info) => (info.getValue() ? <IconCircleCheck size={20} color="white" fill="green" /> : ""),
    meta: { className: "small-centered" },
  }),
  nameColumnHelper.accessor("isErc20", {
    header: () => "isErc20",
    cell: (info) => (info.getValue() ? <IconCircleCheck size={20} color="white" fill="green" /> : ""),
    meta: { className: "small-centered" },
  }),
  nameColumnHelper.accessor("isErc721", {
    header: () => "isErc721",
    cell: (info) => (info.getValue() ? <IconCircleCheck size={20} color="white" fill="green" /> : ""),
    meta: { className: "small-centered" },
  }),
];
