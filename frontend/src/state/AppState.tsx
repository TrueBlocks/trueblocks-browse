import { createContext, ReactNode, useContext, useEffect, useState } from "react";
import {
  AbiPage,
  GetMeta,
  GetChain,
  SetChain,
  GetWizardState,
  HistoryPage,
  IndexPage,
  ManifestPage,
  MonitorPage,
  NamePage,
  ProjectPage,
  StatusPage,
  StepWizard,
} from "@gocode/app/App";
import { base, messages, types, wizard } from "@gocode/models";
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

  status: types.StatusContainer;
  fetchStatus: (currentItem: number, itemsPerPage: number) => void;

  address: base.Address;
  setAddress: (address: base.Address) => void;

  chain: string;
  changeChain: (newChain: string) => void;

  meta: types.MetaData;
  setMeta: (meta: types.MetaData) => void;

  isConfigured: boolean;
  wizardState: wizard.State;
  stepWizard: (step: wizard.Step) => void;
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

  const [address, setAddress] = useState<base.Address>("0x0" as unknown as base.Address);
  const [chain, setChain] = useState<string>("");
  const [meta, setMeta] = useState<types.MetaData>({} as types.MetaData);

  const [wizardState, setWizardState] = useState<wizard.State>(wizard.State.NOTOKAY);
  const [isConfigured, setIsConfigured] = useState<boolean>(false);

  const fetchProject = async (currentItem: number, itemsPerPage: number) => {
    ProjectPage(currentItem, itemsPerPage).then((item: types.ProjectContainer) => {
      if (item) {
        setProject(item);
      }
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

  const fetchStatus = async (currentItem: number, itemsPerPage: number) => {
    StatusPage(currentItem, itemsPerPage).then((item: types.StatusContainer) => {
      if (item) {
        setStatus(item);
      }
    });
  };

  const fetchChain = async () => {
    GetChain().then((chain) => {
      setChain(chain);
    });
  };

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

  const changeChain = (newChain: string) => {
    setChain(newChain);
    SetChain(newChain, address) // disables refresh
      .then(() => {})
      .catch((error) => {
        console.error("Error setting chain:", error);
      });
  };

  useEffect(() => {
    fetchChain();
    fetchMeta();
    fetchWizard();
    fetchStatus(0, 100);
  }, []);

  useEffect(() => {
    const handleRefresh = () => {
      fetchChain();
      fetchMeta();
      fetchWizard();
      fetchStatus(0, 100);
    };

    const { Message } = messages;
    EventsOn(Message.DAEMON, handleRefresh);
    return () => {
      EventsOff(Message.DAEMON);
    };
  }, []);

  useEffect(() => {
    fetchHistory(0, 15);
    HistoryPage(address as unknown as string, 0, 15).then((item: types.HistoryContainer) => {
      setHistory(item);
    });
  }, []); // eslint-disable-line react-hooks/exhaustive-deps

  const stepWizard = (step: wizard.Step) => {
    StepWizard(step).then((state) => {
      setWizardState(state);
      setIsConfigured(state == wizard.State.OKAY);
    });
  };

  const state = {
    address,
    chain,
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
    status,
    fetchStatus,
    setAddress,
    changeChain,
    isConfigured,
    wizardState,
    stepWizard,
    meta,
    setMeta,
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
