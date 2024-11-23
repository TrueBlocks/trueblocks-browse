import { Text } from "@mantine/core";
import { updater } from "@gocode/models";
import { useRenderCounter } from "@hooks";
import { useAppState, useViewState } from "@state";

const debug = false;

export const DebugState = ({ u }: { u: updater.Updater }) => {
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
