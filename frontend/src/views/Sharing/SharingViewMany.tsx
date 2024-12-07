// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { useCallback } from "react";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { DebugState, FormTable, View, ViewForm } from "@components";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { AbisFormDef, AbisTableDef } from "../Abis";
import { NamesFormDef, NamesTableDef } from "../Names";
// EXISTING_CODE

export const SharingView = () => {
  const { names, fetchNames, abis, fetchAbis } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  // eslint-disable-next-line prefer-const
  let customTabs: string[] = [];
  // eslint-disable-next-line prefer-const
  let customForms: Record<string, JSX.Element> = {};
  // EXISTING_CODE
  customTabs = ["pin", "upload"];
  customForms["pin"] = <div>This is a custom tab</div>;
  // EXISTING_CODE

  const fetchSharing = useCallback(
    (currentItem: number, itemsPerPage: number) => {
      fetchNames(currentItem, itemsPerPage);
      fetchAbis(currentItem, itemsPerPage);
    },
    [fetchNames, fetchAbis]
  );

  const namesTable = useReactTable({
    data: names?.items || [],
    columns: NamesTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const abisTable = useReactTable({
    data: abis?.items || [],
    columns: AbisTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const tabs = ["names", "abis", ...(customTabs || [])];
  const forms: ViewForm = {
    names: <FormTable data={names} groups={NamesFormDef(namesTable)} />,
    abis: <FormTable data={abis} groups={AbisFormDef(abisTable)} />,
    ...customForms,
  };

  // if (!(status?.items?.length > 0)) {
  //   return <>Loading...</>;
  // }

  return (
    <ViewStateProvider
      // do not remove - delint
      nItems={names.nItems}
      fetchFn={fetchSharing}
      onEnter={handleEnter}
      modifyFn={handleModify}
      tabs={tabs}
    >
      <DebugState u={[names.updater, abis.updater]} />
      <View forms={forms} searchable />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
