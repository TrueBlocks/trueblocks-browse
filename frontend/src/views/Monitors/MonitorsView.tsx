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

  // EXISTING_CODE
  // EXISTING_CODE

  const table = useReactTable({
    data: monitors?.items || [],
    columns: MonitorsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "monitors";
  const tabs = ["monitors"];
  const forms: ViewForm = {
    monitors: <FormTable data={monitors} groups={MonitorsFormDef(table)} />,
  };

  // if (!(monitors?.items?.length > 0)) {
  //   return <>Loading...</>;
  // }

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={monitors.nItems}
      fetchFn={fetchMonitors}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState u={monitors.updater} />
      <View tabs={tabs} forms={forms} searchable />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
