import { Text } from "@mantine/core";
import { useAppState } from "@state";

export const AppStatus = () => {
  const { chain, status } = useAppState();
  return <Text size={"sm"}>{`${status.clientVersion} / ${chain} / ${"not loaded"} / ${status.lastUpdate}`}</Text>;
};
