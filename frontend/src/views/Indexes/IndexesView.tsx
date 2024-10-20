import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { DataTable, FormTable, FieldGroup, View, SpecButton, PinButton } from "@components";
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

  const route = "indexes";
  return (
    <ViewStateProvider route={route} nItems={indexes.nItems} fetchFn={fetchIndexes} modifyFn={ModifyNoop}>
      <View>
        <FormTable data={indexes} groups={createIndexForm(table)} />
      </View>
    </ViewStateProvider>
  );
};

const createIndexForm = (table: any): FieldGroup<types.IndexContainer>[] => {
  return [
    {
      legend: "Index Data",
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
      legend: "Statistics",
      colSpan: 6,
      fields: [
        { label: "nItems", type: "int", accessor: "nItems" },
        { label: "addrsPerBlock", type: "float", accessor: "addrsPerBlock" },
        { label: "appsPerAddr", type: "float", accessor: "appsPerAddr" },
        { label: "appsPerBlock", type: "float", accessor: "appsPerBlock" },
      ],
    },
    {
      legend: "Buttons",
      buttons: [
        <PinButton key={"pin"} value="https://trueblocks.io" />,
        <SpecButton
          key={"spec"}
          value="https://trueblocks.io/papers/2023/specification-for-the-unchained-index-v2.0.0-release.pdf"
        />,
      ],
    },
    {
      legend: "Chunks",
      collapsable: false,
      components: [
        {
          component: <DataTable<types.ChunkStats> table={table} loading={false} />,
        },
      ],
    },
  ];
};
