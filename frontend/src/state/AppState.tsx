import React, { createContext, useContext, useState, ReactNode, useEffect, useRef } from "react";
import { types, messages, base, wizard } from "@gocode/models";
import { EventsOn, EventsOff } from "@runtime";
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

interface AppStateProps {
  address: base.Address;
  history: types.TransactionContainer;
  fetchHistory: (currentItem: number, itemsPerPage: number, item?: any) => void;
  setHistory: React.Dispatch<React.SetStateAction<types.TransactionContainer>>;
  monitors: types.MonitorContainer;
  fetchMonitors: (currentItem: number, itemsPerPage: number, item?: any) => void;
  names: types.NameContainer;
  fetchNames: (currentItem: number, itemsPerPage: number, item?: any) => void;
  abis: types.AbiContainer;
  fetchAbis: (currentItem: number, itemsPerPage: number, item?: any) => void;
  indexes: types.IndexContainer;
  fetchIndexes: (currentItem: number, itemsPerPage: number, item?: any) => void;
  manifests: types.ManifestContainer;
  fetchManifests: (currentItem: number, itemsPerPage: number, item?: any) => void;
  status: types.StatusContainer;
  fetchStatus: (currentItem: number, itemsPerPage: number, item?: any) => void;

  setAddress: (address: base.Address) => void;

  isConfigured: boolean;
  wizardState: wizard.State;
  stepWizard: (step: wizard.Step) => void;

  meta: types.MetaData;
  setMeta: (meta: types.MetaData) => void;

  getCounters: () => Counters;
}

const AppState = createContext<AppStateProps | undefined>(undefined);

export const AppStateProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const myRef = useRef(false);

  const [meta, setMeta] = useState<types.MetaData>({} as types.MetaData);

  const [wizardState, setWizardState] = useState<wizard.State>(wizard.State.NOTOKAY);
  const [isConfigured, setIsConfigured] = useState<boolean>(false);

  const [address, setAddress] = useState<base.Address>("0x0" as unknown as base.Address);

  const [history, setHistory] = useState<types.TransactionContainer>({} as types.TransactionContainer);
  const [monitors, setMonitors] = useState<types.MonitorContainer>({} as types.MonitorContainer);
  const [names, setNames] = useState<types.NameContainer>({} as types.NameContainer);
  const [abis, setAbis] = useState<types.AbiContainer>({} as types.AbiContainer);
  const [indexes, setIndexes] = useState<types.IndexContainer>({} as types.IndexContainer);
  const [manifests, setManifests] = useState<types.ManifestContainer>({} as types.ManifestContainer);
  const [status, setStatus] = useState<types.StatusContainer>({} as types.StatusContainer);

  const fetchMeta = async () => {
    GetMeta().then((meta) => {
      setMeta(meta);
    });
  };

  const fetchWizard = async () => {
    GetWizardState().then((state) => {
      setWizardState(state);
      setIsConfigured(state == wizard.State.OKAY);
    });
  };

  const fetchHistory = async (currentItem: number, itemsPerPage: number, item?: any) => {
    myRef.current = true;
    GetLastSub("/history").then((subRoute: string) => {
      if (subRoute !== "") {
        console.log("subRoute-app1: ", subRoute);
        subRoute = subRoute.replace("/", "");
        console.log("subRoute-app2: ", subRoute);
        console.log("subRoute-app3: ", subRoute as unknown as base.Address);
        setAddress(subRoute as unknown as base.Address);
      }
    });
  };

  const fetchMonitors = async (currentItem: number, itemsPerPage: number, item?: any) => {
    MonitorPage(currentItem, itemsPerPage).then((item: types.MonitorContainer) => {
      if (item) {
        setMonitors(item);
      }
    });
  };

  const fetchNames = async (currentItem: number, itemsPerPage: number, item?: any) => {
    NamePage(currentItem, itemsPerPage).then((item: types.NameContainer) => {
      if (item) {
        setNames(item);
      }
    });
  };

  const fetchAbis = async (currentItem: number, itemsPerPage: number, item?: any) => {
    AbiPage(currentItem, itemsPerPage).then((item: types.AbiContainer) => {
      if (item) {
        setAbis(item);
      }
    });
  };

  const fetchIndexes = async (currentItem: number, itemsPerPage: number, item?: any) => {
    IndexPage(currentItem, itemsPerPage).then((item: types.IndexContainer) => {
      if (item) {
        setIndexes(item);
      }
    });
  };

  const fetchManifests = async (currentItem: number, itemsPerPage: number, item?: any) => {
    ManifestPage(currentItem, itemsPerPage).then((item: types.ManifestContainer) => {
      if (item) {
        setManifests(item);
      }
    });
  };

  const fetchStatus = async (currentItem: number, itemsPerPage: number, item?: any) => {
    StatusPage(currentItem, itemsPerPage).then((item: types.StatusContainer) => {
      if (item) {
        setStatus(item);
      }
    });
  };

  useEffect(() => {
    myRef.current = true;
    fetchMeta();
    fetchWizard();
    fetchStatus(0, 1);
  }, []);

  useEffect(() => {
    myRef.current = true;
    const handleRefresh = () => {
      fetchMeta();
      fetchWizard();
      fetchStatus(0, 1);
    };

    var { Message } = messages;
    EventsOn(Message.DAEMON, handleRefresh);
    return () => {
      EventsOff(Message.DAEMON);
    };
  }, []);

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

  let state = {
    address,
    history,
    fetchHistory,
    setHistory,
    monitors,
    fetchMonitors,
    names,
    fetchNames,
    abis,
    fetchAbis,
    indexes,
    fetchIndexes,
    manifests,
    fetchManifests,
    status,
    fetchStatus,
    setAddress,
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

type Counters = {
  nTxs: number;
  nMonitors: number;
  nNames: number;
  nAbis: number;
  nIndexes: number;
  nManifests: number;
  nStatus: number;
};
