import { useState, useEffect, Dispatch, SetStateAction } from "react";
import { SimpleGrid, Stack, Box } from "@mantine/core";
import { FieldGroup, FieldsetWrapper, FormTable, View } from "@components";
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
  logMessages: messages.DaemonMsg[];
  toggleDaemon: (name: string) => void;
}

export const DaemonsView = () => {
  const { fetchNoop, modifyNoop } = useNoops();
  const [scraper, setScraper] = useState<daemons.Daemon>(empty);
  const [freshen, setFreshen] = useState<daemons.Daemon>(empty);
  const [ipfs, setIpfs] = useState<daemons.Daemon>(empty);
  const [logMessages, setLogMessages] = useState<messages.DaemonMsg[]>([]);

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

  const handleMessage = (sMsg: messages.DaemonMsg) => {
    switch (sMsg.name) {
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
      const newLogs = [...prev, sMsg];
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

  const route = "daemons";
  const data: Nope = {
    toggleDaemon,
    scraper,
    freshen,
    ipfs,
    logMessages,
  };

  return (
    <ViewStateProvider route={route} fetchFn={fetchNoop} modifyFn={modifyNoop}>
      <View>
        <FormTable data={data} groups={createDaemonForm(data)} />
      </View>
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
  ];
};
