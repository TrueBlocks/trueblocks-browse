import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { DebugState } from "../../components";

export const ConfigView = () => {
  const { config, fetchConfig } = useAppState();
  const { modifyNoop } = useNoops();

  const route = "config";
  // const tabs = ["config"];
  // const forms: ViewForm = {
  //   names: <div>ViewForm</div>,
  // };
  return (
    <ViewStateProvider route={route} nItems={0} fetchFn={fetchConfig} modifyFn={modifyNoop}>
      <DebugState n={config.lastUpdate} />
      <pre>{JSON.stringify(config?.settings, null, 2)}</pre>
    </ViewStateProvider>
  );
};
