// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
import { useState, useEffect, useContext, useCallback, createContext, ReactNode } from "react";
import {
  ProjectPage,
  HistoryPage,
  MonitorPage,
  NamePage,
  AbiPage,
  IndexPage,
  ManifestPage,
  StatusPage,
  SettingsPage,
  DaemonPage,
  SessionPage,
  ConfigPage,
  WizardPage,
  GetAppInfo,
  Logger,
} from "@gocode/app/App";
import { app, base, messages, types } from "@gocode/models";
import { EventsOff, EventsOn } from "@runtime";

interface AppStateProps {
  project: types.ProjectContainer;
  fetchProject: (currentItem: number, itemsPerPage: number) => void;

  history: types.HistoryContainer;
  fetchHistory: (currentItem: number, itemsPerPage: number) => void;

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

  status: types.StatusContainer;
  fetchStatus: (currentItem: number, itemsPerPage: number) => void;

  settings: types.SettingsContainer;
  fetchSettings: (currentItem: number, itemsPerPage: number) => void;

  daemons: types.DaemonContainer;
  fetchDaemons: (currentItem: number, itemsPerPage: number) => void;

  session: types.SessionContainer;
  fetchSession: (currentItem: number, itemsPerPage: number) => void;

  config: types.ConfigContainer;
  fetchConfig: (currentItem: number, itemsPerPage: number) => void;

  wizard: types.WizardContainer;
  fetchWizard: (currentItem: number, itemsPerPage: number) => void;

  address: base.Address;
  setAddress: (address: base.Address) => void;

  info: app.AppInfo;
  chain: string;
  meta: types.MetaData;
  isConfigured: boolean;
  wizState: types.WizState;
  setHistory: React.Dispatch<React.SetStateAction<types.HistoryContainer>>;
  setMeta: (meta: types.MetaData) => void;
  setWizState: (state: types.WizState) => void;
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
  const [status, setStatus] = useState<types.StatusContainer>({} as types.StatusContainer);
  const [settings, setSettings] = useState<types.SettingsContainer>({} as types.SettingsContainer);
  const [daemons, setDaemons] = useState<types.DaemonContainer>({} as types.DaemonContainer);
  const [session, setSession] = useState<types.SessionContainer>({} as types.SessionContainer);
  const [config, setConfig] = useState<types.ConfigContainer>({} as types.ConfigContainer);
  const [wizard, setWizard] = useState<types.WizardContainer>({} as types.WizardContainer);

  const [address, setAddress] = useState<base.Address>("0x0" as unknown as base.Address);

  const [chain, setChain] = useState<string>("mainnet");
  const [isConfigured, setIsConfigured] = useState<boolean>(false);
  const [wizState, setWizState] = useState<types.WizState>(types.WizState.WELCOME);
  const [meta, setMeta] = useState<types.MetaData>({} as types.MetaData);
  const [info, setInfo] = useState<app.AppInfo>({} as app.AppInfo);

  const fetchProject = useCallback((currentItem: number, itemsPerPage: number) => {
    ProjectPage(currentItem, itemsPerPage).then((item: types.ProjectContainer) => {
      if (item) {
        setProject(item);
      }
    });
  }, []);

  const fetchHistory = useCallback((currentItem: number, itemsPerPage: number) => {
    HistoryPage(currentItem, itemsPerPage).then((item: types.HistoryContainer) => {
      if (item) {
        setHistory(item);
      }
    });
  }, []);

  const fetchMonitors = useCallback((currentItem: number, itemsPerPage: number) => {
    MonitorPage(currentItem, itemsPerPage).then((item: types.MonitorContainer) => {
      if (item) {
        setMonitors(item);
      }
    });
  }, []);

  const fetchNames = useCallback((currentItem: number, itemsPerPage: number) => {
    NamePage(currentItem, itemsPerPage).then((item: types.NameContainer) => {
      if (item) {
        setNames(item);
      }
    });
  }, []);

  const fetchAbis = useCallback((currentItem: number, itemsPerPage: number) => {
    AbiPage(currentItem, itemsPerPage).then((item: types.AbiContainer) => {
      if (item) {
        setAbis(item);
      }
    });
  }, []);

  const fetchIndexes = useCallback((currentItem: number, itemsPerPage: number) => {
    IndexPage(currentItem, itemsPerPage).then((item: types.IndexContainer) => {
      if (item) {
        setIndexes(item);
      }
    });
  }, []);

  const fetchManifests = useCallback((currentItem: number, itemsPerPage: number) => {
    ManifestPage(currentItem, itemsPerPage).then((item: types.ManifestContainer) => {
      if (item) {
        setManifests(item);
      }
    });
  }, []);

  const fetchStatus = useCallback((currentItem: number, itemsPerPage: number) => {
    StatusPage(currentItem, itemsPerPage).then((item: types.StatusContainer) => {
      if (item) {
        setStatus(item);
      }
    });
  }, []);

  const fetchSettings = useCallback((currentItem: number, itemsPerPage: number) => {
    SettingsPage(currentItem, itemsPerPage).then((item: types.SettingsContainer) => {
      if (item) {
        setSettings(item);
      }
    });
  }, []);

  const fetchDaemons = useCallback((currentItem: number, itemsPerPage: number) => {
    DaemonPage(currentItem, itemsPerPage).then((item: types.DaemonContainer) => {
      if (item) {
        setDaemons(item);
      }
    });
  }, []);

  const fetchSession = useCallback((currentItem: number, itemsPerPage: number) => {
    SessionPage(currentItem, itemsPerPage).then((item: types.SessionContainer) => {
      if (item) {
        setSession(item);
      }
    });
  }, []);

  const fetchConfig = useCallback((currentItem: number, itemsPerPage: number) => {
    ConfigPage(currentItem, itemsPerPage).then((item: types.ConfigContainer) => {
      if (item) {
        setConfig(item);
      }
    });
  }, []);

  const fetchWizard = useCallback((currentItem: number, itemsPerPage: number) => {
    Logger(["Calling into WizardPage"]);
    WizardPage(currentItem, itemsPerPage).then((item: types.WizardContainer) => {
      if (item) {
        Logger(["got items", JSON.stringify(item)]);
        setWizard(item);
      }
    });
  }, []);

  const fetchAppInfo = () => {
    GetAppInfo().then((info) => {
      setChain(info.chain);
      setMeta(info.meta);
      setWizState(info.state);
      setIsConfigured(info.isConfigured);
      setInfo(info);
    });
  };

  useEffect(() => {
    fetchHistory(0, 15);
    HistoryPage(0, 15).then((item: types.HistoryContainer) => {
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
    EventsOn(Message.REFRESH, handleRefresh);
    return () => {
      EventsOff(Message.REFRESH);
    };
  }, [fetchStatus]);

  const state = {
    project,
    fetchProject,
    history,
    fetchHistory,
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
    settings,
    fetchSettings,
    daemons,
    fetchDaemons,
    session,
    fetchSession,
    config,
    fetchConfig,
    wizard,
    fetchWizard,

    address,
    info,
    chain,
    meta,
    isConfigured,
    wizState,
    setHistory,
    setAddress,
    setMeta,
    setWizState,
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
