import React from "react";
import { types } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns } from "./NamesTable";
import { View, FormTable, DataTable, GroupDefinition } from "@components";
import { useAppState, ViewStateProvider } from "@state";

export function NamesView() {
  const { names, fetchNames } = useAppState();

  const table = useReactTable({
    data: names.names || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <ViewStateProvider route={"names"} nItems={names.nItems} fetchFn={fetchNames}>
      <View>
        <FormTable data={names} definition={createNameForm(table)} />
      </View>
    </ViewStateProvider>
  );
}

type theInstance = InstanceType<typeof types.NameContainer>;
function createNameForm(table: any): GroupDefinition<theInstance>[] {
  return [
    {
      title: "Name Data",
      colSpan: 6,
      fields: [
        { label: "nNames", type: "int", accessor: "nItems" },
        { label: "nContracts", type: "int", accessor: "nContracts" },
        { label: "nErc20s", type: "int", accessor: "nErc20s" },
        { label: "nErc721s", type: "int", accessor: "nErc721s" },
        { label: "nDeleted", type: "int", accessor: "nDeleted" },
      ],
    },
    {
      title: "Database Parts",
      colSpan: 6,
      fields: [
        { label: "sizeOnDisc", type: "bytes", accessor: "sizeOnDisc" },
        { label: "nCustom", type: "int", accessor: "nCustom" },
        { label: "nRegular", type: "int", accessor: "nRegular" },
        { label: "nPrefund", type: "int", accessor: "nPrefund" },
        { label: "nBaddress", type: "int", accessor: "nBaddress" },
      ],
    },
    {
      title: "Names",
      fields: [],
      components: [
        {
          component: <DataTable<types.Name> table={table} loading={false} />,
        },
      ],
    },
  ];
}
