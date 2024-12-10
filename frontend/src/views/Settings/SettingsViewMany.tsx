// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { useCallback } from "react";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { DebugState, TabItem, View, ViewForm } from "@components";
import { useNoops } from "@hooks";
import { ViewStateProvider, useAppState } from "@state";
import { ConfigFormDef, ConfigTableDef } from "../Config";
import { SessionFormDef, SessionTableDef } from "../Session";
import { StatusFormDef, StatusTableDef } from "../Status";
// EXISTING_CODE

export const SettingsView = () => {
  const { status, fetchStatus, config, fetchConfig, session, fetchSession } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  const fetchSettings = useCallback(
    (currentItem: number, itemsPerPage: number) => {
      fetchStatus(currentItem, itemsPerPage);
      fetchConfig(currentItem, itemsPerPage);
      fetchSession(currentItem, itemsPerPage);
    },
    [fetchStatus, fetchConfig, fetchSession]
  );

  const statusTable = useReactTable({
    data: status?.items || [],
    columns: StatusTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const configTable = useReactTable({
    data: config?.items || [],
    columns: ConfigTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const sessionTable = useReactTable({
    data: session?.items || [],
    columns: SessionTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const tabItems: ViewForm = {
    status: <TabItem tabName="status" data={status} groups={StatusFormDef(statusTable)} />,
    config: <TabItem tabName="config" data={config} groups={ConfigFormDef(configTable)} />,
    session: <TabItem tabName="session" data={session} groups={SessionFormDef(sessionTable)} />,
  };

  // if (!(status?.items?.length > 0)) {
  //   return <>Loading...</>;
  // }

  return (
    <ViewStateProvider
      // do not remove - delint
      nItems={status.nItems}
      fetchFn={fetchSettings}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState u={[status.updater, config.updater, session.updater]} />
      <View tabItems={tabItems} />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
