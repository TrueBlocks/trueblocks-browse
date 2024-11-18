// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { useState, useEffect, Dispatch, SetStateAction } from "react";
import { SimpleGrid, Stack, Box } from "@mantine/core";
import { FieldGroup, FieldsetWrapper, FormTable, ViewForm, PinButton, View, DebugState } from "@components";
import { DaemonCard, DaemonLog } from "@components";
import { GetDaemon, ToggleDaemon } from "@gocode/app/App";
import { daemons, messages } from "@gocode/models";
import { useNoops } from "@hooks";
import { EventsOn, EventsOff } from "@runtime";
import { ViewStateProvider } from "@state";
// EXISTING_CODE

const empty = {} as daemons.Daemon;

interface Nope {
  scraper: daemons.Daemon;
  freshen: daemons.Daemon;
  ipfs: daemons.Daemon;
  logMessages: messages.MessageMsg[];
  toggleDaemon: (name: string) => void;
}

export const DaemonsView = () => {
  const { fetchNoop, enterNoop, modifyNoop } = useNoops();

  // EXISTING_CODE
  // TODO BOGUS: The daemon state should be in the AppState
  const [scraper, setScraper] = useState<daemons.Daemon>(empty);
  const [freshen, setFreshen] = useState<daemons.Daemon>(empty);
  const [ipfs, setIpfs] = useState<daemons.Daemon>(empty);
  const [logMessages, setLogMessages] = useState<messages.MessageMsg[]>([]);

  const updateDaemon = (daemon: string, setDaemon: Dispatch<SetStateAction<daemons.Daemon>>) => {
    GetDaemon(daemon).then((json: string) => {
      setDaemon(daemons.Daemon.createFrom(json));
    });
  };

  useEffect(() => {
    updateDaemon("scraper", setScraper);
    updateDaemon("freshen", setFreshen);
    updateDaemon("ipfs", setIpfs);
  }, []);

  const handleMessage = (msg: messages.MessageMsg) => {
    if (msg.num1 != 1) return; // ignore non-daemon refreshes here
    switch (msg.name) {
      case "scraper":
        updateDaemon("scraper", setScraper);
        break;
      case "freshen":
        updateDaemon("freshen", setFreshen);
        break;
      case "ipfs":
        updateDaemon("ipfs", setIpfs);
        break;
      default:
        break;
    }
    setLogMessages((prev) => {
      const newLogs = [...prev, msg];
      return newLogs.length > 8 ? newLogs.slice(-8) : newLogs;
    });
  };

  useEffect(() => {
    const { Message } = messages;
    EventsOn(Message.REFRESH, handleMessage);
    return () => {
      EventsOff(Message.REFRESH);
    };
  });

  const toggleDaemon = (name: string) => {
    ToggleDaemon(name);
  };

  const data: Nope = {
    toggleDaemon,
    scraper,
    freshen,
    ipfs,
    logMessages,
  };
  // EXISTING_CODE

  const route = "daemons";
  const tabs = ["daemons"];
  const forms: ViewForm = {
    daemons: <FormTable data={data} groups={createDaemonForm(data)} />,
  };

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={0}
      fetchFn={fetchNoop}
      onEnter={enterNoop}
      modifyFn={modifyNoop}
    >
      <DebugState n={0} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

const createDaemonForm = (data: Nope): FieldGroup<Nope>[] => {
  return [
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
  ];
};

// EXISTING_CODE
// EXISTING_CODE
