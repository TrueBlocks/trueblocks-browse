// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { Table } from "@tanstack/react-table";
import { ExploreButton, ExportButton, DataTable, FieldGroup, GoogleButton } from "@components";
import { types } from "@gocode/models";
import { useAppState } from "../../state";
// EXISTING_CODE

export const HistoryFormDef = (table: Table<types.Transaction>): FieldGroup<types.HistoryContainer>[] => {
  // EXISTING_CODE
  const { info } = useAppState();
  // EXISTING_CODE
  return [
    // EXISTING_CODE
    {
      label: "DalleDress",
      colSpan: 2,
      fields: [{ label: "", type: "dalle", accessor: "address" }],
    },
    {
      label: "Transaction Data",
      colSpan: 7,
      fields: [
        { label: "name", type: "address-name-only", accessor: "address" },
        { label: "address", type: "address-address-only", accessor: "address" },
        { label: "balance", type: "ether", accessor: "balance" },
      ],
    },
    {
      label: "Transaction Data",
      colSpan: 3,
      fields: [
        { label: "nTransactions", type: "int", accessor: "nItems" },
        { label: "nLogs", type: "int", accessor: "nLogs" },
        { label: "nTokens", type: "int", accessor: "nTokens" },
        { label: "nErrors", type: "int", accessor: "nErrors" },
      ],
    },
    {
      label: "Buttons",
      buttons: [
        <ExploreButton key={"explore"} value={info.address} />,
        <GoogleButton key={"google"} value={info.address} />,
        <ExportButton key={"export"} value={info.address} />,
      ],
    },
    {
      label: "Transactions",
      collapsable: false,
      components: [<DataTable<types.Transaction> key={"dataTable"} table={table} loading={false} />],
    },
    // EXISTING_CODE
  ];
};

// EXISTING_CODE
// EXISTING_CODE
