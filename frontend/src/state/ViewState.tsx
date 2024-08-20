import React, { createContext, useContext, useState, ReactNode } from "react";
import { Pager } from "@components";
import { Route } from "@/Routes";
import { useKeyboardPaging } from "@hooks";

// TODO: Complicated situation during development. Will be corrected. There are two pagers
// TODO: for the same route. This one (in the ViewContext) and the one in the useAppState.
// TODO: We would rather have only one (this one) but we need to refactor the useAppState
// TODO: so that it's not required there. Since both pagers (per route) listen for and
// TODO: respond to "DAEMON" (i.e. freshen) messages, it works fine. Just inefficient.
// TODO: The useAppState pager is used to grab a page from the backend state. This pager
// TODO: causes the page to advance.

interface ViewStateProps {
  route: Route;
  nItems: number;
  pager: Pager;
  setPager: React.Dispatch<React.SetStateAction<Pager>>;
}

const ViewContext = createContext<ViewStateProps | undefined>(undefined);

export const ViewStateProvider: React.FC<{
  route: Route;
  nItems: number;
  children: ReactNode;
}> = ({
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
