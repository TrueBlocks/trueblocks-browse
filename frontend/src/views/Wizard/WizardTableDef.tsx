import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";
import { types } from "@gocode/models";

const columnHelper = createColumnHelper<types.WizError>();

export const WizardTableDef: CustomColumnDef<types.WizError, any>[] = [
  columnHelper.accessor("count", {
    header: () => "count",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("error", {
    header: () => "error",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
];
