import React, { createContext, useContext, useState, ReactNode, useEffect } from "react";
import { types, messages, base, wizard, daemons } from "@gocode/models";
import { useKeyboardPaging } from "@hooks";
import { Pager, EmptyPager } from "@components";
import { EventsOn, EventsOff } from "@runtime";
import { Route } from "@/Routes";
import {
  HistoryPage,
  MonitorPage,
  NamePage,
  AbiPage,
  IndexPage,
  ManifestPage,
  StatusPage,
  GetLastSub,
  StepWizard,
  GetWizardState,
  GetMeta,
} from "@gocode/app/App";

type Counters = {
  nTxs: number;
  nMonitors: number;
  nNames: number;
  nAbis: number;
  nIndexes: number;
  nManifests: number;
  nStatus: number;
};

interface AppStateProps {
  address: base.Address;
  history: types.TransactionContainer;
  monitors: types.MonitorContainer;
  names: types.NameContainer;
  abis: types.AbiContainer;
  indexes: types.IndexContainer;
  manifests: types.ManifestContainer;
  status: types.StatusContainer;

  setAddress: (address: base.Address) => void;
  setHistory: (history: types.TransactionContainer) => void;
  setMonitors: (monitors: types.MonitorContainer) => void;
  setNames: (names: types.NameContainer) => void;
  setAbis: (abis: types.AbiContainer) => void;
  setIndexes: (indexes: types.IndexContainer) => void;
  setManifests: (manifests: types.ManifestContainer) => void;
  setStatus: (status: types.StatusContainer) => void;

  getPager: (name: Route) => Pager;
  resetPager: (name: Route) => void;

  isConfigured: boolean;
  wizardState: wizard.State;
  stepWizard: (step: wizard.Step) => void;

  meta: types.MetaData;
  setMeta: (meta: types.MetaData) => void;

  getCounters: () => Counters;
}

const AppState = createContext<AppStateProps | undefined>(undefined);

export const AppStateProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [meta, setMeta] = useState<types.MetaData>({} as types.MetaData);

  const [wizardState, setWizardState] = useState<wizard.State>(wizard.State.NOTOKAY);
  const [isConfigured, setIsConfigured] = useState<boolean>(false);

  const [address, setAddress] = useState<base.Address>("0x0" as unknown as base.Address);

  const [history, setHistory] = useState<types.TransactionContainer>({} as types.TransactionContainer);
  let historyPgr = useKeyboardPaging("history", history.nItems, [], 15);

  const [monitors, setMonitors] = useState<types.MonitorContainer>({} as types.MonitorContainer);
  let monitorPgr = useKeyboardPaging("monitors", monitors.nItems, [], 15);

  const [names, setNames] = useState<types.NameContainer>({} as types.NameContainer);
  let namesPgr = useKeyboardPaging("names", names.nItems, [], 15);

  const [abis, setAbis] = useState<types.AbiContainer>({} as types.AbiContainer);
  let abiPgr = useKeyboardPaging("abis", abis.nItems, [], 15);

  const [indexes, setIndexes] = useState<types.IndexContainer>({} as types.IndexContainer);
  let indexPgr = useKeyboardPaging("indexes", indexes.nItems, [], 15);

  const [manifests, setManifests] = useState<types.ManifestContainer>({} as types.ManifestContainer);
  let manifestPgr = useKeyboardPaging("manifest", manifests.nItems, [], 15);

  const [status, setStatus] = useState<types.StatusContainer>({} as types.StatusContainer);
  let statusPgr = useKeyboardPaging("status", status.nItems, [], 10);

  useEffect(() => {
    const fetchMeta = async () => {
      GetMeta().then((meta) => {
        setMeta(meta);
      });
    };
    fetchMeta();

    const fetchWizard = async () => {
      GetWizardState().then((state) => {
        setWizardState(state);
        setIsConfigured(state == wizard.State.OKAY);
      });
    };
    fetchWizard();

    const fetchHistory = async (address: base.Address, currentItem: number, itemsPerPage: number) => {
      GetLastSub("/history").then((subRoute: string) => {
        if (subRoute !== "") {
          subRoute = subRoute.replace("/", "");
          setAddress(subRoute as unknown as base.Address);
          HistoryPage(address as unknown as string, currentItem, itemsPerPage).then(
            (item: types.TransactionContainer) => {
              if (item) {
                setHistory(item);
              }
            }
          );
        }
      });
    };
    fetchHistory(address, historyPgr.curItem, historyPgr.perPage);

    const fetchMonitors = async (currentItem: number, itemsPerPage: number) => {
      MonitorPage(currentItem, itemsPerPage).then((item: types.MonitorContainer) => {
        if (item) {
          setMonitors(item);
        }
      });
    };
    fetchMonitors(monitorPgr.curItem, monitorPgr.perPage);

    const fetchNames = async (currentItem: number, itemsPerPage: number) => {
      NamePage(currentItem, itemsPerPage).then((item: types.NameContainer) => {
        if (item) {
          setNames(item);
        }
      });
    };
    fetchNames(namesPgr.curItem, namesPgr.perPage);

    const fetchAbis = async (currentItem: number, itemsPerPage: number) => {
      AbiPage(currentItem, itemsPerPage).then((item: types.AbiContainer) => {
        if (item) {
          setAbis(item);
        }
      });
    };
    fetchAbis(abiPgr.curItem, abiPgr.perPage);

    const fetchIndexes = async (currentItem: number, itemsPerPage: number) => {
      IndexPage(currentItem, itemsPerPage).then((item: types.IndexContainer) => {
        if (item) {
          setIndexes(item);
        }
      });
    };
    fetchIndexes(indexPgr.curItem, indexPgr.perPage);

    const fetchManifest = async (currentItem: number, itemsPerPage: number) => {
      ManifestPage(currentItem, itemsPerPage).then((item: types.ManifestContainer) => {
        if (item) {
          setManifests(item);
        }
      });
    };
    fetchManifest(manifestPgr.curItem, manifestPgr.perPage);

    const fetchStatus = async (currentItem: number, itemsPerPage: number) => {
      StatusPage(currentItem, itemsPerPage).then((item: types.StatusContainer) => {
        if (item) {
          setStatus(item);
        }
      });
    };
    fetchStatus(statusPgr.curItem, statusPgr.perPage);

    const handleRefresh = () => {
      fetchMeta();
      fetchWizard();
      // fetchHistory(address, historyPgr.curItem, historyPgr.perPage);
      fetchMonitors(monitorPgr.curItem, monitorPgr.perPage);
      fetchNames(namesPgr.curItem, namesPgr.perPage);
      fetchAbis(abiPgr.curItem, abiPgr.perPage);
      fetchIndexes(indexPgr.curItem, indexPgr.perPage);
      fetchManifest(manifestPgr.curItem, manifestPgr.perPage);
      fetchStatus(statusPgr.curItem, statusPgr.perPage);
    };

    var { Message } = messages;
    EventsOn(Message.DAEMON, handleRefresh);
    return () => {
      EventsOff(Message.DAEMON);
    };
  }, [
    // historyPgr.curItem,
    // historyPgr.perPage,
    monitorPgr.curItem,
    monitorPgr.perPage,
    namesPgr.curItem,
    namesPgr.perPage,
    abiPgr.curItem,
    abiPgr.perPage,
    indexPgr.curItem,
    indexPgr.perPage,
    manifestPgr.curItem,
    manifestPgr.perPage,
    statusPgr.curItem,
    statusPgr.perPage,
  ]);

  const stepWizard = (step: wizard.Step) => {
    StepWizard(step).then((state) => {
      setWizardState(state);
      setIsConfigured(state == wizard.State.OKAY);
    });
  };

  function getCounters(): Counters {
    return {
      nTxs: history.nItems,
      nMonitors: monitors.nItems,
      nNames: names.nItems,
      nAbis: abis.nItems,
      nIndexes: indexes.nItems,
      nManifests: manifests.nItems,
      nStatus: status.nItems,
    };
  }

  const getPager = (name: Route): Pager => {
    switch (name) {
      case "history":
        return historyPgr;
      case "monitors":
        return monitorPgr;
      case "names":
        return namesPgr;
      case "abis":
        return abiPgr;
      case "indexes":
        return indexPgr;
      case "manifest":
        return manifestPgr;
      case "status":
        return statusPgr;
      case "settings":
      case "daemons":
      case "":
      default:
        break;
    }
    return EmptyPager;
  };

  const resetPager = (name: Route) => {
    switch (name) {
      case "history":
        historyPgr = useKeyboardPaging("history", history.nItems, [], 15);
        break;
    }
  };

  let state = {
    address,
    history,
    monitors,
    names,
    abis,
    indexes,
    manifests,
    status,
    setAddress,
    setHistory,
    setMonitors,
    setNames,
    setAbis,
    setIndexes,
    setManifests,
    setStatus,
    getPager,
    resetPager,
    isConfigured,
    wizardState,
    stepWizard,
    meta,
    setMeta,
    getCounters,
  };

  return <AppState.Provider value={state}>{children}</AppState.Provider>;
};

export const useAppState = () => {
  const context = useContext(AppState);
  if (!context) {
    throw new Error("useAppState must be used within a AppStateProvider");
  }
  return context;
};
