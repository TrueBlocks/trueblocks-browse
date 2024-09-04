import React, { useState, useEffect, Dispatch, SetStateAction } from "react";
import { SimpleGrid, Fieldset } from "@mantine/core";
import { View } from "@components";
import { GetDaemon, ToggleDaemon } from "@gocode/app/App";
import { daemons, messages } from "@gocode/models";
import { EventsOn, EventsOff } from "@runtime";
import { ViewStateProvider } from "@state";
import { DaemonCard, DaemonLog } from ".";

const empty = {} as daemons.Daemon;

export function DaemonsView() {
  const [scraper, setScraper] = useState<daemons.Daemon>(empty);
  const [freshen, setFreshen] = useState<daemons.Daemon>(empty);
  const [ipfs, setIpfs] = useState<daemons.Daemon>(empty);
  const [logMessages, setLogMessages] = useState<messages.DaemonMsg[]>([]);

  const updateDaemon = (daemon: string, setStateFn: Dispatch<SetStateAction<daemons.Daemon>>) => {
    GetDaemon(daemon).then((s) => {
      setStateFn(s);
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
  }, []);

  const toggleDaemon = (name: string) => {
    ToggleDaemon(name);
  };

  return (
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    <ViewStateProvider route="daemons" fetchFn={(_unused1: number, _unused2: number) => {}}>
      <View>
        <Fieldset legend={"Daemons"} bg={"white"}>
          <SimpleGrid cols={2}>
            <DaemonCard daemon={scraper} toggle={toggleDaemon} />
            <DaemonCard daemon={freshen} toggle={toggleDaemon} />
            <DaemonCard daemon={ipfs} toggle={toggleDaemon} />
          </SimpleGrid>
        </Fieldset>
        <Fieldset legend={"Logs"} bg={"white"}>
          <DaemonLog logMessages={logMessages} />
        </Fieldset>
      </View>
    </ViewStateProvider>
  );
}
