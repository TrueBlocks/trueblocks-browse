import React from "react";
import { Text } from "@mantine/core";
import { useAppState } from "@state";

export function IndexStatus() {
  const { indexes } = useAppState();

  if (!indexes.items) {
    return <Text size="sm">loading indexes...</Text>;
  }

  return <Text size="sm">{`unchained index: block ${indexes.items[indexes.nItems - 1]} ${indexes.nItems}`}</Text>;
}
