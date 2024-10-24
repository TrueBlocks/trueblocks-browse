import { ReactNode } from "react";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, DataTable, FieldGroup, CleanButton, PublishButton, AddButton } from "@components";
import { GoToHistory, ModifyName } from "@gocode/app/App";
import { types } from "@gocode/models";
import { Page } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { tableColumns } from "./NamesTable";

export const NamesView = () => {
  const { names, fetchNames } = useAppState();

  const handleEnter = (page: Page) => {
    const address = names.names[page.getRecord()].address;
    GoToHistory(address).then(() => {});
  };

  const table = useReactTable({
    data: names.names || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "names";
  const tabs = ["names"];
  const forms: Record<string, ReactNode> = {
    names: <FormTable data={names} groups={createNameForm(table)} />,
  };
  return (
    <ViewStateProvider
      route={route}
      nItems={names.nItems}
      fetchFn={fetchNames}
      onEnter={handleEnter}
      modifyFn={ModifyName}
    >
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

const createNameForm = (table: any): FieldGroup<types.NamesContainer>[] => {
  return [
    {
      label: "Name Data",
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
      label: "Database Parts",
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
      label: "Buttons",
      buttons: [
        <AddButton key={"add"} value={"https://trueblocks.io"} />,
        <CleanButton key={"clean"} value={"https://trueblocks.io"} />,
        <PublishButton key={"publish"} value={"https://trueblocks.io"} />,
      ],
    },
    {
      label: "Names",
      collapsable: false,
      components: [<DataTable<types.Name> key={"dataTable"} table={table} loading={false} />],
    },
  ];
};
