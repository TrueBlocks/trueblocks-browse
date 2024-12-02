// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { DebugState, FormTable, View, ViewForm } from "@components";
import { types } from "@gocode/models";
import { useNoops } from "@hooks";
import { ViewStateProvider, useAppState } from "@state";
import { ConfigFormDef, ConfigTableDef } from "../Config";
import { SessionFormDef, SessionTableDef } from "../Session";
import { StatusFormDef } from "../Status";
import { SettingsTableDef } from ".";
// EXISTING_CODE

export const SettingsView = () => {
  const { settings, fetchSettings } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  // EXISTING_CODE
  const status = settings.status ?? types.StatusContainer.createFrom({});
  const config = settings.config ?? types.ConfigContainer.createFrom({});
  const session = settings.session ?? types.SessionContainer.createFrom({});

  const table = useReactTable({
    data: status?.items || [],
    columns: SettingsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const table2 = useReactTable({
    data: [], // config?.items || [],
    columns: ConfigTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const table3 = useReactTable({
    data: session?.items || [],
    columns: SessionTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "settings";
  const tabs = ["config", "status", "session"];
  const forms: ViewForm = {
    status: <FormTable data={status} groups={StatusFormDef(table)} />,
    config: <FormTable data={config} groups={ConfigFormDef(table2)} />,
    session: <FormTable data={session} groups={SessionFormDef(table3)} />,
  };

  // if (!(status?.items?.length > 0)) {
  //   return <>Loading...</>;
  // }
  // EXISTING_CODE

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={status.nItems}
      fetchFn={fetchSettings}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState u={[settings.updater]} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
