import { Text } from "@mantine/core";
import { useRenderCounter } from "../../hooks";
import { useViewState } from "../../state";

export const DebugState = ({ n }: { n: number }) => {
  const { nItems } = useViewState();
  const renderCount = useRenderCounter();
  return (
    <div>
      <Text>nItems: {nItems}</Text>
      <Text>lastUpdate: {n}</Text>
      <Text>renderCount: {renderCount}</Text>
    </div>
  );
};
