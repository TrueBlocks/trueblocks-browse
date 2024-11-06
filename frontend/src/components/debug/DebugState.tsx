import { Text } from "@mantine/core";
import { useRenderCounter } from "../../hooks";

export const DebugState = ({ n }: { n: number }) => {
  const renderCount = useRenderCounter();
  return (
    <div>
      <Text>lastUpdate: {n}</Text>
      <Text>renderCount: {renderCount}</Text>
    </div>
  );
};
