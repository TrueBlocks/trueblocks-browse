import React, { createContext, useContext, ReactNode } from "react";

interface AppStateProps {}

const AppState = createContext<AppStateProps | undefined>(undefined);

export const AppStateProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  let state = {};
  return <AppState.Provider value={state}>{children}</AppState.Provider>;
};

export const useAppState = () => {
  const context = useContext(AppState);
  if (!context) {
    throw new Error("useAppState must be used within a AppStateProvider");
  }
  return context;
};
