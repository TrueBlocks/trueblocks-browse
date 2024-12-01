// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm, DebugState } from "@components";
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

  // eslint-disable-next-line prefer-const
  let customTabs: string[] = [];
  // eslint-disable-next-line prefer-const
  let customForms: Record<string, JSX.Element> = {};
  // EXISTING_CODE
  // EXISTING_CODE

  const table = useReactTable({
    data: monitors?.items || [],
    columns: MonitorsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "monitors";
  const tabs = ["monitors", ...(customTabs || [])];
  const forms: ViewForm = {
    monitors: <FormTable data={monitors} groups={MonitorsFormDef(table)} />,
    ...customForms,
  };

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={monitors.nItems}
      fetchFn={fetchMonitors}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState u={[monitors.updater]} />
      <View tabs={tabs} forms={forms} searchable />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
