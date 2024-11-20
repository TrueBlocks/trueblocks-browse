// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm, DebugState } from "@components";
import { ModifyName } from "@gocode/app/App";
import { Page } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { NamesFormDef, NamesTableDef } from ".";
// EXISTING_CODE

export const NamesView = () => {
  const { names, fetchNames, loadAddress } = useAppState();
  const handleEnter = (page: Page) => {
    loadAddress(names.items[page.getRecord()].address);
  };
  const handleModify = ModifyName;

  // EXISTING_CODE
  // EXISTING_CODE

  const table = useReactTable({
    data: names?.items || [],
    columns: NamesTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "names";
  const tabs = ["names"];
  const forms: ViewForm = {
    names: <FormTable data={names} groups={NamesFormDef(table)} />,
  };

  // if (!(names?.items?.length > 0)) {
  //   return <>Loading...</>;
  // }

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={names.nItems}
      fetchFn={fetchNames}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState n={names.lastUpdate} />
      <View tabs={tabs} forms={forms} searchable />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
