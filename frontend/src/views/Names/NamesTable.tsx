import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter, CrudButton } from "@components";
import { types } from "@gocode/models";
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
    header: () => "Crud Buttons",
    cell: (info) => {
      const { deleted, isCustom, address } = info.row.original;
      return isCustom ? <CrudButton value={address} withEdit isDeleted={deleted as boolean} /> : <></>;
    },
    meta: { className: "small center cell" },
  }),
];
