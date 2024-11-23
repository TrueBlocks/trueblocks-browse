// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { Box, SimpleGrid, Stack } from "@mantine/core";
import { DaemonCard, DaemonLog, FieldGroup, FieldsetWrapper, PinButton } from "@components";
import { daemons, messages, updater } from "@gocode/models";
export interface Nope {
  scraper: daemons.Daemon;
  freshen: daemons.Daemon;
  ipfs: daemons.Daemon;
  logMessages: messages.MessageMsg[];
  toggleDaemon: (name: string) => void;
  updater: updater.Updater;
}
// EXISTING_CODE

export const DaemonsFormDef = (data: Nope): FieldGroup<Nope>[] => {
  return [
    // EXISTING_CODE
    {
      label: "Daemons",
      collapsable: false,
      components: [
        <SimpleGrid key={"cards"} cols={2}>
          <DaemonCard daemon={data.scraper} toggle={data.toggleDaemon} />
          <DaemonCard daemon={data.freshen} toggle={data.toggleDaemon} />
          <DaemonCard daemon={data.ipfs} toggle={data.toggleDaemon} />
        </SimpleGrid>,
        <Stack key={"logs"}>
          <Box />
          <FieldsetWrapper legend="Logs">
            <DaemonLog logMessages={data.logMessages} />
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
