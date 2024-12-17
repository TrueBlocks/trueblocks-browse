import { createContext, useEffect, useContext, ReactNode } from "react";
import { Pager } from "@components";
import { messages, app } from "@gocode/models";
import { Page, useKeyboardPaging } from "@hooks";
import { EventsOn, EventsOff } from "@runtime";

type ModifyFnType = (arg1: app.ModifyData) => Promise<void>;
type FetchFnType = (selected: number, perPage: number) => void;
type ClickFnType = (value: string) => void;

interface ViewStateProps {
  nItems: number;
  pager: Pager;
  fetchFn: FetchFnType;
  modifyFn: ModifyFnType;
  clickFn?: ClickFnType;
}

const ViewContext = createContext<ViewStateProps | undefined>(undefined);

type ViewContextType = {
  nItems: number;
  fetchFn: FetchFnType;
  modifyFn: ModifyFnType;
  onEnter: (page: Page) => void;
  clickFn?: ClickFnType;
  children: ReactNode;
};

export const ViewStateProvider = ({ nItems, fetchFn, modifyFn, onEnter, clickFn, children }: ViewContextType) => {
  const lines = 10;
  const pager = useKeyboardPaging(nItems, lines, onEnter);

  // - pagination -----------------------------------------------------------
  useEffect(() => {
    fetchFn(pager.getOffset(), pager.perPage);
  }, [pager.pageNumber, pager.perPage]);

  // - pagination -----------------------------------------------------------
  useEffect(() => {
    const handleRefresh = () => {
      fetchFn(pager.getOffset(), pager.perPage);
    };

    const { Message } = messages;
    EventsOn(Message.REFRESH, handleRefresh);
    return () => {
      EventsOff(Message.REFRESH);
    };
  }, [fetchFn, nItems, pager]);

  // - state -----------------------------------------------------------
  const state = {
    nItems,
    fetchFn,
    modifyFn,
    clickFn,
    pager,
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
