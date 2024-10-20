import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, DataTable, FieldGroup, CleanButton, AddButton } from "@components";
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

const createMonitorForm = (table: any): FieldGroup<types.MonitorContainer>[] => {
  return [
    {
      label: "Monitor Data",
      colSpan: 6,
      fields: [
        { label: "nMonitors", type: "int", accessor: "nItems" },
        { label: "nRecords", type: "int", accessor: "nRecords" },
        { label: "nNamed", type: "int", accessor: "nNamed" },
        { label: "fileSize", type: "bytes", accessor: "fileSize" },
      ],
    },
    {
      label: "Other",
      colSpan: 6,
      fields: [
        { label: "nEmpty", type: "int", accessor: "nEmpty" },
        { label: "nStaged", type: "int", accessor: "nStaged" },
        { label: "nDeleted", type: "int", accessor: "nDeleted" },
      ],
    },
    {
      label: "Buttons",
      buttons: [
        <AddButton key={"add"} value={"https://trueblocks.io"} />,
        <CleanButton key={"clean"} value={"https://trueblocks.io"} />,
      ],
    },
    {
      label: "Available Monitors",
      collapsable: false,
      components: [<DataTable<types.Monitor> key={"dataTable"} table={table} loading={false} />],
    },
  ];
};
