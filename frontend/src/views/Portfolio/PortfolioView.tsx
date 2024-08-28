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
  const { portfolio, fetchPortfolio } = useAppState();

  const handleEnter = (page: Page) => {
    const record = page.selected - page.getOffset();
    const address = portfolio.items[record].address;
    SetLast("route", `/history/${address}`);
    EventsEmit(messages.Message.NAVIGATE, {
      route: `/history/${address}`,
    });
  };

  //   if (!address) {
  //     return <Text>Address not found</Text>;
  //   }

  const table = useReactTable({
    data: portfolio.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <ViewStateProvider route={""} nItems={portfolio.myCount} fetchFn={fetchPortfolio} onEnter={handleEnter}>
      <View>
        <FormTable data={portfolio} definition={createPortfolioForm(table)} />
      </View>
    </ViewStateProvider>
  );
}

type theInstance = InstanceType<typeof types.PortfolioContainer>;
function createPortfolioForm(table: any): GroupDefinition<theInstance>[] {
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
        { label: "historySize", type: "bytes", accessor: "historySize" },
      ],
    },
  ];
}