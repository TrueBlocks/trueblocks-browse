import React from "react";
import { types } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns } from "./StatusTable";
import { View, FormTable, DataTable, GroupDefinition } from "@components";
import { useAppState } from "@state";

export function StatusView() {
  const { status } = useAppState();

  const table = useReactTable({
    data: status.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <FormTable data={status} definition={createStatusForm(table)} />
    </View>
  );
}

type theInstance = InstanceType<typeof types.StatusContainer>;
function createStatusForm(table: any): GroupDefinition<theInstance>[] {
  return [
    {
      title: "System Data",
      colSpan: 7,
      fields: [
        { label: "trueblocks", accessor: "version" },
        { label: "client", accessor: "clientVersion" },
        { label: "isArchive", type: "boolean", accessor: "isArchive" },
        { label: "isTracing", type: "boolean", accessor: "isTracing" },
      ],
    },
    {
      title: "API Keys",
      colSpan: 5,
      fields: [
        { label: "hasEsKey", type: "boolean", accessor: "hasEsKey" },
        { label: "hasPinKey", type: "boolean", accessor: "hasPinKey" },
        { label: "rpcProvider", accessor: "rpcProvider" },
      ],
    },
    {
      title: "Configuration Paths",
      colSpan: 7,
      fields: [
        { label: "rootConfig", accessor: "rootConfig" },
        { label: "chainConfig", accessor: "chainConfig" },
        { label: "indexPath", accessor: "indexPath" },
        { label: "cachePath", accessor: "cachePath" },
      ],
    },
    {
      title: "Statistics",
      colSpan: 5,
      fields: [
        { label: "latestUpdate", accessor: "latestUpdate" },
        { label: "nCaches", type: "int", accessor: "nItems" },
        { label: "nFiles", type: "int", accessor: "nFiles" },
        { label: "nFolders", type: "int", accessor: "nFolders" },
        { label: "sizeInBytes", type: "bytes", accessor: "nBytes" },
      ],
    },
    {
      title: "Caches",
      fields: [],
      components: [
        {
          component: <DataTable<types.CacheItem> table={table} loading={false} pagerName="status" />,
        },
      ],
    },
  ];
}
