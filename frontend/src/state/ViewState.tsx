import { useState, createContext, useEffect, useContext, ReactNode } from "react";
import { Pager } from "@components";
import { IsShowing, SetShowing, SetFilter, GetFilter } from "@gocode/app/App";
import { messages, app } from "@gocode/models";
import { Page, useKeyboardPaging, useNoops } from "@hooks";
import { Route } from "@layout";
import { EventsOn, EventsOff } from "@runtime";

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
  filter: string;
  updateFilter: (criteria: string) => void;
}

const ViewContext = createContext<ViewStateProps | undefined>(undefined);

type ViewContextType = {
  route: Route;
  nItems: number;
  fetchFn: FetchFnType;
  modifyFn: ModifyFnType;
  onEnter: (page: Page) => void;
  children: ReactNode;
};

export const ViewStateProvider = ({ route, nItems, fetchFn, modifyFn, onEnter, children }: ViewContextType) => {
  const [headerShows, setHeaderShows] = useState<boolean | null>(null);
  const [filter, setFilter] = useState<string>("");
  const lines = route === "status" ? 6 : route === "names" ? 9 : 10;
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
