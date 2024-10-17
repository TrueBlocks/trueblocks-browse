import { Stack } from "@mantine/core";
import { getCoreRowModel, useReactTable, Table } from "@tanstack/react-table";
import { View, FormTable, DataTable, GroupDefinition, PublishButton, CleanButton } from "@components";
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

  return (
    <ViewStateProvider route={"abis"} nItems={abis.nItems} fetchFn={fetchAbis} modifyFn={ModifyAbi}>
      <View>
        <FormTable data={abis} definition={createAbisForm(table)} />
      </View>
    </ViewStateProvider>
  );
};

type theInstance = InstanceType<typeof types.AbiContainer>;
const createAbisForm = (table: Table<types.Abi>): GroupDefinition<theInstance>[] => {
  return [
    {
      legend: "Abi Data",
      colSpan: 5,
      fields: [
        { label: "nItems", type: "int", accessor: "nItems" },
        { label: "nFunctions", type: "int", accessor: "nFunctions" },
        { label: "nEvents", type: "int", accessor: "nEvents" },
        { label: "fileSize", type: "bytes", accessor: "fileSize" },
      ],
    },
    {
      legend: "Bounds",
      colSpan: 5,
      fields: [
        { label: "largestFile", type: "text", accessor: "largestFile" },
        { label: "mostFunctions", type: "text", accessor: "mostFunctions" },
        { label: "mostEvents", type: "text", accessor: "mostEvents" },
      ],
    },
    {
      legend: "Buttons",
      colSpan: 2,
      components: [
        {
          component: (
            <Stack>
              <PublishButton value={"https://trueblocks.io"}>Publish</PublishButton>
              <CleanButton value={"https://trueblocks.io"}>Clean</CleanButton>
            </Stack>
          ),
        },
      ],
    },
    {
      legend: "Files",
      components: [
        {
          component: <DataTable<types.Abi> table={table} loading={false} />,
        },
      ],
    },
  ];
};
