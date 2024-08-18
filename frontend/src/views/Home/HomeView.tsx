import React from "react";
import { Text, Group } from "@mantine/core";
import { View, Formatter } from "@components";
import { useAppState } from "@state";

export function HomeView() {
  const { monitors, names, abis, indexes, manifests, status } = useAppState();
  return (
    <View>
      <Text>Current State</Text>
      <Group>
        <Text>nMonitors:</Text> <Formatter type="int" value={monitors.nItems} />
      </Group>
      <Group>
        <Text>nNames:</Text> <Formatter type="int" value={names.nItems} />
      </Group>
      <Group>
        <Text>nAbis:</Text> <Formatter type="int" value={abis.nItems} />
      </Group>
      <Group>
        <Text>nIndexes:</Text> <Formatter type="int" value={indexes.nItems} />
      </Group>
      <Group>
        <Text>nManifests:</Text> <Formatter type="int" value={manifests.nItems} />
      </Group>
      <Group>
        <Text>nCaches:</Text> <Formatter type="int" value={status.nItems} />
      </Group>
    </View>
  );
}
