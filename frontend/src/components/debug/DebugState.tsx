import { Text } from "@mantine/core";
import { useRenderCounter } from "@hooks";
import { useAppState, useViewState } from "@state";

export const DebugState = ({ n }: { n: number }) => {
  const { info } = useAppState();
  const { nItems } = useViewState();
  const renderCount = useRenderCounter();
  return (
    <div>
      <Text>{`info.Address: ${info.address}`}</Text>
      <Text>{`nItems: ${nItems}`}</Text>
      <Text>{`lastUpdate: ${n}`}</Text>
      <Text>{`renderCount: ${renderCount}`}</Text>
    </div>
  );
};
