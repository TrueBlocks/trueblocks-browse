import { useState, useEffect, Dispatch, SetStateAction, ReactNode } from "react";
import { SimpleGrid, Stack, Box } from "@mantine/core";
import { FieldGroup, FieldsetWrapper, FormTable, PinButton, View } from "@components";
import { GetDaemon, ToggleDaemon } from "@gocode/app/App";
import { daemons, messages } from "@gocode/models";
import { useNoops } from "@hooks";
import { EventsOn, EventsOff } from "@runtime";
import { ViewStateProvider } from "@state";
import { DaemonCard, DaemonLog } from ".";

const empty = {} as daemons.Daemon;

interface Nope {
  scraper: daemons.Daemon;
  freshen: daemons.Daemon;
  ipfs: daemons.Daemon;
  logMessages: messages.MessageMsg[];
  toggleDaemon: (name: string) => void;
}

export const DaemonsView = () => {
  const { fetchNoop, modifyNoop } = useNoops();
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
    EventsOn(Message.DAEMON, handleMessage);
    return () => {
      EventsOff(Message.DAEMON);
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

  const route = "daemons";
  const tabs = ["daemons"];
  const forms: Record<string, ReactNode> = {
    daemons: <FormTable data={data} groups={createDaemonForm(data)} />,
  };

  return (
    <ViewStateProvider route={route} fetchFn={fetchNoop} modifyFn={modifyNoop}>
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
