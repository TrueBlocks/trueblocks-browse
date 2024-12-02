// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { useState } from "react";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { FormTable, ViewForm, View, DebugState } from "@components";
import { ToggleDaemon } from "@gocode/app/App";
import { types, messages, updater } from "@gocode/models";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { DaemonsFormDef, Nope } from "./DaemonsFormDef";
import { DaemonsTableDef } from "./DaemonsTableDef";

const empty = {} as types.Daemon;

// EXISTING_CODE

export const DaemonsView = () => {
  const { daemons, fetchDaemons } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  // EXISTING_CODE
  // TODO BOGUS: The daemon state should be in the AppState
  const [scraper] = useState<types.Daemon>(empty);
  const [freshen] = useState<types.Daemon>(empty);
  const [ipfs] = useState<types.Daemon>(empty);
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

  // const upd = updater.Updater.createFrom({});
  // const daemons: Nope = {
  //   toggleDaemon,
  //   scraper,
  //   freshen,
  //   ipfs,
  //   logMessages,
  //   updater: upd,
  // };
  // EXISTING_CODE

  const table = useReactTable({
    data: daemons?.items || [],
    columns: DaemonsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "daemons";
  const tabs = ["daemons"];
  const forms: ViewForm = {
    daemons: <FormTable data={daemons} groups={DaemonsFormDef(table)} />,
  };

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={daemons.nItems}
      fetchFn={fetchDaemons}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState u={daemons.updater} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
