import React, { createContext, useEffect, useContext, ReactNode } from "react";
import { Pager } from "@components";
import { Route } from "@/Routes";
import { useKeyboardPaging } from "@hooks";
import { types, messages } from "@gocode/models";
import { HistoryPage } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";
import { useAppState } from "@state";

interface ViewStateProps {
  route: Route;
  nItems: number;
  pager: Pager;
  getViewPager: (route: Route) => Pager | null;
}

const ViewContext = createContext<ViewStateProps | undefined>(undefined);

export const ViewStateProvider: React.FC<{
  route: Route;
  nItems?: number;
  fetchFn: (curItem: number, perPage: number, item?: any) => void;
  children: ReactNode;
}> = ({ route, nItems = -1, fetchFn, children }) => {
  const { address, setHistory } = useAppState();
  const lines = route === "status" ? 10 : 14;
  const pager = useKeyboardPaging(route, nItems, [], lines);

  useEffect(() => {
    fetchFn(pager.curItem, pager.perPage, null);
  }, [pager.curItem, pager.perPage]);

  useEffect(() => {
    const handleRefresh = () => {
      fetchFn(pager.curItem, pager.perPage);
    };

    var { Message } = messages;
    EventsOn(Message.DAEMON, handleRefresh);
    return () => {
      EventsOff(Message.DAEMON);
    };
  }, [EventsOn, EventsOff, fetchFn]);

  useEffect(() => {
    if (route === "history") {
      HistoryPage(address as unknown as string, pager.curItem, pager.perPage).then(
        (item: types.TransactionContainer) => {
          if (item) {
            setHistory(item);
          }
        }
      );
    }
  }, [address, pager.curItem, pager.perPage]);

  const getViewPager = (route: Route): Pager => {
    return pager;
  };

  let state = {
    route,
    nItems,
    pager,
    getViewPager,
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
