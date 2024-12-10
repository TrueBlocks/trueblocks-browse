// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { Table } from "@tanstack/react-table";
import { types } from "../../../wailsjs/go/models";
import { FieldGroup } from "../../components";
// EXISTING_CODE

export const BalancesFormDef = (table: Table<types.Transaction>): FieldGroup<types.BalanceContainer>[] => {
  // EXISTING_CODE
  // EXISTING_CODE
  return [
    // EXISTING_CODE
    // EXISTING_CODE
  ];
};

// EXISTING_CODE
/*
// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";
import { types } from "@gocode/models";

const columnHelper = createColumnHelper<types.Transaction>();

export const HistoryTableDef: CustomColumnDef<types.Transaction, any>[] = [
    columnHelper.accessor("blockNumber", {
    header: () => "TxId",
    cell: (info) => {
      const { blockNumber, transactionIndex, hash } = info.row.original;
      return <Formatter type="appearance" value={`${blockNumber}.${transactionIndex}`} value2={hash} />;
    },
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("timestamp", {
    id: "Timestamp",
    cell: (info) => <Formatter type="timestamp" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("from", {
    header: () => "From",
    cell: (info) => <Formatter type="address-editor" value={info.renderValue()} />,
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("to", {
    header: () => "To",
    cell: (info) => <Formatter type="address-editor" value={info.renderValue()} />,
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("value", {
    header: () => "Ether",
    cell: (info) => {
      const { value, isError } = info.row.original;
      return <Formatter type="ether" value={value} value2={isError} />;
    },
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("gasUsed", {
    header: () => "Gas",
    cell: (info) => {
      const { gasUsed, gasPrice, from } = info.row.original;
      const gasCost = gasUsed * gasPrice;
      return <Formatter type="gas" value={gasCost} value2={from} />;
    },
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("isError", {
    header: () => "isError",
    cell: (info) => <Formatter type="error" value={info.renderValue()} />,
    meta: { className: "small center cell" },
  }),
  columnHelper.accessor("input", {
    header: () => "Function",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium center cell" },
  }),
  ];

*/
// EXISTING_CODE
