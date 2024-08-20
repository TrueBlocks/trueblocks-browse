import React, { createContext, useContext, useState, ReactNode } from "react";
import { Pager } from "@components";
import { Route } from "@/Routes";
import { useKeyboardPaging } from "@hooks";

interface ViewStateProps {
  route: Route;
  nItems: number;
  pager: Pager;
  setPager: React.Dispatch<React.SetStateAction<Pager>>;
}

const ViewContext = createContext<ViewStateProps | undefined>(undefined);

export const ViewStateProvider: React.FC<{ route: Route; nItems: number; children: ReactNode }> = ({
  route,
  nItems,
  children,
}) => {
  const [pager, setPager] = useState<Pager>(useKeyboardPaging(route, nItems, [], 15));
  let state = {
    route,
    nItems,
    pager,
    setPager,
  };

  return <ViewContext.Provider value={state}>{children}</ViewContext.Provider>;
};

export const useViewState = () => {
  const context = useContext(ViewContext);
  if (!context) {
    throw new Error("useViewState must be used within a ViewProvider");
  }
  return context;
};
