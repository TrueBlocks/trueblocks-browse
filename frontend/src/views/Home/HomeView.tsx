import React from "react";
import { Text } from "@mantine/core";
import { View2 } from "@components";
import { useAppState } from "@state";

export function HomeView() {
  const { names, abis, indexes, manifests, status } = useAppState();
  return (
    <View2>
      <Text>Home View Content</Text>
      <Text>{`nNames: ${names.nItems}`}</Text>
      <Text>{`nAbis: ${abis.nItems}`}</Text>
      <Text>{`nIndexes: ${indexes.nItems}`}</Text>
      <Text>{`nManifests: ${manifests.nItems}`}</Text>
      <Text>{`nCaches: ${status.nItems}`}</Text>
    </View2>
  );
}
