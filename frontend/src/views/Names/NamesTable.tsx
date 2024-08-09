import React from "react";
import { IconCircleCheck } from "@tabler/icons-react";
import { types } from "@gocode/models";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, AddressNameEditor, AddressNameViewer } from "@components";
import { NameTags } from "./NameTag";

const columnHelper = createColumnHelper<types.Name>();

// Find: NewViews
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
      editor: (getValue: () => any) => <AddressNameViewer address={getValue} />,
    },
  }),
  columnHelper.accessor("name", {
    header: () => "Name",
    cell: (info) => info.renderValue(),
    meta: {
      className: "wide cell",
      editor: (getValue: () => any) => (
        <AddressNameEditor value={getValue} onSubmit={(newValue) => console.log(newValue)} />
      ),
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
    cell: (info) => (info.getValue() ? <IconCircleCheck size={20} color="white" fill="green" /> : ""),
    meta: { className: "small center cell" },
  }),
  columnHelper.accessor("isErc20", {
    header: () => "isErc20",
    cell: (info) => (info.getValue() ? <IconCircleCheck size={20} color="white" fill="green" /> : ""),
    meta: { className: "small center cell" },
  }),
  columnHelper.accessor("isErc721", {
    header: () => "isErc721",
    cell: (info) => (info.getValue() ? <IconCircleCheck size={20} color="white" fill="green" /> : ""),
    meta: { className: "small center cell" },
  }),
];
