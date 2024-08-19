import React from "react";
import { Text, Group } from "@mantine/core";
import { useAppState } from "@state";
import { Formatter } from "@components";

export function IndexStatus() {
  const { meta } = useAppState();

  return (
    <Group justify={"space-between"}>
      <Text size="sm">unchained index: </Text>
      {" / "}
      <Formatter size="sm" type="int" value={meta.client} />
      {" / "}
      <Formatter size="sm" type="int" value={meta.finalized - meta.client} />
      {" / "}
      <Formatter size="sm" type="int" value={meta.ripe - meta.client} />{" "}
    </Group>
  );
}
