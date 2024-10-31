import { Stack, Group } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { CleanButton, DataTable, EditButton, FieldGroup, FormTable, View, ViewForm } from "@components";
import { configtypes, types } from "@gocode/models";
import { useNoops } from "@hooks";
import { ViewStateProvider, useAppState } from "@state";
import { SessionFormDef } from "../Session";
import { tableColumns } from ".";

export const SettingsView = () => {
  const { modifyNoop } = useNoops();
  const { settings, fetchSettings } = useAppState();

  const status = settings.status ?? types.StatusContainer.createFrom({});
  const config = settings.config ?? types.ConfigContainer.createFrom({});
  const session = settings.session ?? types.SessionContainer.createFrom({});

  const table = useReactTable({
    data: status.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "settings";
  const tabs = ["config", "status", "session"];
  const forms: ViewForm = {
    status: <FormTable data={status} groups={createStatusForm(table)} />,
    config: <FormTable data={config} groups={createConfigForm(config)} />,
    session: <FormTable data={session} groups={SessionFormDef()} />,
  };

  if (!settings) {
    return <div>Loading...</div>;
  }

  return (
    <ViewStateProvider route={route} fetchFn={fetchSettings} modifyFn={modifyNoop}>
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

const createConfigForm = (cfg: types.ConfigContainer): FieldGroup<types.ConfigContainer>[] => {
  /*
  	    version: configtypes.VersionGroup;
	    settings: configtypes.SettingsGroup;
	    keys: {[key: string]: configtypes.KeyGroup};
	    pinning: configtypes.PinningGroup;
	    unchained: configtypes.UnchainedGroup;
	    chains: {[key: string]: configtypes.ChainGroup};
	    // Go type: time
	    lastUpdate: any;
 */
  // const vg = <VG key={"version"} version={cfg.version} />;
  return [
    {
      label: "Version",
      components: [<VG key={"version"} version={cfg.version} />, <SG key={"settings"} settings={cfg.settings} />],
    },
    {
      label: "Buttons",
      buttons: [<EditButton key="edit" value="https://trueblocks.io" />],
    },
    // {
    //   label: "Version Group",
    //   colSpan: 12, // Full width since it's the only field
    //   components: [<div key={"1"}>Hello world</div>],
    // },
  ];
};

type ConfigProps = {
  version?: configtypes.VersionGroup;
  settings?: configtypes.SettingsGroup;
};

const VG = ({ version }: ConfigProps) => {
  return (
    <Stack>
      <Group>
        <div>current:</div>
        <div>{version?.current}</div>
      </Group>
    </Stack>
  );
};

const SG = ({ settings }: ConfigProps) => {
  return (
    <Stack>
      <Group>
        <div>cachePath:</div>
        <div>{settings?.cachePath}</div>
      </Group>
      <Group>
        <div>indexPath:</div>
        <div>{settings?.indexPath}</div>
      </Group>
      <Group>
        <div>defaultChain:</div>
        <div>{settings?.defaultChain}</div>
      </Group>
      <Group>
        <div>defaultGateway:</div>
        <div>{settings?.defaultGateway}</div>
      </Group>
    </Stack>
  );
};

const createStatusForm = (table: any): FieldGroup<types.StatusContainer>[] => {
  return [
    {
      label: "System Data",
      colSpan: 7,
      fields: [
        { label: "trueblocks", type: "text", accessor: "version" },
        { label: "client", type: "text", accessor: "clientVersion" },
        { label: "isArchive", type: "boolean", accessor: "isArchive" },
        { label: "isTracing", type: "boolean", accessor: "isTracing" },
      ],
    },
    {
      label: "API Keys",
      colSpan: 5,
      fields: [
        { label: "hasEsKey", type: "boolean", accessor: "hasEsKey" },
        { label: "hasPinKey", type: "boolean", accessor: "hasPinKey" },
        { label: "rpcProvider", type: "url", accessor: "rpcProvider" },
      ],
    },
    {
      label: "Configuration Paths",
      colSpan: 7,
      fields: [
        { label: "rootConfig", type: "path", accessor: "rootConfig" },
        { label: "chainConfig", type: "path", accessor: "chainConfig" },
        { label: "indexPath", type: "path", accessor: "indexPath" },
        { label: "cachePath", type: "path", accessor: "cachePath" },
      ],
    },
    {
      label: "Statistics",
      colSpan: 5,
      fields: [
        { label: "lastUpdate", type: "date", accessor: "lastUpdate" },
        { label: "nCaches", type: "int", accessor: "nItems" },
        { label: "nFiles", type: "int", accessor: "nFiles" },
        { label: "nFolders", type: "int", accessor: "nFolders" },
        { label: "sizeInBytes", type: "bytes", accessor: "nBytes" },
      ],
    },
    {
      label: "Buttons",
      buttons: [<CleanButton key={"clean"} value={"https://trueblocks.io"} />],
    },
    {
      label: "Caches",
      fields: [],
      collapsable: false,
      components: [<DataTable<types.CacheItem> key={"dataTable"} table={table} loading={false} />],
    },
  ];
};
