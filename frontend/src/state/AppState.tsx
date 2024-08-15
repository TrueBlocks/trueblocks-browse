import React, { createContext, useContext, useState, ReactNode } from "react";
import { types } from "@gocode/models";

interface AppStateProps {
  history: types.TransactionContainer;
  monitors: types.MonitorContainer;
  names: types.NameContainer;
  abis: types.AbiContainer;
  indexes: types.IndexContainer;
  manifests: types.ManifestContainer;
  status: types.StatusContainer;

  setHistory: (history: types.TransactionContainer) => void;
  setMonitors: (monitors: types.MonitorContainer) => void;
  setNames: (names: types.NameContainer) => void;
  setAbis: (abis: types.AbiContainer) => void;
  setIndexes: (indexes: types.IndexContainer) => void;
  setManifests: (manifests: types.ManifestContainer) => void;
  setStatus: (status: types.StatusContainer) => void;

  // settings: types.SettingsContainer;
  // setSettings: (settings: types.SettingsContainer) => void;
}

const AppState = createContext<AppStateProps | undefined>(undefined);

export const AppStateProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [history, setHistory] = useState<types.TransactionContainer>({} as types.TransactionContainer);
  const [monitors, setMonitors] = useState<types.MonitorContainer>({} as types.MonitorContainer);
  const [names, setNames] = useState<types.NameContainer>({} as types.NameContainer);
  const [abis, setAbis] = useState<types.AbiContainer>({} as types.AbiContainer);
  const [indexes, setIndexes] = useState<types.IndexContainer>({} as types.IndexContainer);
  const [manifests, setManifests] = useState<types.ManifestContainer>({} as types.ManifestContainer);
  const [status, setStatus] = useState<types.StatusContainer>({} as types.StatusContainer);
  return (
    <AppState.Provider
      value={{
        history,
        monitors,
        names,
        abis,
        indexes,
        manifests,
        status,
        setHistory,
        setMonitors,
        setNames,
        setAbis,
        setIndexes,
        setManifests,
        setStatus,
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
