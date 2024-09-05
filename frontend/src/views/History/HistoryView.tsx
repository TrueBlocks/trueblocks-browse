import { useEffect } from "react";
import { Stack } from "@mantine/core";
import { getCoreRowModel, useReactTable, Table } from "@tanstack/react-table";
import { useParams } from "wouter";
import { ExploreButton, ExportButton, View, FormTable, DataTable, GroupDefinition } from "@components";
import { GetLastSub, CancleContexts } from "@gocode/app/App";
import { types, base } from "@gocode/models";
import { useAppState, ViewStateProvider } from "@state";
import { tableColumns } from "./HistoryTable";

export function HistoryView() {
  const { setAddress, history, fetchHistory } = useAppState();

  const aa = useParams().address;
  useEffect(() => {
    CancleContexts();
    if (aa === ":address") {
      GetLastSub("/history").then((subRoute) => {
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
    <ViewStateProvider route={"history"} nItems={history.nItems} fetchFn={fetchHistory}>
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
      title: "Transaction Data",
      colSpan: 6,
      fields: [
        { label: "address", type: "address-address-only", accessor: "address" },
        { label: "name", type: "address-name-only", accessor: "address" },
        { label: "balance", type: "ether", accessor: "balance" },
      ],
    },
    {
      title: "Data 0",
      colSpan: 4,
      fields: [
        { label: "nTransactions", type: "int", accessor: "nItems" },
        { label: "nLogs", type: "int", accessor: "nLogs" },
        { label: "nTokens", type: "int", accessor: "nTokens" },
        { label: "nErrors", type: "int", accessor: "nErrors" },
      ],
    },
    {
      title: "Buttons",
      colSpan: 2,
      fields: [],
      components: [
        {
          component: (
            <Stack>
              <ExploreButton size="sm" endpoint="address" value={address as unknown as string} />
              <ExportButton size="sm" value={address as unknown as string} />
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
