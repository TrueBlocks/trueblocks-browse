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
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  // EXISTING_CODE
  // EXISTING_CODE

  const route = "config";

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={0}
      fetchFn={fetchConfig}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState n={config.lastUpdate} />
      <pre>{JSON.stringify(config, null, 2)}</pre>
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE

//-------------------------------------------------------------------
// Template variables:
// class:         Config
// lower:         config
// routeLabel:    Config
// routeLower:    config
