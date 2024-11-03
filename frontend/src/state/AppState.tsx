import { useState, useEffect, useContext, createContext, ReactNode } from "react";
import {
  ProjectPage,
  HistoryPage,
  MonitorPage,
  NamePage,
  AbiPage,
  IndexPage,
  ManifestPage,
  SettingsPage,
  SetChain,
  StatusPage,
  SessionPage,
  GetAppInfo,
} from "@gocode/app/App";
import { app, base, messages, types } from "@gocode/models";
import { EventsOff, EventsOn } from "@runtime";

interface AppStateProps {
  project: types.ProjectContainer;
  fetchProject: (currentItem: number, itemsPerPage: number) => void;

  history: types.HistoryContainer;
  fetchHistory: (currentItem: number, itemsPerPage: number) => void;
  setHistory: React.Dispatch<React.SetStateAction<types.HistoryContainer>>;

  monitors: types.MonitorContainer;
  fetchMonitors: (currentItem: number, itemsPerPage: number) => void;

  names: types.NameContainer;
  fetchNames: (currentItem: number, itemsPerPage: number) => void;

  abis: types.AbiContainer;
  fetchAbis: (currentItem: number, itemsPerPage: number) => void;

  indexes: types.IndexContainer;
  fetchIndexes: (currentItem: number, itemsPerPage: number) => void;

  manifests: types.ManifestContainer;
  fetchManifests: (currentItem: number, itemsPerPage: number) => void;

  settings: types.SettingsGroup;
  fetchSettings: (currentItem: number, itemsPerPage: number) => void;

  status: types.StatusContainer;
  fetchStatus: (currentItem: number, itemsPerPage: number) => void;

  session: types.SessionContainer;
  fetchSession: (currentItem: number, itemsPerPage: number) => void;

  address: base.Address;
  setAddress: (address: base.Address) => void;

  info: app.AppInfo;
  chain: string;
  meta: types.MetaData;
  isConfigured: boolean;
  wizardState: types.WizState;
  selectChain: (newChain: string) => void;
  setMeta: (meta: types.MetaData) => void;
  setWizardState: (state: types.WizState) => void;
}

const AppState = createContext<AppStateProps | undefined>(undefined);

export const AppStateProvider: React.FC<{ children: ReactNode }> = ({ children }: { children: ReactNode }) => {
  const [project, setProject] = useState<types.ProjectContainer>({} as types.ProjectContainer);
  const [history, setHistory] = useState<types.HistoryContainer>({} as types.HistoryContainer);
  const [monitors, setMonitors] = useState<types.MonitorContainer>({} as types.MonitorContainer);
  const [names, setNames] = useState<types.NameContainer>({} as types.NameContainer);
  const [abis, setAbis] = useState<types.AbiContainer>({} as types.AbiContainer);
  const [indexes, setIndexes] = useState<types.IndexContainer>({} as types.IndexContainer);
  const [manifests, setManifests] = useState<types.ManifestContainer>({} as types.ManifestContainer);
  const [settings, setSettings] = useState<types.SettingsGroup>({} as types.SettingsGroup);
  const [status, setStatus] = useState<types.StatusContainer>({} as types.StatusContainer);
  const [session, setSession] = useState<types.SessionContainer>({} as types.SessionContainer);
  // TODO BOGUS: The daemon state should be in the AppState

  const [address, setAddress] = useState<base.Address>("0x0" as unknown as base.Address);

  const [chain, setChain] = useState<string>("mainnet");
  const [isConfigured, setIsConfigured] = useState<boolean>(false);
  const [wizardState, setWizardState] = useState<types.WizState>(types.WizState.WELCOME);
  const [meta, setMeta] = useState<types.MetaData>({} as types.MetaData);
  const [info, setInfo] = useState<app.AppInfo>({} as app.AppInfo);

  const fetchProject = async (currentItem: number, itemsPerPage: number) => {
    ProjectPage(currentItem, itemsPerPage).then((item: types.ProjectContainer) => {
      setProject(item);
    });
  };

  const fetchHistory = async (currentItem: number, itemsPerPage: number) => {
    HistoryPage(String(address), currentItem, itemsPerPage).then((item: types.HistoryContainer) => {
      setHistory(item);
    });
  };

  const fetchMonitors = async (currentItem: number, itemsPerPage: number) => {
    MonitorPage(currentItem, itemsPerPage).then((item: types.MonitorContainer) => {
      if (item) {
        setMonitors(item);
      }
    });
  };

  const fetchNames = async (currentItem: number, itemsPerPage: number) => {
    NamePage(currentItem, itemsPerPage).then((item: types.NameContainer) => {
      if (item) {
        setNames(item);
      }
    });
  };

  const fetchAbis = async (currentItem: number, itemsPerPage: number) => {
    AbiPage(currentItem, itemsPerPage).then((item: types.AbiContainer) => {
      if (item) {
        setAbis(item);
      }
    });
  };

  const fetchIndexes = async (currentItem: number, itemsPerPage: number) => {
    IndexPage(currentItem, itemsPerPage).then((item: types.IndexContainer) => {
      if (item) {
        setIndexes(item);
      }
    });
  };

  const fetchManifests = async (currentItem: number, itemsPerPage: number) => {
    ManifestPage(currentItem, itemsPerPage).then((item: types.ManifestContainer) => {
      if (item) {
        setManifests(item);
      }
    });
  };

  const fetchSettings = async (currentItem: number, itemsPerPage: number) => {
    SettingsPage(currentItem, itemsPerPage).then((item: types.SettingsGroup) => {
      if (item) {
        setSettings(item);
      }
    });
  };

  const fetchStatus = async (currentItem: number, itemsPerPage: number) => {
    StatusPage(currentItem, itemsPerPage).then((item: types.StatusContainer) => {
      if (item) {
        setStatus(item);
      }
    });
  };

  const fetchSession = async (currentItem: number, itemsPerPage: number) => {
    SessionPage(currentItem, itemsPerPage).then((item: types.SessionContainer) => {
      if (item) {
        setSession(item);
      }
    });
  };

  const fetchAppInfo = () => {
    GetAppInfo().then((info) => {
      setChain(info.chain);
      setMeta(info.meta);
      setWizardState(info.state);
      setIsConfigured(info.isConfigured);
      setInfo(info);
    });
  };

  const selectChain = (newChain: string) => {
    setChain(newChain);
    SetChain(newChain, address) // disables refresh
      .then(() => {})
      .catch((error) => {
        console.error("Error setting chain:", error);
      });
  };

  useEffect(() => {
    fetchHistory(0, 15);
    HistoryPage(address as unknown as string, 0, 15).then((item: types.HistoryContainer) => {
      setHistory(item);
    });
  }, []); // eslint-disable-line react-hooks/exhaustive-deps

  useEffect(() => {
    const handleRefresh = () => {
      fetchAppInfo();
      fetchStatus(0, 100);
    };
    handleRefresh(); // first load

    // when messaged
    const { Message } = messages;
    EventsOn(Message.DAEMON, handleRefresh);
    EventsOn(Message.WIZARD, handleRefresh);
    return () => {
      EventsOff(Message.DAEMON);
      EventsOff(Message.WIZARD);
    };
  }, []);

  const state = {
    project,
    fetchProject,
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
    settings,
    fetchSettings,
    status,
    fetchStatus,
    session,
    fetchSession,
    address,
    info,
    chain,
    meta,
    isConfigured,
    wizardState,
    setAddress,
    selectChain,
    setMeta,
    setWizardState,
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
