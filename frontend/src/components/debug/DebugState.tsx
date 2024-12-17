import { useEffect } from "react";
import { Text } from "@mantine/core";
import { sdk } from "@gocode/models";
import { useAppState, useViewState } from "@state";

const debug = import.meta.env.VITE_DEBUG === "true";

// eslint-disable-next-line @typescript-eslint/no-unused-vars
export const DebugState = ({ u }: { u: sdk.Updater[] }) => {
  const { counters, route, activeTab, headerOn, info } = useAppState();
  const { nItems } = useViewState();

  useEffect(() => {
    if (!counters.current[route]) {
      counters.current[route] = 0;
    }
    counters.current[route] += 1;
  });

  if (!debug) {
    return null;
  }

  return (
    <div>
      <Text>{`info.Address: ${info.address}`}</Text>
      <Text>{`nItems: ${nItems}`}</Text>
      <Text>{`renderCount: ${counters.current[route]}`}</Text>
      <Text>{`route: ${route} activeTab: ${activeTab} headerOn: ${headerOn}`}</Text>
    </div>
  );
};
