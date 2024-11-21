import { Text } from "@mantine/core";
import { useRenderCounter } from "@hooks";
import { useAppState, useViewState } from "@state";
// import { types } from "../../../wailsjs/go/models";

const debug = true;

export const DebugState = ({ u }: { u: any }) => {
  const { info } = useAppState();
  const { nItems } = useViewState();
  const renderCount = useRenderCounter();

  if (!debug) {
    return null;
  }

  return (
    <div>
      <Text>{`info.Address: ${info.address}`}</Text>
      <Text>{`nItems: ${nItems}`}</Text>
      <Text>{`updater: ${JSON.stringify(u, null, 2)}`}</Text>
      <Text>{`renderCount: ${renderCount}`}</Text>
    </div>
  );
};
