import { useEffect } from "react";
import { getCoreRowModel, useReactTable, Table } from "@tanstack/react-table";
import { useParams } from "wouter";
import { ExploreButton, ExportButton, View, FormTable, DataTable, FieldGroup, GoogleButton } from "@components";
import { types, base } from "@gocode/models";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { tableColumns } from "./HistoryTable";

export const HistoryView = () => {
  const { modifyNoop } = useNoops();
  const { setAddress, history, fetchHistory } = useAppState();

  const address = useParams().address as unknown as base.Address;
  useEffect(() => {
    setAddress(address);
  }, [address, setAddress]);

  const table = useReactTable({
    data: history.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "history";
  return (
    <ViewStateProvider route={route} nItems={history.nItems} fetchFn={fetchHistory} modifyFn={modifyNoop}>
      <View>
        <FormTable data={history} groups={createHistoryForm(address, table)} />
      </View>
    </ViewStateProvider>
  );
};

const createHistoryForm = (
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
        { label: "address", type: "address-address-only", accessor: "address" },
        { label: "name", type: "address-name-only", accessor: "address" },
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
