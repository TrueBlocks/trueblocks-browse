// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";
import { types } from "@gocode/models";
// EXISTING_CODE

const columnHelper = createColumnHelper<types.WizError>();

export const WizardTableDef: CustomColumnDef<types.WizError, any>[] = [
  // EXISTING_CODE
  columnHelper.accessor("index", {
    header: () => "index",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("state", {
    header: () => "state",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("reason", {
    header: () => "reason",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("error", {
    header: () => "error",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  // EXISTING_CODE
];

// EXISTING_CODE
// EXISTING_CODE
