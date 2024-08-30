import React, { createContext, useEffect, useContext, ReactNode, useMemo } from "react";
import { EmptyPager, Pager } from "@components";
import { Route } from "@/Routes";
import { Page, useKeyboardPaging } from "@hooks";
import { types, messages } from "@gocode/models";
import { HistoryPage } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";
import { useAppState } from "@state";

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
  const { address, setHistory } = useAppState();
  const lines = route === "status" ? 6 : route === "names" ? 9 : 10;
  const ignoreEnter = (page: Page) => {};
  const pager = useKeyboardPaging(route, nItems, lines, onEnter ? onEnter : ignoreEnter);

  useEffect(() => {
    fetchFn(pager.getOffset(), pager.perPage, null);
  }, [pager.pageNumber, pager.perPage]);

  useEffect(() => {
    const handleRefresh = () => {
      fetchFn(pager.getOffset(), pager.perPage);
    };

    var { Message } = messages;
    EventsOn(Message.DAEMON, handleRefresh);
    return () => {
      EventsOff(Message.DAEMON);
    };
  }, [EventsOn, EventsOff, fetchFn]);

  useEffect(() => {
    if (route === "history") {
      const addr = address as unknown as string;
      HistoryPage(addr, pager.getOffset(), pager.perPage).then((item: types.HistoryContainer) => {
        if (item) {
          setHistory(item);
        }
      });
    }
  }, [address, pager.pageNumber, pager.perPage]);

  let state = useMemo(() => ({
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
