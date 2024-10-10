import { useEffect } from "react";
import { Group, Stack } from "@mantine/core";
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
import { GetSessionSubVal, ModifyNoop } from "@gocode/app/App";
import { types, base } from "@gocode/models";
import { useAppState, ViewStateProvider } from "@state";
import { tableColumns } from "./HistoryTable";

export function HistoryView() {
  const { setAddress, history, fetchHistory } = useAppState();

  const aa = useParams().address;
  useEffect(() => {
    if (aa === ":address") {
      GetSessionSubVal("/history").then((subRoute) => {
        subRoute = subRoute.replace("/", "");
        setAddress(subRoute as unknown as base.Address);
      });
    } else {
      setAddress(aa as unknown as base.Address);
    }
  }, [aa]);

  const table = useReactTable({
    data: history.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <ViewStateProvider route={"history"} nItems={history.nItems} fetchFn={fetchHistory} modifyFn={ModifyNoop}>
      <View>
        <FormTable data={history} definition={CreateHistoryForm(table)} />
      </View>
    </ViewStateProvider>
  );
}

type theInstance = InstanceType<typeof types.HistoryContainer>;
function CreateHistoryForm(table: Table<types.Transaction>): GroupDefinition<theInstance>[] {
  const { address } = useAppState();
  return [
    {
      title: "DalleDress",
      colSpan: 2,
      fields: [{ label: "", type: "dalle", accessor: "address" }],
    },
    {
      title: "Transaction Data",
      colSpan: 5,
      fields: [
        { label: "address", type: "address-address-only", accessor: "address" },
        { label: "name", type: "address-name-only", accessor: "address" },
        { label: "balance", type: "ether", accessor: "balance" },
      ],
    },
    {
      title: "Transaction Data",
      colSpan: 2,
      fields: [
        { label: "nTransactions", type: "int", accessor: "nItems" },
        { label: "nLogs", type: "int", accessor: "nLogs" },
        { label: "nTokens", type: "int", accessor: "nTokens" },
        { label: "nErrors", type: "int", accessor: "nErrors" },
      ],
    },
    {
      title: "Buttons",
      colSpan: 3,
      fields: [],
      components: [
        {
          component: (
            <Stack>
              <Group>
                <ExploreButton value={address}>Explore</ExploreButton>
                <DalleButton value={address}>Dalle</DalleButton>
              </Group>
              <Group>
                <GoogleButton value={address}>Google</GoogleButton>
                <ExportButton value={address}>Export</ExportButton>
              </Group>
            </Stack>
          ),
        },
      ],
    },
    {
      title: "Transaction History",
      fields: [],
      components: [
        {
          component: <DataTable<types.Transaction> table={table} loading={false} />,
        },
      ],
    },
  ];
}
