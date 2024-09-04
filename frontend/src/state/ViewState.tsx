import React, { createContext, useEffect, useContext, ReactNode, useMemo } from "react";
import { Pager } from "@components";
import { Route } from "@/Routes";
import { messages } from "@gocode/models";
import { Page, useKeyboardPaging } from "@hooks";
import { EventsOn, EventsOff } from "@runtime";

import { Route } from "@/Routes";

interface ViewStateProps {
  route: Route;
  nItems: number;
  pager: Pager;
}

const ViewContext = createContext<ViewStateProps | undefined>(undefined);

export const ViewStateProvider: React.FC<{
  route: Route;
  nItems?: number;
  fetchFn: (selected: number, perPage: number, item?: any) => void;
  onEnter?: (page: Page) => void;
  children: ReactNode;
}> = ({ route, nItems = -1, fetchFn, onEnter, children }) => {
  const lines = route === "status" ? 6 : route === "names" ? 9 : 10;
  const ignoreEnter = (page: Page) => {};
  const pager = useKeyboardPaging(route, nItems, lines, onEnter ? onEnter : ignoreEnter);

  useEffect(() => {
    fetchFn(pager.getOffset(), pager.perPage, null);
  }, [pager.pageNumber, pager.perPage]);

  useEffect(() => {
    const handleRefresh = () => {
      // Fetch page only if it makes sense: the current page is the first page
      // (showing the latest transactions) and is incomplete.
      // Otherwise we get into constant rerendering
      if (pager.pageNumber === 1 && nItems < pager.perPage) {
        fetchFn(pager.getOffset(), pager.perPage);
      }
    };

    const { Message } = messages;
    EventsOn(Message.DAEMON, handleRefresh);
    return () => {
      EventsOff(Message.DAEMON);
    };
  }, [fetchFn, nItems, pager]);

  const state = {
    route,
    nItems,
    pager,
  }), [route, nItems, pager]);

  return <ViewContext.Provider value={state}>{children}</ViewContext.Provider>;
};

export const useViewState = () => {
  const context = useContext(ViewContext);
  if (!context) {
    throw new Error("useViewState must be used within a ViewProvider");
  }
  return context;
};
