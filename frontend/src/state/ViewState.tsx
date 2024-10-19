import { useState, createContext, useEffect, useContext, ReactNode } from "react";
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
  headerShows: boolean | null;
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
  const [headerShows, setHeaderShows] = useState<boolean | null>(null);
  const { address, setHistory } = useAppState();
  const lines = route === "status" ? 6 : route === "names" ? 9 : 10;
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const ignoreEnter = (_unused: Page) => {};
  const pager = useKeyboardPaging(route, nItems, lines, onEnter ? onEnter : ignoreEnter);

  const handleCollapse = (newState: string | null) => {
    const isShowing = newState === "header";
    SetShowing(route, isShowing).then(() => {
      setHeaderShows(isShowing);
    });
  };

  useEffect(() => {
    IsShowing(route).then((onOff) => {
      setHeaderShows(onOff);
    });
  }, [route]);

  useEffect(() => {
    const handleToggle = (msg: messages.ToggleMsg) => {
      if (msg.layout === "" && route === msg.route) {
        IsShowing(route).then((onOff) => {
          SetShowing(route, !onOff).then(() => {
            setHeaderShows(!onOff);
          });
        });
      }
    };
    EventsOn(messages.Message.TOGGLEHEADER, handleToggle);
    return () => {
      EventsOff(messages.Message.TOGGLEHEADER);
    };
  }, [route]);

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
    headerShows,
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
