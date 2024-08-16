import React from "react";
import { Text } from "@mantine/core";
import { View2 } from "@components";
import { useAppState } from "@state";

export function HomeView() {
  const { manifests, status, indexes } = useAppState();
  return (
    <View2>
      <Text>Home View Content</Text>
      <Text>{`nManifests: ${manifests.nItems}`}</Text>
      <Text>{`nCaches: ${status.nItems}`}</Text>
      <Text>{`nIndexes: ${indexes.nItems}`}</Text>
    </View2>
  );
}
