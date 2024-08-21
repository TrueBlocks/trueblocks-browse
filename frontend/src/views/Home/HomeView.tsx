import React from "react";
import { Text, Group } from "@mantine/core";
import { View, Formatter } from "@components";
import { useAppState, ViewStateProvider } from "@state";
import { HistorySize } from "@gocode/app/App";

export function HomeView() {
  const [size, setSize] = React.useState(0);
  const { address, getCounters } = useAppState();
  const { fetchHistory, fetchMonitors, fetchNames, fetchAbis, fetchIndexes, fetchManifests, fetchStatus } =
    useAppState();

  const fetchFn = (curItem: number, perPage: number, item?: any) => {
    // fetchHistory(curItem, perPage, item);
    fetchMonitors(curItem, perPage, item);
    fetchNames(curItem, perPage, item);
    fetchAbis(curItem, perPage, item);
    fetchIndexes(curItem, perPage, item);
    fetchManifests(curItem, perPage, item);
    fetchStatus(curItem, perPage, item);
    HistorySize(address as unknown as string).then((size) => setSize(size));
  };

  var counters = getCounters();
  return (
    <ViewStateProvider route="" fetchFn={fetchFn}>
      <View>
        <Text>Current State</Text>
        <Group>
          <Text>current address:</Text> <Formatter type="address-address-only" value={address} />
        </Group>
        <Group>
          <Text>current address size:</Text> <Formatter type="bytes" value={size} />
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
    </ViewStateProvider>
  );
}
