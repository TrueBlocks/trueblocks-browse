import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm } from "@components";
import { GoToAddress, ModifyName } from "@gocode/app/App";
import { Page } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { NamesFormDef, NamesTableDef } from ".";

export const NamesView = () => {
  const { names, fetchNames } = useAppState();

  const handleEnter = (page: Page) => {
    const address = names.items[page.getRecord()].address;
    GoToAddress(address).then(() => {});
  };

  const table = useReactTable({
    data: names.items || [],
    columns: NamesTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "names";
  const tabs = ["names"];
  const forms: ViewForm = {
    names: <FormTable data={names} groups={NamesFormDef(table)} />,
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
