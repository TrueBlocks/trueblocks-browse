// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter, CrudButton } from "@components";
import { NameTags } from "@components";
import { types } from "@gocode/models";
// EXISTING_CODE

const columnHelper = createColumnHelper<types.Name>();

export const NamesTableDef: CustomColumnDef<types.Name, any>[] = [
  // EXISTING_CODE
  columnHelper.accessor("parts", {
    header: () => "Type",
    cell: (row) => <NameTags name={row.row.original} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("tags", {
    header: () => "Tags",
    cell: (info) => <Formatter type="tag" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("address", {
    header: () => "Name/Address",
    cell: (info) => {
      return <Formatter type="address-editor" value={info.renderValue()} />;
    },
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("symbol", {
    header: () => "Symbol",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("decimals", {
    header: () => "Decimals",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
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
  columnHelper.accessor("deleted", {
    header: () => "Buttons",
    cell: (info) => {
      const { deleted, isCustom, address } = info.row.original;
      return isCustom ? <CrudButton value={address} withEdit isDeleted={deleted as boolean} /> : <></>;
    },
    meta: { className: "small center cell" },
  }),
  // EXISTING_CODE
];

// EXISTING_CODE
// EXISTING_CODE

//-------------------------------------------------------------------
// Template variables:
// class:         Name
// lower:         name
// routeLabel:    Names
// routeLower:    names
// itemName:      Name
// embedType:     .
