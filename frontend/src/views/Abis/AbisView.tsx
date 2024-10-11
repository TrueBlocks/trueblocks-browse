import { getCoreRowModel, useReactTable, Table, ColumnDef } from "@tanstack/react-table";
import { View, FormTable, DataTable, GroupDefinition, EditableTable } from "@components";
import { ModifyAbi } from "@gocode/app/App";
import { types } from "@gocode/models";
import { useAppState, ViewStateProvider } from "@state";
import { tableColumns } from "./AbisTable";

export function AbisView() {
  const { abis, fetchAbis } = useAppState();

  const table = useReactTable({
    data: abis.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <ViewStateProvider route={"abis"} nItems={abis.nItems} fetchFn={fetchAbis} modifyFn={ModifyAbi}>
      <View>
        <FormTable data={abis} definition={createAbisForm(table)} />
      </View>
    </ViewStateProvider>
  );
}

type theInstance = InstanceType<typeof types.AbiContainer>;
function createAbisForm(table: Table<types.Abi>): GroupDefinition<theInstance>[] {
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
          component: <DataTable<types.Abi> table={table} loading={false} />,
        },
      ],
    },
  ];
}
