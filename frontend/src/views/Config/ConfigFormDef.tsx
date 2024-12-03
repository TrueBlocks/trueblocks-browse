// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { Stack, Group } from "@mantine/core";
import { Table } from "@tanstack/react-table";
import { DataTable, EditButton, FieldGroup } from "@components";
import { types } from "@gocode/models";
import { useAppState } from "../../state";
// EXISTING_CODE

export const ConfigFormDef = (table: Table<types.Chain>): FieldGroup<types.ConfigContainer>[] => {
  // EXISTING_CODE
  // EXISTING_CODE
  return [
    // EXISTING_CODE
    {
      label: "Version",
      colSpan: 6,
      components: [<VG key={"version"} />],
    },
    {
      label: "Settings",
      colSpan: 6,
      components: [<SG key={"settings"} />],
    },
    {
      label: "Buttons",
      buttons: [<EditButton key="edit" value="https://trueblocks.io" />],
    },
    {
      label: "Configured Chains",
      collapsable: false,
      components: [<DataTable<types.Chain> key={"dataTable"} table={table} loading={false} />],
    },
    // EXISTING_CODE
  ];
};

// EXISTING_CODE
const VG = () => {
  const { config } = useAppState();
  return (
    <Stack>
      <Group>
        <div>current:</div>
        <div>{config.version?.current}</div>
      </Group>
    </Stack>
  );
};

const SG = () => {
  const { config } = useAppState();
  return (
    <Stack>
      <Group>
        <div>cachePath:</div>
        <div>{config.settings?.cachePath}</div>
      </Group>
      <Group>
        <div>indexPath:</div>
        <div>{config.settings?.indexPath}</div>
      </Group>
      <Group>
        <div>defaultChain:</div>
        <div>{config.settings?.defaultChain}</div>
      </Group>
      <Group>
        <div>defaultGateway:</div>
        <div>{config.settings?.defaultGateway}</div>
      </Group>
    </Stack>
  );
};

// EXISTING_CODE
