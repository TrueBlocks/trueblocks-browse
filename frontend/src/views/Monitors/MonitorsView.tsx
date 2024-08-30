import React, { useState } from "react";
import { types, messages, base } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns } from "./MonitorsTable";
import { View, FormTable, DataTable, GroupDefinition } from "@components";
import { useAppState, ViewStateProvider } from "@state";
import { SetLast } from "@gocode/app/App";
import { EventsEmit } from "@runtime";
import { Page } from "@hooks";

export function MonitorsView() {
  const { monitors, fetchMonitors } = useAppState();

  const handleEnter = (page: Page) => {
    const record = page.selected - page.getOffset();
    const address = monitors.items[record].address;
    SetLast("route", `/history/${address}`);
    EventsEmit(messages.Message.NAVIGATE, {
      route: `/history/${address}`,
    });
  };

  const table = useReactTable({
    data: monitors.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <ViewStateProvider route="monitors" nItems={monitors.nItems} fetchFn={fetchMonitors} onEnter={handleEnter}>
      <View>
        <FormTable data={monitors} definition={createMonitorForm(table)} />
      </View>
    </ViewStateProvider>
  );
}

type theInstance = InstanceType<typeof types.MonitorContainer>;
function createMonitorForm(table: any): GroupDefinition<theInstance>[] {
  return [
    {
      title: "Monitor Data",
      colSpan: 6,
      fields: [
        { label: "nMonitors", type: "int", accessor: "nItems" },
        { label: "nRecords", type: "int", accessor: "nRecords" },
        { label: "nNamed", type: "int", accessor: "nNamed" },
        { label: "fileSize", type: "bytes", accessor: "fileSize" },
      ],
    },
    {
      title: "Other",
      colSpan: 6,
      fields: [
        { label: "nEmpty", type: "int", accessor: "nEmpty" },
        { label: "nStaged", type: "int", accessor: "nStaged" },
        { label: "nDeleted", type: "int", accessor: "nDeleted" },
      ],
    },
    {
      title: "Available Monitors",
      fields: [],
      components: [
        {
          component: <DataTable<types.Monitor> table={table} loading={false} />,
        },
      ],
    },
  ];
}
