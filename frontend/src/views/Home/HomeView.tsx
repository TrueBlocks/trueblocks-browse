import React from "react";
import { Text } from "@mantine/core";
import { View } from "@components";
import { useAppState } from "@state";

export function HomeView() {
  const { monitors, names, abis, indexes, manifests, status } = useAppState();
  return (
    <View>
      <Text>Current State</Text>
      <Text>{`nMonitors: ${monitors.nItems}`}</Text>
      <Text>{`nNames: ${names.nItems}`}</Text>
      <Text>{`nAbis: ${abis.nItems}`}</Text>
      <Text>{`nIndexes: ${indexes.nItems}`}</Text>
      <Text>{`nManifests: ${manifests.nItems}`}</Text>
      <Text>{`nCaches: ${status.nItems}`}</Text>
    </View>
  );
}
