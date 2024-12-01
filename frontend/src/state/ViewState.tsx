import { useState, createContext, useEffect, useContext, ReactNode } from "react";
import { Pager } from "@components";
import { IsShowing, SetShowing, SetFilter, GetFilter } from "@gocode/app/App";
import { messages, app } from "@gocode/models";
import { Page, useKeyboardPaging } from "@hooks";
import { Route } from "@layout";
import { EventsOn, EventsOff } from "@runtime";

type ModifyFnType = (arg1: app.ModifyData) => Promise<void>;
type FetchFnType = (selected: number, perPage: number) => void;
type ClickFnType = (value: string) => void;

interface ViewStateProps {
  route: Route;
  nItems: number;
  headerShows: boolean | null;
  handleCollapse: (value: string | null) => void;
  pager: Pager;
  fetchFn: FetchFnType;
  modifyFn: ModifyFnType;
  filter: string;
  updateFilter: (criteria: string) => void;
  clickFn?: ClickFnType;
}

const ViewContext = createContext<ViewStateProps | undefined>(undefined);

type ViewContextType = {
  route: Route;
  nItems: number;
  fetchFn: FetchFnType;
  modifyFn: ModifyFnType;
  onEnter: (page: Page) => void;
  clickFn?: ClickFnType;
  children: ReactNode;
};

export const ViewStateProvider = ({
  route,
  nItems,
  fetchFn,
  modifyFn,
  onEnter,
  clickFn,
  children,
}: ViewContextType) => {
  const [headerShows, setHeaderShows] = useState<boolean | null>(null);
  const [filter, setFilter] = useState<string>("");
  // TODO: This used to be different for session and names, but those are
  // TODO: now tabs. This points to the fact that we need per-tab state.
  // TODO: There are other places -- active tab, header open/close, etc.
  const lines = 10;
  const pager = useKeyboardPaging(nItems, lines, onEnter);

  const handleCollapse = (newState: string | null) => {
    const isShowing = newState === "header";
    SetShowing(route, isShowing).then(() => {
      setHeaderShows(isShowing);
    });
  };

  const updateFilter = (criteria: string) => {
    setFilter(criteria);
    SetFilter(route, criteria).then(() => {
      fetchFn(pager.getOffset(), pager.perPage);
    });
  };

  useEffect(() => {
    IsShowing(route).then((onOff) => {
      setHeaderShows(onOff);
    });

    GetFilter(route).then((filterData) => {
      setFilter(filterData.criteria);
    });
  }, [route]);

  useEffect(() => {
    const handleAccordion = (msg: messages.MessageMsg) => {
      const cmp = route === "" ? "project" : route;
      if (msg.string2 === "" && cmp === msg.string1) {
        IsShowing(cmp).then((onOff) => {
          SetShowing(cmp, !onOff).then(() => {
            setHeaderShows(!onOff);
          });
        });
      }
    };

    const { Message } = messages;
    EventsOn(Message.TOGGLEACCORDION, handleAccordion);
    return () => {
      EventsOff(Message.TOGGLEACCORDION);
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
    EventsOn(Message.REFRESH, handleRefresh);
    return () => {
      EventsOff(Message.REFRESH);
    };
  }, [fetchFn, nItems, pager]);

  const state = {
    route,
    nItems,
    headerShows,
    handleCollapse,
    pager,
    fetchFn,
    modifyFn,
    filter,
    updateFilter,
    clickFn,
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
