// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
import { useState, useEffect, useContext, useCallback, createContext, ReactNode, useRef } from "react";
import { useLocation } from "wouter";
import {
  FetchProject,
  FetchBalance,
  FetchIncoming,
  FetchOutgoing,
  FetchInternal,
  FetchChart,
  FetchLog,
  FetchStatement,
  FetchNeighbor,
  FetchTrace,
  FetchReceipt,
  FetchMonitor,
  FetchName,
  FetchAbi,
  FetchIndex,
  FetchManifest,
  FetchPin,
  FetchUpload,
  FetchStatus,
  FetchDaemon,
  FetchSession,
  FetchConfig,
  FetchWizard,
  FetchAppInfo,
  LoadAddress,
  GetLastRoute,
  SetLastRoute,
  GetLastTab,
  SetLastTab,
} from "@gocode/app/App";
import { app, base, messages, types } from "@gocode/models";
import { EventsOff, EventsOn } from "@runtime";

interface AppStateProps {
  project: types.ProjectContainer;
  fetchProject: (currentItem: number, itemsPerPage: number) => void;
  balances: types.BalanceContainer;
  fetchBalances: (currentItem: number, itemsPerPage: number) => void;
  incoming: types.IncomingContainer;
  fetchIncoming: (currentItem: number, itemsPerPage: number) => void;
  outgoing: types.OutgoingContainer;
  fetchOutgoing: (currentItem: number, itemsPerPage: number) => void;
  internals: types.InternalContainer;
  fetchInternals: (currentItem: number, itemsPerPage: number) => void;
  charts: types.ChartContainer;
  fetchCharts: (currentItem: number, itemsPerPage: number) => void;
  logs: types.LogContainer;
  fetchLogs: (currentItem: number, itemsPerPage: number) => void;
  statements: types.StatementContainer;
  fetchStatements: (currentItem: number, itemsPerPage: number) => void;
  neighbors: types.NeighborContainer;
  fetchNeighbors: (currentItem: number, itemsPerPage: number) => void;
  traces: types.TraceContainer;
  fetchTraces: (currentItem: number, itemsPerPage: number) => void;
  receipts: types.ReceiptContainer;
  fetchReceipts: (currentItem: number, itemsPerPage: number) => void;
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
  pins: types.PinContainer;
  fetchPins: (currentItem: number, itemsPerPage: number) => void;
  uploads: types.UploadContainer;
  fetchUploads: (currentItem: number, itemsPerPage: number) => void;
  status: types.StatusContainer;
  fetchStatus: (currentItem: number, itemsPerPage: number) => void;
  daemons: types.DaemonContainer;
  fetchDaemons: (currentItem: number, itemsPerPage: number) => void;
  session: types.SessionContainer;
  fetchSession: (currentItem: number, itemsPerPage: number) => void;
  config: types.ConfigContainer;
  fetchConfig: (currentItem: number, itemsPerPage: number) => void;
  wizard: types.WizardContainer;
  fetchWizard: (currentItem: number, itemsPerPage: number) => void;
  info: app.AppInfo;
  loadAddress: (address: base.Address) => void;
  counters: React.MutableRefObject<Record<string, number>>;
  route: string;
  routeChanged: (newVal: string) => void;
  activeTab: string;
  tabChanged: (newVal: string) => void;
}

const AppState = createContext<AppStateProps | undefined>(undefined);

export const AppStateProvider: React.FC<{ children: ReactNode }> = ({ children }: { children: ReactNode }) => {
  const [project, setProject] = useState<types.ProjectContainer>({} as types.ProjectContainer);
  const [balances, setBalances] = useState<types.BalanceContainer>({} as types.BalanceContainer);
  const [incoming, setIncoming] = useState<types.IncomingContainer>({} as types.IncomingContainer);
  const [outgoing, setOutgoing] = useState<types.OutgoingContainer>({} as types.OutgoingContainer);
  const [internals, setInternals] = useState<types.InternalContainer>({} as types.InternalContainer);
  const [charts, setCharts] = useState<types.ChartContainer>({} as types.ChartContainer);
  const [logs, setLogs] = useState<types.LogContainer>({} as types.LogContainer);
  const [statements, setStatements] = useState<types.StatementContainer>({} as types.StatementContainer);
  const [neighbors, setNeighbors] = useState<types.NeighborContainer>({} as types.NeighborContainer);
  const [traces, setTraces] = useState<types.TraceContainer>({} as types.TraceContainer);
  const [receipts, setReceipts] = useState<types.ReceiptContainer>({} as types.ReceiptContainer);
  const [monitors, setMonitors] = useState<types.MonitorContainer>({} as types.MonitorContainer);
  const [names, setNames] = useState<types.NameContainer>({} as types.NameContainer);
  const [abis, setAbis] = useState<types.AbiContainer>({} as types.AbiContainer);
  const [indexes, setIndexes] = useState<types.IndexContainer>({} as types.IndexContainer);
  const [manifests, setManifests] = useState<types.ManifestContainer>({} as types.ManifestContainer);
  const [pins, setPins] = useState<types.PinContainer>({} as types.PinContainer);
  const [uploads, setUploads] = useState<types.UploadContainer>({} as types.UploadContainer);
  const [status, setStatus] = useState<types.StatusContainer>({} as types.StatusContainer);
  const [daemons, setDaemons] = useState<types.DaemonContainer>({} as types.DaemonContainer);
  const [session, setSession] = useState<types.SessionContainer>({} as types.SessionContainer);
  const [config, setConfig] = useState<types.ConfigContainer>({} as types.ConfigContainer);
  const [wizard, setWizard] = useState<types.WizardContainer>({} as types.WizardContainer);
  const [info, setInfo] = useState<app.AppInfo>({} as app.AppInfo);
  const [route, setRoute] = useState<string>("project");
  const [activeTab, setActiveTab] = useState<string>("project");
  const [, setLocation] = useLocation();
  const counters = useRef<Record<string, number>>({});

  // ------------------- Route/Tab State -------------------
  useEffect(() => {
    GetLastRoute().then((lastRoute) => {
      setRoute(lastRoute);
      GetLastTab().then((lastTab) => {
        setActiveTab(lastTab);
      });
    });
  }, []);

  useEffect(() => {
    const handleNavigation = (msg: messages.MessageMsg) => {
      setLocation("/" + msg.string1);
      setRoute(msg.string1);
      setActiveTab(msg.string2);
    };

    const { Message } = messages;
    EventsOn(Message.NAVIGATE, handleNavigation);
    return () => {
      EventsOff(Message.NAVIGATE);
    };
  }, [setLocation]);

  const tabChanged = (newVal: string) => {
    SetLastTab(route, newVal).then(() => {
      setActiveTab(newVal);
    });
  };

  const routeChanged = (newVal: string) => {
    SetLastRoute(newVal).then(() => {
      setRoute(newVal);
      GetLastTab().then((lastTab) => {
        tabChanged(lastTab);
      });
    });
  };

  // ------------------- Data Fetches -------------------
  const fetchProject = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchProject(currentItem, itemsPerPage).then((item: types.ProjectContainer) => {
      setProject(item);
    });
  }, []);

  const fetchBalances = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchBalance(currentItem, itemsPerPage).then((item: types.BalanceContainer) => {
      setBalances(item);
    });
  }, []);

  const fetchIncoming = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchIncoming(currentItem, itemsPerPage).then((item: types.IncomingContainer) => {
      setIncoming(item);
    });
  }, []);

  const fetchOutgoing = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchOutgoing(currentItem, itemsPerPage).then((item: types.OutgoingContainer) => {
      setOutgoing(item);
    });
  }, []);

  const fetchInternals = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchInternal(currentItem, itemsPerPage).then((item: types.InternalContainer) => {
      setInternals(item);
    });
  }, []);

  const fetchCharts = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchChart(currentItem, itemsPerPage).then((item: types.ChartContainer) => {
      setCharts(item);
    });
  }, []);

  const fetchLogs = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchLog(currentItem, itemsPerPage).then((item: types.LogContainer) => {
      setLogs(item);
    });
  }, []);

  const fetchStatements = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchStatement(currentItem, itemsPerPage).then((item: types.StatementContainer) => {
      setStatements(item);
    });
  }, []);

  const fetchNeighbors = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchNeighbor(currentItem, itemsPerPage).then((item: types.NeighborContainer) => {
      setNeighbors(item);
    });
  }, []);

  const fetchTraces = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchTrace(currentItem, itemsPerPage).then((item: types.TraceContainer) => {
      setTraces(item);
    });
  }, []);

  const fetchReceipts = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchReceipt(currentItem, itemsPerPage).then((item: types.ReceiptContainer) => {
      setReceipts(item);
    });
  }, []);

  const fetchMonitors = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchMonitor(currentItem, itemsPerPage).then((item: types.MonitorContainer) => {
      setMonitors(item);
    });
  }, []);

  const fetchNames = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchName(currentItem, itemsPerPage).then((item: types.NameContainer) => {
      setNames(item);
    });
  }, []);

  const fetchAbis = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchAbi(currentItem, itemsPerPage).then((item: types.AbiContainer) => {
      setAbis(item);
    });
  }, []);

  const fetchIndexes = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchIndex(currentItem, itemsPerPage).then((item: types.IndexContainer) => {
      setIndexes(item);
    });
  }, []);

  const fetchManifests = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchManifest(currentItem, itemsPerPage).then((item: types.ManifestContainer) => {
      setManifests(item);
    });
  }, []);

  const fetchPins = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchPin(currentItem, itemsPerPage).then((item: types.PinContainer) => {
      setPins(item);
    });
  }, []);

  const fetchUploads = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchUpload(currentItem, itemsPerPage).then((item: types.UploadContainer) => {
      setUploads(item);
    });
  }, []);

  const fetchStatus = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchStatus(currentItem, itemsPerPage).then((item: types.StatusContainer) => {
      setStatus(item);
    });
  }, []);

  const fetchDaemons = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchDaemon(currentItem, itemsPerPage).then((item: types.DaemonContainer) => {
      setDaemons(item);
    });
  }, []);

  const fetchSession = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchSession(currentItem, itemsPerPage).then((item: types.SessionContainer) => {
      setSession(item);
    });
  }, []);

  const fetchConfig = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchConfig(currentItem, itemsPerPage).then((item: types.ConfigContainer) => {
      setConfig(item);
    });
  }, []);

  const fetchWizard = useCallback((currentItem: number, itemsPerPage: number) => {
    // Note that this only fetches a single page after sorting and filtering (if any)
    FetchWizard(currentItem, itemsPerPage).then((item: types.WizardContainer) => {
      setWizard(item);
    });
  }, []);

  const loadAddress = (address: base.Address) => {
    const addressStr = address as unknown as string;
    LoadAddress(addressStr).then(() => {
      info.address = address;
      setInfo(info);
    });
  };

  const fetchAppInfo = () => {
    FetchAppInfo().then((info) => {
      setInfo(info);
    });
  };

  useEffect(() => {
    const handleRefresh = () => {
      fetchAppInfo();
      fetchWizard(0, 100);
      fetchStatus(0, 100);
    };
    handleRefresh(); // first load

    // when messaged
    const { Message } = messages;
    EventsOn(Message.REFRESH, handleRefresh);
    return () => {
      EventsOff(Message.REFRESH);
    };
  }, [fetchStatus, fetchWizard]);

  return (
    <AppState.Provider
      value={{
        project,
        fetchProject,
        balances,
        fetchBalances,
        incoming,
        fetchIncoming,
        outgoing,
        fetchOutgoing,
        internals,
        fetchInternals,
        charts,
        fetchCharts,
        logs,
        fetchLogs,
        statements,
        fetchStatements,
        neighbors,
        fetchNeighbors,
        traces,
        fetchTraces,
        receipts,
        fetchReceipts,
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
        pins,
        fetchPins,
        uploads,
        fetchUploads,
        status,
        fetchStatus,
        daemons,
        fetchDaemons,
        session,
        fetchSession,
        config,
        fetchConfig,
        wizard,
        fetchWizard,
        info,
        loadAddress,
        counters,
        route,
        routeChanged,
        activeTab,
        tabChanged,
      }}
    >
      {children}
    </AppState.Provider>
  );
};

export const useAppState = () => {
  const context = useContext(AppState);
  if (!context) {
    throw new Error("useAppState must be used within a AppStateProvider");
  }
  return context;
};
