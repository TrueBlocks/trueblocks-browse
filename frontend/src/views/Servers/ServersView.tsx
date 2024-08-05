import React, { useState, useEffect, Dispatch, SetStateAction } from "react";
import { View, ViewStatus, ViewTitle } from "@components";
import { ServerCard, ServerLog } from "./";
import { servers, messages } from "@gocode/models";
import { GetServer, ToggleServer } from "@gocode/app/App";
import { Stack, Title, SimpleGrid, Fieldset } from "@mantine/core";
import { EventsOn, EventsOff } from "@runtime";
import classes from "@/App.module.css";

var empty = {} as servers.Server;

export function ServersView() {
  const [scraper, setScraper] = useState<servers.Server>(empty);
  const [fileServer, setFileServer] = useState<servers.Server>(empty);
  const [freshen, setFreshen] = useState<servers.Server>(empty);
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
    updateServer("freshen", setFreshen);
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
      case "freshen":
        updateServer("freshen", setFreshen);
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
      <ViewTitle />
      <Stack className={classes.mainContent}>
        <Fieldset legend={"Servers"} bg={"white"} style={{ padding: "1rem", margin: "1rem" }}>
          <SimpleGrid cols={2} spacing="lg" style={{ padding: "lg" }}>
            <ServerCard server={scraper} toggle={toggleServer} />
            <ServerCard server={freshen} toggle={toggleServer} />
            <ServerCard server={ipfs} toggle={toggleServer} />
            <ServerCard server={fileServer} toggle={toggleServer} />
          </SimpleGrid>
          <ServerLog logMessages={logMessages} />
        </Fieldset>
      </Stack>
      <ViewStatus />
    </View>
  );
}
