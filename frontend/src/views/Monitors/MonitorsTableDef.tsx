// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter, CrudButton } from "@components";
import { types } from "@gocode/models";
// EXISTING_CODE

const columnHelper = createColumnHelper<types.Monitor>();

export const MonitorsTableDef: CustomColumnDef<types.Monitor, any>[] = [
  // EXISTING_CODE
  columnHelper.accessor("address", {
    header: () => "Name/Address",
    cell: (info) => {
      return <Formatter type="address-editor" value={info.renderValue()} />;
    },
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
  columnHelper.accessor("isEmpty", {
    header: () => "isEmpty",
    cell: (info) => <Formatter type="check" value={info.renderValue()} />,
    meta: { className: "small centered cell" },
  }),
  columnHelper.accessor("isStaged", {
    header: () => "isStaged",
    cell: (info) => <Formatter type="check" value={info.renderValue()} />,
    meta: { className: "small centered cell" },
  }),
  columnHelper.accessor("deleted", {
    header: () => "Deleted",
    cell: (info) => <Formatter type="check" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("deleted", {
    header: () => "Buttons",
    cell: (info) => {
      const { deleted, address } = info.row.original;
      return <CrudButton value={address} withEdit isDeleted={deleted} />;
    },
    meta: { className: "small center cell" },
  }),
  // EXISTING_CODE
];

// EXISTING_CODE
// EXISTING_CODE
