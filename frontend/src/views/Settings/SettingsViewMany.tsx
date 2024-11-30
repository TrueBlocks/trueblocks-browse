// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { useCallback } from "react";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { DebugState, FormTable, View, ViewForm } from "@components";
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

  // EXISTING_CODE
  // EXISTING_CODE

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

  const route = "settings";
  const tabs = ["status", "config", "session"];
  const forms: ViewForm = {
    status: <FormTable data={status} groups={StatusFormDef(statusTable)} />,
    config: <FormTable data={config} groups={ConfigFormDef(configTable)} />,
    session: <FormTable data={session} groups={SessionFormDef(sessionTable)} />,
  };

  // if (!(status?.items?.length > 0)) {
  //   return <>Loading...</>;
  // }

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={status.nItems}
      fetchFn={fetchSettings}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState u={[status.updater, config.updater, session.updater]} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
