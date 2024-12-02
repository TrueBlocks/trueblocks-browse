// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { DebugState, FormTable, View, ViewForm } from "@components";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { ConfigFormDef, ConfigTableDef } from ".";
// EXISTING_CODE

export const ConfigView = () => {
  const { config, fetchConfig } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  // EXISTING_CODE
  // EXISTING_CODE

  const table = useReactTable({
    data: [], // config?.items || [],
    columns: ConfigTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "config";
  const tabs = ["config"];
  const forms: ViewForm = {
    config: <FormTable data={config} groups={ConfigFormDef(table)} />,
  };

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={config.nItems}
      fetchFn={fetchConfig}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState u={config.updater} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
