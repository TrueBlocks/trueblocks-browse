import { Text } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { FormTable, View, ViewForm } from "@components";
import { types } from "@gocode/models";
import { useNoops, useRenderCounter } from "@hooks";
import { ViewStateProvider, useAppState } from "@state";
import { ConfigFormDef } from "../Config";
import { SessionFormDef } from "../Session";
import { StatusFormDef } from "../Status";
import { SettingsTableDef } from ".";

export const SettingsView = () => {
  const { modifyNoop } = useNoops();
  const { settings, fetchSettings } = useAppState();
  const renderCount = useRenderCounter();

  const status = settings.status ?? types.StatusContainer.createFrom({});
  const config = settings.config ?? types.ConfigContainer.createFrom({});
  const session = settings.session ?? types.SessionContainer.createFrom({});

  const table = useReactTable({
    data: status.items || [],
    columns: SettingsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "settings";
  const tabs = ["config", "status", "session"];
  const forms: ViewForm = {
    status: <FormTable data={status} groups={StatusFormDef(table)} />,
    config: <FormTable data={config} groups={ConfigFormDef(config)} />,
    session: <FormTable data={session} groups={SessionFormDef()} />,
  };

  if (!settings) {
    return <div>Loading...</div>;
  }

  return (
    <ViewStateProvider route={route} fetchFn={fetchSettings} modifyFn={modifyNoop}>
      <Text>Render count: {renderCount}</Text>
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};
