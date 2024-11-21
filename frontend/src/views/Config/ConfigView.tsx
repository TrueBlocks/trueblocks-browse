// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { DebugState, FormTable, View, ViewForm } from "@components";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { ConfigFormDef } from "./ConfigFormDef";
// EXISTING_CODE

export const ConfigView = () => {
  const { config, fetchConfig } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  // EXISTING_CODE
  const table = config;
  // EXISTING_CODE

  const route = "config";
  const tabs = ["config"];
  const forms: ViewForm = {
    config: <FormTable data={config} groups={ConfigFormDef(table)} />,
  };

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={0}
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
