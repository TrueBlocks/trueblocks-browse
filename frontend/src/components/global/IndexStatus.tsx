import React from "react";
import { Text, Group } from "@mantine/core";
import { useAppState } from "@state";
import { Formatter } from "@components";

export function IndexStatus() {
  const { indexes, meta } = useAppState();

  if (!indexes.items) {
    return <Text size="sm">loading indexes...</Text>;
  }

  return (
    <Group justify={"space-between"} style={{ marginRight: "1em" }}>
      <Text>unchained index: </Text>
      <Formatter size="sm" type="int" value={indexes.nItems} />
      {" / "}
      <Formatter size="sm" type="int" value={meta.client} />
      {" / "}
      <Formatter size="sm" type="int" value={meta.finalized - meta.client} />
      {" / "}
      <Formatter size="sm" type="int" value={meta.ripe - meta.client} />{" "}
    </Group>
  );
}
