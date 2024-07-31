import React, { useState, useEffect, Dispatch, SetStateAction } from "react";
import { View, ViewStatus } from "@components";
import { ServerCard, ServerLog } from "./";
import { servers, messages } from "@gocode/models";
import { GetServer, ToggleServer } from "@gocode/app/App";
import { Stack, Title, SimpleGrid } from "@mantine/core";
import { EventsOn, EventsOff } from "@runtime";
import classes from "@/App.module.css";

var empty = {} as servers.Server;

export function ServersView() {
  const [scraper, setScraper] = useState<servers.Server>(empty);
  const [fileServer, setFileServer] = useState<servers.Server>(empty);
  const [monitor, setMonitor] = useState<servers.Server>(empty);
  const [ipfs, setIpfs] = useState<servers.Server>(empty);
  const [logMessages, setLogMessages] = useState<messages.ServerMsg[]>([]);

  const updateServer = (server: string, setStateFn: Dispatch<SetStateAction<servers.Server>>) => {
    GetServer(server).then((s) => {
      setStateFn(s);
    });
  };

  useEffect(() => {
    updateServer("scraper", setScraper);
    updateServer("fileserver", setFileServer);
    updateServer("monitor", setMonitor);
    updateServer("ipfs", setIpfs);
  }, []);

  const handleMessage = (sMsg: messages.ServerMsg) => {
    switch (sMsg.name) {
      case "scraper":
        updateServer("scraper", setScraper);
        break;
      case "fileserver":
        updateServer("fileserver", setFileServer);
        break;
      case "monitor":
        updateServer("monitor", setMonitor);
        break;
      case "ipfs":
        updateServer("ipfs", setIpfs);
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
    EventsOn("SERVER", handleMessage);
    return () => {
      EventsOff("SERVER");
    };
  }, []);

  const toggleServer = (name: string) => {
    ToggleServer(name);
  };

  return (
    <View>
      <Title order={3}>Servers View</Title>
      <Stack className={classes.mainContent}>
        <SimpleGrid cols={2} spacing="lg" style={{ padding: "lg" }}>
          <ServerCard server={scraper} toggle={toggleServer} />
          <ServerCard server={monitor} toggle={toggleServer} />
          <ServerCard server={ipfs} toggle={toggleServer} />
          <ServerCard server={fileServer} toggle={toggleServer} />
        </SimpleGrid>
        <ServerLog logMessages={logMessages} />
      </Stack>
      <ViewStatus />
    </View>
  );
}
