// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { Stack, Group } from "@mantine/core";
import { Table } from "@tanstack/react-table";
import { DataTable, EditButton, FieldGroup } from "@components";
import { configtypes, types } from "@gocode/models";
// EXISTING_CODE

export const ConfigFormDef = (
  table: Table<configtypes.ChainGroup>,
  config: configtypes.Config
): FieldGroup<types.ConfigContainer>[] => {
  return [
    // EXISTING_CODE
    {
      label: "Version",
      colSpan: 6,
      components: [<VG key={"version"} version={config.version} />],
    },
    {
      label: "Settings",
      colSpan: 6,
      components: [<SG key={"settings"} settings={config.settings} />],
    },
    {
      label: "Buttons",
      buttons: [<EditButton key="edit" value="https://trueblocks.io" />],
    },
    {
      label: "Configured Chains",
      collapsable: false,
      components: [<DataTable<configtypes.ChainGroup> key={"dataTable"} table={table} loading={false} />],
    },
    // EXISTING_CODE
  ];
};

// EXISTING_CODE
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

// EXISTING_CODE
