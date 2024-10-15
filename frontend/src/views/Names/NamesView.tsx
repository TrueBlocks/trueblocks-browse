import { useState } from "react";
import { Stack } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { NameEditor, BaseButton } from "@components";
import { View, FormTable, DataTable, GroupDefinition } from "@components";
import { GoToHistory, ModifyName } from "@gocode/app/App";
import { types } from "@gocode/models";
import { Page } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { tableColumns } from "./NamesTable";

export const NamesView = () => {
  const { names, fetchNames } = useAppState();
  const [showEditor, setShowEditor] = useState(false);

  const handleEnter = (page: Page) => {
    const record = page.selected - page.getOffset();
    const address = names.names[record].address;
    GoToHistory(address).then(() => {});
  };

  const table = useReactTable({
    data: names.names || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <ViewStateProvider
      route={"names"}
      nItems={names.nItems}
      fetchFn={fetchNames}
      onEnter={handleEnter}
      modifyFn={ModifyName}
    >
      <View>
        <Stack justify="space-between">
          <BaseButton onClick={() => setShowEditor(!showEditor)}>{showEditor ? "Hide" : "Show"}</BaseButton>
          {showEditor ? <NameEditor /> : <></>}
        </Stack>
        <FormTable data={names} definition={createNameForm(table)} />
      </View>
    </ViewStateProvider>
  );
};

type theInstance = InstanceType<typeof types.NameContainer>;
const createNameForm = (table: any): GroupDefinition<theInstance>[] => {
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
        { label: "nSystem", type: "int", accessor: "nSystem" },
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
};
