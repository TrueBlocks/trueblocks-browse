// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { DebugState } from "@components";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
// EXISTING_CODE

export const ConfigView = () => {
  const { config, fetchConfig } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();

  // EXISTING_CODE
  // EXISTING_CODE

  const route = "config";
  // const tabs = ["config"];
  // const forms: ViewForm = {
  //   names: <div>ViewForm</div>,
  // };
  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={0}
      fetchFn={fetchConfig}
      onEnter={enterNoop}
      modifyFn={modifyNoop}
    >
      <DebugState n={config.lastUpdate} />
      <pre>{JSON.stringify(config, null, 2)}</pre>
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
