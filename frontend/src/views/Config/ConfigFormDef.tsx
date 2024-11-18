// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { Stack, Group } from "@mantine/core";
import { EditButton, FieldGroup } from "@components";
import { configtypes, types } from "@gocode/models";
// EXISTING_CODE

export const ConfigFormDef = (config: types.ConfigContainer): FieldGroup<types.ConfigContainer>[] => {
  return [
    // EXISTING_CODE
    {
      label: "Version",
      components: [<VG key={"version"} version={config.version} />, <SG key={"settings"} settings={config.settings} />],
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

//-------------------------------------------------------------------
// Template variables:
// class:         Config
// routeLabel:    Config
// itemName:
// isHistory:     false
// isSession:     false
// isConfig:      true
