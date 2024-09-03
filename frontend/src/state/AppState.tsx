import React, { createContext, ReactNode, useContext, useEffect, useState } from "react";
import { Route } from "@/Routes";
import {
  AbiPage,
  GetLastSub,
  GetMeta,
  GetWizardState,
  HistoryPage,
  IndexPage,
  ManifestPage,
  MonitorPage,
  NamePage,
  PortfolioPage,
  StatusPage,
  StepWizard,
} from "@gocode/app/App";
import { base, messages, types, wizard } from "@gocode/models";
import { EventsOff, EventsOn } from "@runtime";

interface AppStateProps {
  address: base.Address;

  portfolio: types.PortfolioContainer;
  fetchPortfolio: (currentItem: number, itemsPerPage: number, item?: any) => void;

  history: types.HistoryContainer;
  fetchHistory: (currentItem: number, itemsPerPage: number, item?: any) => void;
  setHistory: React.Dispatch<React.SetStateAction<types.HistoryContainer>>;

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

  crudOperation(route: Route, selected: number, op: string): void;
}

const AppState = createContext<AppStateProps | undefined>(undefined);

export const AppStateProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [meta, setMeta] = useState<types.MetaData>({} as types.MetaData);

  const [wizardState, setWizardState] = useState<wizard.State>(wizard.State.NOTOKAY);
  const [isConfigured, setIsConfigured] = useState<boolean>(false);

  const [address, setAddress] = useState<base.Address>("0x0" as unknown as base.Address);

  const [portfolio, setPortfolio] = useState<types.PortfolioContainer>({} as types.PortfolioContainer);
  const [history, setHistory] = useState<types.HistoryContainer>({} as types.HistoryContainer);
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

  const fetchPortfolio = async (currentItem: number, itemsPerPage: number, item?: any) => {
    PortfolioPage(currentItem, itemsPerPage).then((item: types.PortfolioContainer) => {
      if (item) {
        setPortfolio(item);
      }
    });
  };

  const fetchHistory = async (currentItem: number, itemsPerPage: number, item?: any) => {
    HistoryPage(String(address), currentItem, itemsPerPage).then((item: types.HistoryContainer) => {
      if (item) {
        setHistory(item);
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

  const crudOperation = (route: Route, selected: number, op: string) => {
    console.log(
      "crudOperation",
      route,
      selected,
      op === "delete" ? "should delete" : op === "remove" ? "should remove" : "should undelete"
    );
    names.names[selected].deleted = op === "delete" ? true : op === "remove" ? false : false;
  };

  useEffect(() => {
    fetchMeta();
    fetchWizard();
    fetchStatus(0, 100);
  }, []);

  useEffect(() => {
    const handleRefresh = () => {
      fetchMeta();
      fetchWizard();
      fetchStatus(0, 100);
    };

    var { Message } = messages;
    EventsOn(Message.DAEMON, handleRefresh);
    return () => {
      EventsOff(Message.DAEMON);
    };
  }, []);

  useEffect(() => {
    fetchHistory(0, 15, null);
    HistoryPage(address as unknown as string, 0, 15).then((item: types.HistoryContainer) => {
      if (item) {
        setHistory(item);
      }
    });
  }, []);

  const stepWizard = (step: wizard.Step) => {
    StepWizard(step).then((state) => {
      setWizardState(state);
      setIsConfigured(state == wizard.State.OKAY);
    });
  };

  let state = {
    address,
    portfolio,
    fetchPortfolio,
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
    crudOperation,
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
