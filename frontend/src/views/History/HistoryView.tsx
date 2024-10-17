import { useEffect } from "react";
import { Stack } from "@mantine/core";
import { getCoreRowModel, useReactTable, Table } from "@tanstack/react-table";
import { useParams } from "wouter";
import {
  ExploreButton,
  ExportButton,
  View,
  FormTable,
  DataTable,
  GroupDefinition,
  DalleButton,
  GoogleButton,
} from "@components";
import { GetAddress, ModifyNoop } from "@gocode/app/App";
import { types, base } from "@gocode/models";
import { useAppState, ViewStateProvider } from "@state";
import { tableColumns } from "./HistoryTable";

export const HistoryView = () => {
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

  return (
    <ViewStateProvider route={"history"} nItems={history.nItems} fetchFn={fetchHistory} modifyFn={ModifyNoop}>
      <View>
        <FormTable data={history} definition={createHistoryForm(address, table)} />
      </View>
    </ViewStateProvider>
  );
};

type theInstance = InstanceType<typeof types.HistoryContainer>;
const createHistoryForm = (address: base.Address, table: Table<types.Transaction>): GroupDefinition<theInstance>[] => {
  return [
    {
      legend: "DalleDress",
      colSpan: 2,
      fields: [{ label: "", type: "dalle", accessor: "address" }],
    },
    {
      legend: "Transaction Data",
      colSpan: 5,
      fields: [
        { label: "address", type: "address-address-only", accessor: "address" },
        { label: "name", type: "address-name-only", accessor: "address" },
        { label: "balance", type: "ether", accessor: "balance" },
      ],
    },
    {
      legend: "Transaction Data",
      colSpan: 3,
      fields: [
        { label: "nTransactions", type: "int", accessor: "nItems" },
        { label: "nLogs", type: "int", accessor: "nLogs" },
        { label: "nTokens", type: "int", accessor: "nTokens" },
        { label: "nErrors", type: "int", accessor: "nErrors" },
      ],
    },
    {
      legend: "Buttons",
      colSpan: 2,
      components: [
        {
          component: (
            <Stack>
              <ExploreButton value={address}>Explore</ExploreButton>
              <GoogleButton value={address}>Google</GoogleButton>
              <ExportButton value={address}>Export</ExportButton>
            </Stack>
          ),
        },
      ],
    },
    {
      legend: "Transaction History",
      components: [
        {
          component: <DataTable<types.Transaction> table={table} loading={false} />,
        },
      ],
    },
  ];
};
