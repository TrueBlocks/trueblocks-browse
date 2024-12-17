// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";
import { types } from "@gocode/models";
// EXISTING_CODE

const columnHelper = createColumnHelper<types.Nothing>();

export const SessionTableDef: CustomColumnDef<types.Nothing, any>[] = [
  // EXISTING_CODE
  columnHelper.accessor("unused", {
    header: () => "Symbol",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  // EXISTING_CODE
];

// EXISTING_CODE
// EXISTING_CODE
