import { ReactNode } from "react";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { DataTable, FormTable, FieldGroup, View, SpecButton, PinButton } from "@components";
import { types } from "@gocode/models";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { tableColumns } from "./IndexesTable";

export const IndexesView = () => {
  const { modifyNoop } = useNoops();
  const { indexes, fetchIndexes } = useAppState();

  const table = useReactTable({
    data: indexes.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "indexes";
  const tabs = ["indexes"];
  const forms: Record<string, ReactNode> = {
    indexes: <FormTable data={indexes} groups={createIndexForm(table)} />,
  };
  return (
    <ViewStateProvider route={route} nItems={indexes.nItems} fetchFn={fetchIndexes} modifyFn={modifyNoop}>
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

const createIndexForm = (table: any): FieldGroup<types.IndexContainer>[] => {
  return [
    {
      label: "Index Data",
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
      label: "Statistics",
      colSpan: 6,
      fields: [
        { label: "nItems", type: "int", accessor: "nItems" },
        { label: "addrsPerBlock", type: "float", accessor: "addrsPerBlock" },
        { label: "appsPerAddr", type: "float", accessor: "appsPerAddr" },
        { label: "appsPerBlock", type: "float", accessor: "appsPerBlock" },
      ],
    },
    {
      label: "Buttons",
      buttons: [
        <PinButton key={"pin"} value="https://trueblocks.io" />,
        <SpecButton
          key={"spec"}
          value="https://trueblocks.io/papers/2023/specification-for-the-unchained-index-v2.0.0-release.pdf"
        />,
      ],
    },
    {
      label: "Chunks",
      collapsable: false,
      components: [<DataTable<types.ChunkStats> key={"dataTable"} table={table} loading={false} />],
    },
  ];
};
