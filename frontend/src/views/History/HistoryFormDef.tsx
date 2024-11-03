import { Table } from "@tanstack/react-table";
import { ExploreButton, ExportButton, DataTable, FieldGroup, GoogleButton } from "@components";
import { types, base } from "@gocode/models";

export const HistoryFormDef = (
  address: base.Address,
  table: Table<types.Transaction>
): FieldGroup<types.HistoryContainer>[] => {
  return [
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
        <ExploreButton key={"explore"} value={address} />,
        <GoogleButton key={"google"} value={address} />,
        <ExportButton key={"export"} value={address} />,
      ],
    },
    {
      label: "Transaction History",
      collapsable: false,
      components: [<DataTable<types.Transaction> key={"dataTable"} table={table} loading={false} />],
    },
  ];
};
