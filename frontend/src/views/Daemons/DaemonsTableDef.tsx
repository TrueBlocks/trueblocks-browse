// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef } from "@components";
import { types } from "@gocode/models";
// EXISTING_CODE

const columnHelper = createColumnHelper<types.Nothing>();

export const DaemonsTableDef: CustomColumnDef<types.Nothing, any>[] = [
  // EXISTING_CODE
  // EXISTING_CODE
];

// EXISTING_CODE
// TODO BOGUS: The daemon state should be in the AppState
// const [scraper] = useState<types.Daemon>(empty);
// const [freshen] = useState<types.Daemon>(empty);
// const [ipfs] = useState<types.Daemon>(empty);
// const [logMessages] = useState<messages.MessageMsg[]>([]);
// // const [scraper, setScraper] = useState<types.Daemon>(empty);
// const [freshen, setFreshen] = useState<types.Daemon>(empty);
// const [ipfs, setIpfs] = useState<types.Daemon>(empty);
// const [logMessages, setLogMessages] = useState<messages.MessageMsg[]>([]);

// const updateDaemon = (daemon: string, setDaemon: Dispatch<SetStateAction<types.Daemon>>) => {
//   GetDaemon(daemon).then((json: string) => {
//     setDaemon(types.Daemon.createFrom(json));
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

// const toggleDaemon = (name: string) => {
//   ToggleDaemon(name);
// };

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
