import { useEffect } from "react";
import { Text } from "@mantine/core";
import { sdk } from "@gocode/models";
import { useAppState, useViewState } from "@state";

const debug = false;

// eslint-disable-next-line @typescript-eslint/no-unused-vars
export const DebugState = ({ u }: { u: sdk.Updater[] }) => {
  const { counters } = useAppState();
  const { route } = useViewState();

  useEffect(() => {
    if (!counters.current[route]) {
      counters.current[route] = 0;
    }
    counters.current[route] += 1;
  });

  const { info } = useAppState();
  const { nItems } = useViewState();

  if (!debug) {
    return null;
  }

  return (
    <div>
      <Text>{`info.Address: ${info.address}`}</Text>
      <Text>{`nItems: ${nItems}`}</Text>
      {/* <Text>{`updater: ${JSON.stringify(u, null, 2)}`}</Text> */}
      <Text>{`renderCount: ${counters.current[route]}`}</Text>
    </div>
  );
};
