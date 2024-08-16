import React from "react";
import { types } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns } from "./AbisTable";
import { View, FormTable, DataTable, GroupDefinition } from "@components";
import { useAppState } from "@state";

export function AbisView() {
  const { abis } = useAppState();

  const table = useReactTable({
    data: abis.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <FormTable data={abis} definition={createAbisForm(table)} />
    </View>
  );
}

type theInstance = InstanceType<typeof types.AbiContainer>;
function createAbisForm(table: any): GroupDefinition<theInstance>[] {
  return [
    {
      title: "Abi Data",
      colSpan: 6,
      fields: [
        { label: "nItems", type: "int", accessor: "nItems" },
        { label: "nFunctions", type: "int", accessor: "nFunctions" },
        { label: "nEvents", type: "int", accessor: "nEvents" },
        { label: "fileSize", type: "bytes", accessor: "fileSize" },
      ],
    },
    {
      title: "Bounds",
      colSpan: 6,
      fields: [
        { label: "largestFile", type: "text", accessor: "largestFile" },
        { label: "mostFunctions", type: "text", accessor: "mostFunctions" },
        { label: "mostEvents", type: "text", accessor: "mostEvents" },
      ],
    },
    {
      title: "Files",
      fields: [],
      components: [
        {
          component: <DataTable<types.Abi> table={table} loading={false} pagerName="abis" />,
        },
      ],
    },
  ];
}
