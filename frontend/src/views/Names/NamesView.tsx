import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm, DebugState } from "@components";
import { LoadAddress, ModifyName } from "@gocode/app/App";
import { Page } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { NamesFormDef, NamesTableDef } from ".";

export const NamesView = () => {
  const { names, fetchNames } = useAppState();

  const handleEnter = (page: Page) => {
    const address = names.items[page.getRecord()].address;
    const addressStr = address as unknown as string;
    LoadAddress(addressStr).then(() => {});
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
