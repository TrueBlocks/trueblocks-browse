import React from "react";
import { types, messages } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns } from "./PortfolioTable";
import { View, FormTable, DataTable, GroupDefinition } from "@components";
import { useAppState, ViewStateProvider } from "@state";
import { Page } from "@hooks";
import { SetLast } from "@gocode/app/App";
import { EventsEmit } from "@runtime";

export function PortfolioView() {
  const { home, fetchHome } = useAppState();

  //   if (!address) {
  //     return <Text>Address not found</Text>;
  //   }

  const handleEnter = (page: Page) => {
    const record = page.selected - page.getOffset();
    const address = home.items[record].address;
    SetLast("route", `/history/${address}`);
    EventsEmit(messages.Message.NAVIGATE, {
      route: `/history/${address}`,
    });
  };

  const table = useReactTable({
    data: home.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <ViewStateProvider route={""} nItems={home.items?.length} fetchFn={fetchHome} onEnter={handleEnter}>
      <View>
        <FormTable data={home} definition={createHomeForm(table)} />
      </View>
    </ViewStateProvider>
  );
}

type theInstance = InstanceType<typeof types.PortfolioContainer>;
function createHomeForm(table: any): GroupDefinition<theInstance>[] {
  return [
    {
      title: "Open Monitors",
      fields: [],
      components: [
        {
          component: <DataTable<types.HistoryContainer> table={table} loading={false} />,
        },
      ],
    },
    {
      title: "Data 1",
      colSpan: 3,
      fields: [
        { label: "nMonitors", type: "int", accessor: "nMonitors" },
        { label: "nNames", type: "int", accessor: "nNames" },
        { label: "nAbis", type: "int", accessor: "nAbis" },
      ],
    },
    {
      title: "Data 2",
      colSpan: 3,
      fields: [
        { label: "nIndexes", type: "int", accessor: "nIndexes" },
        { label: "nManifests", type: "int", accessor: "nManifests" },
        { label: "nCaches", type: "int", accessor: "nCaches" },
      ],
    },
  ];
}
