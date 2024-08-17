import React from "react";
import { Text } from "@mantine/core";
import { useAppState } from "@state";

export const AppStatus = () => {
  const { status } = useAppState();
  const chain = status.chain ? status.chain : "mainnet";
  return <Text size={"sm"}>{`${status.clientVersion} / ${chain} / ${"not loaded"} / ${status.latestUpdate}`}</Text>;
};
