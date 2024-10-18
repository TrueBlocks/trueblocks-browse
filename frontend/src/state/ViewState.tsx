import { useState, createContext, useEffect, useContext, ReactNode } from "react";
import { useLocation } from "wouter";
import { Pager } from "@components";
import { HistoryPage, IsShowing, SetShowing } from "@gocode/app/App";
import { types, messages, app } from "@gocode/models";
import { Page, useKeyboardPaging } from "@hooks";
import { Route } from "@layout";
import { EventsOn, EventsOff } from "@runtime";
import { useAppState } from "@state";

type ModifyFnType = (arg1: app.ModifyData) => Promise<void>;
type FetchFnType = (selected: number, perPage: number) => void;

interface ViewStateProps {
  route: Route;
  nItems: number;
  collapsed: boolean;
  handleCollapse: (value: string | null) => void;
  pager: Pager;
  fetchFn: FetchFnType;
  modifyFn: ModifyFnType;
}

const ViewContext = createContext<ViewStateProps | undefined>(undefined);

type ViewContextType = {
  route: Route;
  nItems?: number;
  fetchFn: FetchFnType;
  modifyFn: ModifyFnType;
  onEnter?: (page: Page) => void;
  children: ReactNode;
};

export const ViewStateProvider = ({ route, nItems = -1, fetchFn, modifyFn, onEnter, children }: ViewContextType) => {
  const [collapsed, setCollapsed] = useState<boolean>(false);
  const [location] = useLocation();

  const { address, setHistory } = useAppState();
  const lines = route === "status" ? 6 : route === "names" ? 9 : 10;
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const ignoreEnter = (_unused: Page) => {};
  const pager = useKeyboardPaging(route, nItems, lines, onEnter ? onEnter : ignoreEnter);

  useEffect(() => {
    IsShowing(route).then((showing) => {
      setCollapsed(showing);
    });
  }, [route, location]);

  function handleCollapse(newState: string | null) {
    const isCollapsed = newState ? false : true;
    SetShowing(route, isCollapsed).then(() => {
      setCollapsed(isCollapsed ? true : false);
    });
  }

  useEffect(() => {
    fetchFn(pager.getOffset(), pager.perPage);
  }, [pager.pageNumber, pager.perPage]);

  useEffect(() => {
    const handleRefresh = () => {
      fetchFn(pager.getOffset(), pager.perPage);
    };

    const { Message } = messages;
    EventsOn(Message.DAEMON, handleRefresh);
    return () => {
      EventsOff(Message.DAEMON);
    };
  }, [fetchFn, nItems, pager]);

  useEffect(() => {
    if (route === "history") {
      HistoryPage(String(address), pager.getOffset(), pager.perPage).then((item: types.HistoryContainer) => {
        setHistory(item);
      });
    }
  }, [address, pager.pageNumber, pager.perPage]);

  const state = {
    route,
    nItems,
    collapsed,
    handleCollapse,
    pager,
    fetchFn,
    modifyFn,
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
