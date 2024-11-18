// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm, DebugState } from "@components";
import { ModifyAbi } from "@gocode/app/App";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { AbisTableDef, AbisFormDef } from ".";
// EXISTING_CODE

export const AbisView = () => {
  const { abis, fetchAbis } = useAppState();
  const { enterNoop } = useNoops();

  // EXISTING_CODE
  // EXISTING_CODE

  const table = useReactTable({
    data: abis.items || [],
    columns: AbisTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "abis";
  const tabs = ["abis"];
  const forms: ViewForm = {
    abis: <FormTable data={abis} groups={AbisFormDef(table)} />,
  };
  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={abis.nItems}
      fetchFn={fetchAbis}
      onEnter={enterNoop}
      modifyFn={ModifyAbi}
    >
      <DebugState n={abis.lastUpdate} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
