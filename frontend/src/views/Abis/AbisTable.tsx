import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter, DeleteButton } from "@components";
import { types } from "@gocode/models";

const columnHelper = createColumnHelper<types.Abi>();

export const tableColumns: CustomColumnDef<types.Abi, any>[] = [
  columnHelper.accessor("address", {
    header: () => "Name/Address",
    cell: (info) => {
      // for known abis, address is zero so we must use the name
      const { address, name } = info.row.original;
      return <Formatter type="address-editor" value={address} value2={name} />;
    },
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("isKnown", {
    header: () => "isKnown",
    cell: (info) => <Formatter type="check" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("nFunctions", {
    header: () => "nFunctions",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("nEvents", {
    header: () => "nEvents",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("isEmpty", {
    header: () => "isEmpty",
    cell: (info) => <Formatter type="error" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("fileSize", {
    header: () => "fileSize",
    cell: (info) => <Formatter type="bytes" value={info.renderValue() === 42 ? 0 : info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("hasConstructor", {
    header: () => "hasConstructor",
    cell: (info) => <Formatter type="check" value={info.renderValue() > 0} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("hasFallback", {
    header: () => "hasFallback",
    cell: (info) => <Formatter type="check" value={info.renderValue() > 0} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("lastModDate", {
    header: () => "lastModDate",
    cell: (info) => <Formatter type="date" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("address", {
    header: () => "Crud Buttons",
    cell: (info) => {
      const { address, isKnown } = info.row.original;
      return isKnown ? <></> : <DeleteButton value={address} isDeleted={false} />;
    },
    meta: { className: "small center cell" },
  }),
];
