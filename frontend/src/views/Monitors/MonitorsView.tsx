import React from "react";
import { types } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns } from "./MonitorsTable";
import { View, FormTable, DataTable, GroupDefinition } from "@components";
import { useAppState } from "@state";

export function MonitorsView() {
  const { monitors } = useAppState();

  const table = useReactTable({
    data: monitors.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View route="monitors" nItems={monitors.nItems}>
      <FormTable data={monitors} definition={createMonitorForm(table)} />
    </View>
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
      ],
    },
    {
      title: "Other",
      colSpan: 6,
      fields: [
        { label: "fileSize", type: "bytes", accessor: "fileSize" },
        { label: "nDeleted", type: "int", accessor: "nDeleted" },
      ],
    },
    {
      title: "Files",
      fields: [],
      components: [
        {
          component: <DataTable<types.Monitor> table={table} loading={false} pagerName="monitors" />,
        },
      ],
    },
  ];
}
