// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { useState } from "react";
import { FormTable, ViewForm, View, DebugState } from "@components";
import { ToggleDaemon } from "@gocode/app/App";
import { daemons, messages } from "@gocode/models";
import { useNoops } from "@hooks";
import { ViewStateProvider } from "@state";
import { DaemonsFormDef, Nope } from "./DaemonsFormDef";

const empty = {} as daemons.Daemon;

// EXISTING_CODE

export const DaemonsView = () => {
  const { fetchNoop, enterNoop, modifyNoop } = useNoops();
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  // EXISTING_CODE
  // TODO BOGUS: The daemon state should be in the AppState
  const [scraper] = useState<daemons.Daemon>(empty);
  const [freshen] = useState<daemons.Daemon>(empty);
  const [ipfs] = useState<daemons.Daemon>(empty);
  const [logMessages] = useState<messages.MessageMsg[]>([]);
  // const [scraper, setScraper] = useState<daemons.Daemon>(empty);
  // const [freshen, setFreshen] = useState<daemons.Daemon>(empty);
  // const [ipfs, setIpfs] = useState<daemons.Daemon>(empty);
  // const [logMessages, setLogMessages] = useState<messages.MessageMsg[]>([]);

  // const updateDaemon = (daemon: string, setDaemon: Dispatch<SetStateAction<daemons.Daemon>>) => {
  //   GetDaemon(daemon).then((json: string) => {
  //     setDaemon(daemons.Daemon.createFrom(json));
  //   });
  // };

  // useEffect(() => {
  //   updateDaemon("scraper", setScraper);
  //   updateDaemon("freshen", setFreshen);
  //   updateDaemon("ipfs", setIpfs);
  // }, []);

  // const handleMessage = (msg: messages.MessageMsg) => {
  //   if (msg.num1 != 1) return; // ignore non-daemon refreshes here
  //   switch (msg.name) {
  //     case "scraper":
  //       updateDaemon("scraper", setScraper);
  //       break;
  //     case "freshen":
  //       updateDaemon("freshen", setFreshen);
  //       break;
  //     case "ipfs":
  //       updateDaemon("ipfs", setIpfs);
  //       break;
  //     default:
  //       break;
  //   }
  //   setLogMessages((prev) => {
  //     const newLogs = [...prev, msg];
  //     return newLogs.length > 8 ? newLogs.slice(-8) : newLogs;
  //   });
  // };

  // useEffect(() => {
  //   const { Message } = messages;
  //   EventsOn(Message.REFRESH, handleMessage);
  //   return () => {
  //     EventsOff(Message.REFRESH);
  //   };
  // });

  const toggleDaemon = (name: string) => {
    ToggleDaemon(name);
  };

  const daemons: Nope = {
    toggleDaemon,
    scraper,
    freshen,
    ipfs,
    logMessages,
  };
  const table = daemons;
  const fetchDaemons = fetchNoop;
  // EXISTING_CODE

  const route = "daemons";
  const tabs = ["daemons"];
  const forms: ViewForm = {
    daemons: <FormTable data={daemons} groups={DaemonsFormDef(table)} />,
  };

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={0}
      fetchFn={fetchDaemons}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState n={0} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
