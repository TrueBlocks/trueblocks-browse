import React from "react";
import { app, types, messages } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns } from "./HomeTable";
import { View, FormTable, DataTable, GroupDefinition } from "@components";
import { useAppState, ViewStateProvider } from "@state";
import { Page } from "@hooks";
import { SetLast } from "@gocode/app/App";
import { EventsEmit } from "@runtime";

export function HomeView() {
  const { home, fetchHome } = useAppState();

  const handleEnter = (page: Page) => {
    const record = page.selected - page.getOffset();
    const address = home.items[record].address;
    SetLast("route", `/history/${address}`);
    EventsEmit(messages.Message.NAVIGATE, {
      route: `/history/${address}`,
    });
  };

  const table = useReactTable({
    data: home.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <ViewStateProvider route={""} nItems={home.items?.length} fetchFn={fetchHome} onEnter={handleEnter}>
      <View>
        <FormTable data={home} definition={createHomeForm(table)} />
      </View>
    </ViewStateProvider>
  );
}

type theInstance = InstanceType<typeof app.HomeContainer>;
function createHomeForm(table: any): GroupDefinition<theInstance>[] {
  return [
    {
      title: "Open Monitors",
      fields: [],
      components: [
        {
          component: <DataTable<types.HistoryContainer> table={table} loading={false} />,
        },
      ],
    },
    {
      title: "Data 1",
      colSpan: 3,
      fields: [
        { label: "nMonitors", type: "int", accessor: "nMonitors" },
        { label: "nNames", type: "int", accessor: "nNames" },
        { label: "nAbis", type: "int", accessor: "nAbis" },
      ],
    },
    {
      title: "Data 2",
      colSpan: 3,
      fields: [
        { label: "nIndexes", type: "int", accessor: "nIndexes" },
        { label: "nManifests", type: "int", accessor: "nManifests" },
        { label: "nCaches", type: "int", accessor: "nCaches" },
      ],
    },
  ];
}

// import React from "react";
// import { Text, Group } from "@mantine/core";
// import { View, Formatter } from "@components";
// import { useAppState, ViewStateProvider } from "@state";
// import { HistorySize } from "@gocode/app/App";

// export function HomeView() {
//   const [size, setSize] = React.useState(0);
//   const { address, getCounters } = useAppState();
//   const { fetchHistory, fetchMonitors, fetchNames, fetchAbis, fetchIndexes, fetchManifests, fetchStatus } =
//     useAppState();

//   if (!address) {
//     return <Text>Address not found</Text>;
//   }

//   const fetchFn = (selected: number, perPage: number, item?: any) => {
//     fetchHistory(selected, perPage, item);
//     fetchMonitors(selected, perPage, item);
//     fetchNames(selected, perPage, item);
//     fetchAbis(selected, perPage, item);
//     fetchIndexes(selected, perPage, item);
//     fetchManifests(selected, perPage, item);
//     fetchStatus(selected, perPage, item);
//     HistorySize(address as unknown as string).then((size) => setSize(size));
//   };

//   var counters = getCounters();
//   return (
//     <ViewStateProvider route="" fetchFn={fetchFn}>
//       <View>
//         <Text>Current State</Text>
//         <Group>
//           <Text>current address:</Text> <Formatter type="address-editor" value={address} />
//         </Group>
//         <Group>
//           <Text>current address size:</Text> <Formatter type="bytes" value={size} />
//         </Group>
//         <Group>
//           <Text>nMonitors:</Text> <Formatter type="int" value={counters.nMonitors} />
//         </Group>
//         <Group>
//           <Text>nNames:</Text> <Formatter type="int" value={counters.nNames} />
//         </Group>
//         <Group>
//           <Text>nAbis:</Text> <Formatter type="int" value={counters.nAbis} />
//         </Group>
//         <Group>
//           <Text>nIndexes:</Text> <Formatter type="int" value={counters.nIndexes} />
//         </Group>
//         <Group>
//           <Text>nManifests:</Text> <Formatter type="int" value={counters.nManifests} />
//         </Group>
//         <Group>
//           <Text>nCaches:</Text> <Formatter type="int" value={counters.nStatus} />
//         </Group>
//       </View>
//     </ViewStateProvider>
//   );
// }
