import { Stack } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { DataTable, FormTable, GroupDefinition, View, SpecButton, PinButton } from "@components";
import { ModifyNoop } from "@gocode/app/App";
import { types } from "@gocode/models";
import { useAppState, ViewStateProvider } from "@state";
import { tableColumns } from "./IndexesTable";

export const IndexesView = () => {
  const { indexes, fetchIndexes } = useAppState();

  const table = useReactTable({
    data: indexes.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <ViewStateProvider route={"indexes"} nItems={indexes.nItems} fetchFn={fetchIndexes} modifyFn={ModifyNoop}>
      <View>
        <FormTable data={indexes} definition={createIndexForm(table)} />
      </View>
    </ViewStateProvider>
  );
};

type theInstance = InstanceType<typeof types.IndexContainer>;
const createIndexForm = (table: any): GroupDefinition<theInstance>[] => {
  return [
    {
      title: "Index Data",
      colSpan: 5,
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
      colSpan: 5,
      fields: [
        { label: "nItems", type: "int", accessor: "nItems" },
        { label: "addrsPerBlock", type: "float", accessor: "addrsPerBlock" },
        { label: "appsPerAddr", type: "float", accessor: "appsPerAddr" },
        { label: "appsPerBlock", type: "float", accessor: "appsPerBlock" },
      ],
    },
    {
      title: "Buttons",
      colSpan: 2,
      components: [
        {
          component: (
            <Stack>
              <PinButton value="https://trueblocks.io">Pin</PinButton>
              <SpecButton value="https://trueblocks.io/papers/2023/specification-for-the-unchained-index-v2.0.0-release.pdf">
                Spec
              </SpecButton>
            </Stack>
          ),
        },
      ],
    },
    {
      title: "Chunks",
      components: [
        {
          component: <DataTable<types.ChunkStats> table={table} loading={false} />,
        },
      ],
    },
  ];
};
