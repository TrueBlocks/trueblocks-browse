import { createContext, useEffect, useContext, ReactNode } from "react";
import { Route } from "@/Routes";
import { Pager } from "@components";
import { HistoryPage } from "@gocode/app/App";
import { types, messages } from "@gocode/models";
import { Page, useKeyboardPaging } from "@hooks";
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
  fetchFn: (selected: number, perPage: number) => void;
  onEnter?: (page: Page) => void;
  children: ReactNode;
}> = ({ route, nItems = -1, fetchFn, onEnter, children }) => {
  const { address, setHistory } = useAppState();
  const lines = route === "status" ? 6 : route === "names" ? 9 : 10;
  const ignoreEnter = (page: Page) => {};
  const pager = useKeyboardPaging(route, nItems, lines, onEnter ? onEnter : ignoreEnter);

  useEffect(() => {
    fetchFn(pager.getOffset(), pager.perPage);
  }, [pager.pageNumber, pager.perPage]);

  useEffect(() => {
    const handleRefresh = () => {
      fetchFn(pager.getOffset(), pager.perPage);
      /*
      // Dawid: This doesn't really work. It ignores, for example, when the data is
      // Dawid: with the latest data at the end. Plus, it has no effect on performance
      // Fetch page only if it makes sense: the current page is the first page
      // (showing the latest transactions) and is incomplete.
      // Otherwise we get into constant rerendering
      if (pager.pageNumber === 1 && nItems < pager.perPage) {
        fetchFn(pager.getOffset(), pager.perPage);
      }
      */
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
