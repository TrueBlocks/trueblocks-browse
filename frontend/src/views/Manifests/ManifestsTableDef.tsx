// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";
import { types } from "@gocode/models";
// EXISTING_CODE

const columnHelper = createColumnHelper<types.ChunkRecord>();

export const ManifestsTableDef: CustomColumnDef<types.ChunkRecord, any>[] = [
  // EXISTING_CODE
  columnHelper.accessor("range", {
    header: () => "Range",
    cell: (info) => <Formatter type="range" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("rangeDates", {
    header: () => "First",
    cell: (info) => <Formatter type="date" value={info.renderValue().firstDate} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("rangeDates", {
    header: () => "Last",
    cell: (info) => <Formatter type="date" value={info.renderValue().lastDate} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("bloomHash", {
    header: () => "BloomHash",
    cell: (info) => <Formatter type="hash" value={info.renderValue()} />,
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("indexHash", {
    header: () => "IndexHash",
    cell: (info) => <Formatter type="hash" value={info.renderValue()} />,
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("bloomSize", {
    header: () => "BloomSize",
    cell: (info) => <Formatter type="bytes" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("indexSize", {
    header: () => "IndexSize",
    cell: (info) => <Formatter type="bytes" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  // EXISTING_CODE
];

// EXISTING_CODE
// EXISTING_CODE
