// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { Box, SimpleGrid, Stack, Text } from "@mantine/core";
import { Table } from "@tanstack/react-table";
// import { DaemonCard, DaemonLog, FieldGroup, FieldsetWrapper, PinButton } from "@components";
import { FieldGroup, FieldsetWrapper, PinButton } from "@components";
import { daemons, messages, types, updater } from "@gocode/models";
export interface Nope {
  scraper: daemons.Daemon;
  freshen: daemons.Daemon;
  ipfs: daemons.Daemon;
  logMessages: messages.MessageMsg[];
  toggleDaemon: (name: string) => void;
  updater: updater.Updater;
}
// EXISTING_CODE

export const DaemonsFormDef = (table: Table<types.Nothing>): FieldGroup<types.DaemonContainer>[] => {
  return [
    // EXISTING_CODE
    {
      label: "Daemons",
      collapsable: false,
      components: [
        <SimpleGrid key={"cards"} cols={2}>
          <Text>Daemons</Text>
          {/* <DaemonCard daemon={data.scraper} toggle={data.toggleDaemon} />
          <DaemonCard daemon={data.freshen} toggle={data.toggleDaemon} />
          <DaemonCard daemon={data.ipfs} toggle={data.toggleDaemon} /> */}
        </SimpleGrid>,
        <Stack key={"logs"}>
          <Box />
          <FieldsetWrapper legend="Logs">
            <Text>Logs</Text>
            {/* <DaemonLog logMessages={data.logMessages} /> */}
          </FieldsetWrapper>
        </Stack>,
      ],
    },
    {
      label: "Buttons",
      buttons: [<PinButton key={"pin"} value="https://trueblocks.io" />],
    },
    // EXISTING_CODE
  ];
};

// EXISTING_CODE
// EXISTING_CODE
