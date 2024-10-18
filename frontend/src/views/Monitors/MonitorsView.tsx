import { Stack } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, DataTable, FieldGroup, CleanButton } from "@components";
import { GoToHistory, ModifyMonitors } from "@gocode/app/App";
import { types } from "@gocode/models";
import { Page } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { tableColumns } from "./MonitorsTable";

export const MonitorsView = () => {
  const { monitors, fetchMonitors } = useAppState();

  const handleEnter = (page: Page) => {
    const address = monitors.items[page.getRecord()].address;
    GoToHistory(address).then(() => {});
  };

  const table = useReactTable({
    data: monitors.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "monitors";
  return (
    <ViewStateProvider
      route={route}
      nItems={monitors.nItems}
      fetchFn={fetchMonitors}
      onEnter={handleEnter}
      modifyFn={ModifyMonitors}
    >
      <View>
        <FormTable data={monitors} groups={createMonitorForm(table)} />
      </View>
    </ViewStateProvider>
  );
};

type theInstance = InstanceType<typeof types.MonitorContainer>;
const createMonitorForm = (table: any): FieldGroup<theInstance>[] => {
  return [
    {
      legend: "Monitor Data",
      colSpan: 5,
      fields: [
        { label: "nMonitors", type: "int", accessor: "nItems" },
        { label: "nRecords", type: "int", accessor: "nRecords" },
        { label: "nNamed", type: "int", accessor: "nNamed" },
        { label: "fileSize", type: "bytes", accessor: "fileSize" },
      ],
    },
    {
      legend: "Other",
      colSpan: 5,
      fields: [
        { label: "nEmpty", type: "int", accessor: "nEmpty" },
        { label: "nStaged", type: "int", accessor: "nStaged" },
        { label: "nDeleted", type: "int", accessor: "nDeleted" },
      ],
    },
    {
      legend: "Buttons",
      colSpan: 2,
      components: [
        {
          component: (
            <Stack align="center">
              <CleanButton value={"https://trueblocks.io"}>Clean</CleanButton>
            </Stack>
          ),
        },
      ],
    },
    {
      legend: "Available Monitors",
      collapsable: false,
      components: [
        {
          component: <DataTable<types.Monitor> table={table} loading={false} />,
        },
      ],
    },
  ];
};
