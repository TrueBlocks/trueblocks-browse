import React from "react";
import { Text, Group } from "@mantine/core";
import { View, Formatter } from "@components";
import { useAppState } from "@state";

export function HomeView() {
  const { address, getCounters } = useAppState();

  var counters = getCounters();
  return (
    <View route="">
      <Text>Current State</Text>
      <Group>
        <Text>current address:</Text> <Formatter type="address-address-only" value={address} />
      </Group>
      <Group>
        <Text>nMonitors:</Text> <Formatter type="int" value={counters.nMonitors} />
      </Group>
      <Group>
        <Text>nNames:</Text> <Formatter type="int" value={counters.nNames} />
      </Group>
      <Group>
        <Text>nAbis:</Text> <Formatter type="int" value={counters.nAbis} />
      </Group>
      <Group>
        <Text>nIndexes:</Text> <Formatter type="int" value={counters.nIndexes} />
      </Group>
      <Group>
        <Text>nManifests:</Text> <Formatter type="int" value={counters.nManifests} />
      </Group>
      <Group>
        <Text>nCaches:</Text> <Formatter type="int" value={counters.nStatus} />
      </Group>
    </View>
  );
}
