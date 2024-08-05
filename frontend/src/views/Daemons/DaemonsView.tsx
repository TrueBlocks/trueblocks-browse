import React, { useState, useEffect, Dispatch, SetStateAction } from "react";
import { View, ViewStatus, ViewTitle } from "@components";
import { DaemonCard, DaemonLog } from ".";
import { daemons, messages } from "@gocode/models";
import { GetDaemon, ToggleDaemon } from "@gocode/app/App";
import { Stack, Title, SimpleGrid, Fieldset } from "@mantine/core";
import { EventsOn, EventsOff } from "@runtime";
import classes from "@/App.module.css";

var empty = {} as daemons.Daemon;

export function DaemonsView() {
  const [scraper, setScraper] = useState<daemons.Daemon>(empty);
  const [fileDaemon, setFileDaemon] = useState<daemons.Daemon>(empty);
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
    updateDaemon("filedaemon", setFileDaemon);
    updateDaemon("freshen", setFreshen);
    updateDaemon("ipfs", setIpfs);
  }, []);

  const handleMessage = (sMsg: messages.DaemonMsg) => {
    switch (sMsg.name) {
      case "scraper":
        updateDaemon("scraper", setScraper);
        break;
      case "filedaemon":
        updateDaemon("filedaemon", setFileDaemon);
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
    EventsOn("DAEMON", handleMessage);
    return () => {
      EventsOff("DAEMON");
    };
  }, []);

  const toggleDaemon = (name: string) => {
    ToggleDaemon(name);
  };

  return (
    <View>
      <Stack className={classes.mainContent}>
        <ViewTitle />
        <Fieldset legend={"Daemons"} bg={"white"}>
          <SimpleGrid cols={2}>
            <DaemonCard daemon={scraper} toggle={toggleDaemon} />
            <DaemonCard daemon={freshen} toggle={toggleDaemon} />
            <DaemonCard daemon={ipfs} toggle={toggleDaemon} />
            <DaemonCard daemon={fileDaemon} toggle={toggleDaemon} />
          </SimpleGrid>
        </Fieldset>
        <Fieldset legend={"Logs"} bg={"white"}>
          <DaemonLog logMessages={logMessages} />
        </Fieldset>
      </Stack>
      <ViewStatus />
    </View>
  );
}
