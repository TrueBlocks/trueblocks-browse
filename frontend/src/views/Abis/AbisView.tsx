import { getCoreRowModel, useReactTable, Table } from "@tanstack/react-table";
import { View, FormTable, DataTable, FieldGroup, PublishButton, CleanButton, AddButton } from "@components";
import { ModifyAbi } from "@gocode/app/App";
import { types } from "@gocode/models";
import { useAppState, ViewStateProvider } from "@state";
import { tableColumns } from "./AbisTable";

export const AbisView = () => {
  const { abis, fetchAbis } = useAppState();

  const table = useReactTable({
    data: abis.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "abis";
  return (
    <ViewStateProvider route={route} nItems={abis.nItems} fetchFn={fetchAbis} modifyFn={ModifyAbi}>
      <View>
        <FormTable data={abis} groups={createAbisForm(table)} />
      </View>
    </ViewStateProvider>
  );
};

const createAbisForm = (table: Table<types.Abi>): FieldGroup<types.AbiContainer>[] => {
  return [
    {
      label: "Abi Data",
      colSpan: 6,
      fields: [
        { label: "nItems", type: "int", accessor: "nItems" },
        { label: "nFunctions", type: "int", accessor: "nFunctions" },
        { label: "nEvents", type: "int", accessor: "nEvents" },
        { label: "fileSize", type: "bytes", accessor: "fileSize" },
      ],
    },
    {
      label: "Bounds",
      colSpan: 6,
      fields: [
        { label: "largestFile", type: "text", accessor: "largestFile" },
        { label: "mostFunctions", type: "text", accessor: "mostFunctions" },
        { label: "mostEvents", type: "text", accessor: "mostEvents" },
      ],
    },
    {
      label: "Buttons",
      buttons: [
        <AddButton key={"add"} value={"https://trueblocks.io"} />,
        <CleanButton key={"clean"} value={"https://trueblocks.io"} />,
        <PublishButton key={"publish"} value={"https://trueblocks.io"} />,
      ],
    },
    {
      label: "Files",
      collapsable: false,
      components: [<DataTable<types.Abi> key={"dataTable"} table={table} loading={false} />],
    },
  ];
};
