import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm, DebugState } from "@components";
import { ModifyName } from "@gocode/app/App";
import { Page } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { NamesFormDef, NamesTableDef } from ".";

export const NamesView = () => {
  const { names, fetchNames, loadAddress } = useAppState();

  const handleEnter = (page: Page) => {
    loadAddress(names.items[page.getRecord()].address);
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
      <DebugState n={names.lastUpdate} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};
