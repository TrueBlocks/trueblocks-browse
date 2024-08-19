import React from "react";
import { types } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, NamePopup, AddressPopup, Formatter } from "@components";
import { NameTags } from "./NameTag";

const columnHelper = createColumnHelper<types.Name>();

export const tableColumns: CustomColumnDef<types.Name, any>[] = [
  columnHelper.accessor("parts", {
    header: () => "Type",
    cell: (row) => <NameTags name={row.row.original} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("tags", {
    header: () => "Tags",
    cell: (info) => info.renderValue(),
    meta: { className: "medium cell" },
  }),
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
      editor: (getValue: () => any) => <NamePopup name={getValue} onSubmit={(newValue) => console.log(newValue)} />,
    },
  }),
  columnHelper.accessor("symbol", {
    header: () => "Symbol",
    cell: (info) => info.renderValue(),
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("decimals", {
    header: () => "Decimals",
    cell: (info) => (info.getValue() === 0 ? "-" : info.getValue()),
    meta: { className: "small center cell" },
  }),
  columnHelper.accessor("isContract", {
    header: () => "isContract",
    cell: (info) => <Formatter type="check" value={info.renderValue()} />,
    meta: { className: "small center cell" },
  }),
  columnHelper.accessor("isErc20", {
    header: () => "isErc20",
    cell: (info) => <Formatter type="check" value={info.renderValue()} />,
    meta: { className: "small center cell" },
  }),
  columnHelper.accessor("isErc721", {
    header: () => "isErc721",
    cell: (info) => <Formatter type="check" value={info.renderValue()} />,
    meta: { className: "small center cell" },
  }),
];
