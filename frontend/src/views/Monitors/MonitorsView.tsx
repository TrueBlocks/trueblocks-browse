// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, TabItem, ViewForm, DebugState } from "@components";
import { ModifyMonitor } from "@gocode/app/App";
import { Page } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { MonitorsTableDef, MonitorsFormDef } from ".";
// EXISTING_CODE

export const MonitorsView = () => {
  const { monitors, fetchMonitors, loadAddress } = useAppState();
  const handleEnter = (page: Page) => {
    loadAddress(monitors.items[page.getRecord()].address);
  };
  const handleModify = ModifyMonitor;

  const table = useReactTable({
    data: monitors?.items || [],
    columns: MonitorsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const tabItems: ViewForm = {
    monitors: <TabItem tabName="monitors" data={monitors} groups={MonitorsFormDef(table)} />,
  };

  return (
    <ViewStateProvider
      // do not remove - delint
      nItems={monitors.nItems}
      fetchFn={fetchMonitors}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState u={[monitors.updater]} />
      <View tabItems={tabItems} searchable />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
