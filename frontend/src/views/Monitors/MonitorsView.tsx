import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm } from "@components";
import { GoToHistory, ModifyMonitors } from "@gocode/app/App";
import { Page } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { MonitorsTableDef, MonitorFormDef } from ".";

export const MonitorsView = () => {
  const { monitors, fetchMonitors } = useAppState();

  const handleEnter = (page: Page) => {
    const address = monitors.items[page.getRecord()].address;
    GoToHistory(address).then(() => {});
  };

  const table = useReactTable({
    data: monitors.items || [],
    columns: MonitorsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "monitors";
  const tabs = ["monitors"];
  const forms: ViewForm = {
    monitors: <FormTable data={monitors} groups={MonitorFormDef(table)} />,
  };
  return (
    <ViewStateProvider
      route={route}
      nItems={monitors.nItems}
      fetchFn={fetchMonitors}
      onEnter={handleEnter}
      modifyFn={ModifyMonitors}
    >
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};
