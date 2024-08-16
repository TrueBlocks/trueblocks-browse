import React from "react";
import { types } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns } from "./IndexesTable";
import { View, FormTable, DataTable, GroupDefinition } from "@components";
import { useAppState } from "@state";

export function IndexesView() {
  const { indexes } = useAppState();

  const table = useReactTable({
    data: indexes.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <FormTable data={indexes} definition={createIndexForm(table)} />
    </View>
  );
}

type theInstance = InstanceType<typeof types.IndexContainer>;
function createIndexForm(table: any): GroupDefinition<theInstance>[] {
  return [
    {
      title: "Index Data",
      colSpan: 6,
      fields: [
        { label: "bloomSz", type: "bytes", accessor: "bloomSz" },
        { label: "chunkSz", type: "bytes", accessor: "chunkSz" },
        { label: "nAddrs", type: "int", accessor: "nAddrs" },
        { label: "nApps", type: "int", accessor: "nApps" },
        { label: "nBlocks", type: "int", accessor: "nBlocks" },
        { label: "nBlooms", type: "int", accessor: "nBlooms" },
      ],
    },
    {
      title: "Statistics",
      colSpan: 6,
      fields: [
        { label: "nItems", type: "int", accessor: "nItems" },
        { label: "addrsPerBlock", type: "float", accessor: "addrsPerBlock" },
        { label: "appsPerAddr", type: "float", accessor: "appsPerAddr" },
        { label: "appsPerBlock", type: "float", accessor: "appsPerBlock" },
      ],
    },
    {
      title: "Chunks",
      fields: [],
      components: [
        {
          component: <DataTable<types.ChunkStats> table={table} loading={false} pagerName="indexes" />,
        },
      ],
    },
  ];
}
