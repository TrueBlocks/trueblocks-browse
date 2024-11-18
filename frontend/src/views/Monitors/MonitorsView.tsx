import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm, DebugState } from "@components";
import { ModifyMonitor } from "@gocode/app/App";
import { Page } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { MonitorsTableDef, MonitorsFormDef } from ".";

export const MonitorsView = () => {
  const { monitors, fetchMonitors, loadAddress } = useAppState();

  const handleEnter = (page: Page) => {
    loadAddress(monitors.items[page.getRecord()].address);
  };

  const table = useReactTable({
    data: monitors.items || [],
    columns: MonitorsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "monitors";
  const tabs = ["monitors"];
  const forms: ViewForm = {
    monitors: <FormTable data={monitors} groups={MonitorsFormDef(table)} />,
  };
  return (
    <ViewStateProvider
      route={route}
      nItems={monitors.nItems}
      fetchFn={fetchMonitors}
      onEnter={handleEnter}
      modifyFn={ModifyMonitor}
    >
      <DebugState n={monitors.lastUpdate} />
      <View tabs={tabs} forms={forms} searchable />
    </ViewStateProvider>
  );
};
