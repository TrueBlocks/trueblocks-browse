import { Stack, Group } from "@mantine/core";
import { EditButton, FieldGroup } from "@components";
import { configtypes, types } from "@gocode/models";

export const ConfigFormDef = (cfg: types.ConfigContainer): FieldGroup<types.ConfigContainer>[] => {
  /*
  	    version: configtypes.VersionGroup;
	    settings: configtypes.SettingsContainer;
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
