import { useState, createContext, useEffect, useContext, ReactNode } from "react";
import { Pager } from "@components";
import { IsShowing, SetShowing, SetFilter, GetActiveTab } from "@gocode/app/App";
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
  headerShows: Record<string, boolean>;
  handleCollapse: (tab: string, newState: string | null) => void;
  pager: Pager;
  fetchFn: FetchFnType;
  modifyFn: ModifyFnType;
  filter: string;
  updateFilter: (criteria: string) => void;
  clickFn?: ClickFnType;
  activeTab: string;
  setActiveTab: React.Dispatch<React.SetStateAction<string>>;
}

const ViewContext = createContext<ViewStateProps | undefined>(undefined);

type ViewContextType = {
  route: Route;
  nItems: number;
  fetchFn: FetchFnType;
  modifyFn: ModifyFnType;
  onEnter: (page: Page) => void;
  clickFn?: ClickFnType;
  tabs: string[];
  children: ReactNode;
};

export const ViewStateProvider = ({
  route,
  nItems,
  fetchFn,
  modifyFn,
  onEnter,
  clickFn,
  tabs,
  children,
}: ViewContextType) => {
  const [headerShows, setHeaderShows] = useState<Record<string, boolean>>({});
  const [activeTab, setActiveTab] = useState<string>(tabs[0] || ""); // Initialize with the first tab or an empty string
  const [filter, setFilter] = useState<string>("");
  // TODO: This used to be different for session and names, but those are
  // TODO: now tabs. This points to the fact that we need per-tab state.
  // TODO: There are other places -- active tab, header open/close, etc.
  const lines = 10;
  const pager = useKeyboardPaging(nItems, lines, onEnter);

  const handleCollapse = (tab: string, newState: string | null) => {
    const isShowing = newState === "header";
    SetShowing(route, tab, isShowing).then(() => {
      setHeaderShows((prev) => ({
        ...prev,
        [tab]: isShowing,
      }));
    });
  };

  const updateFilter = (criteria: string) => {
    setFilter(criteria);
    SetFilter(route, criteria).then(() => {
      fetchFn(pager.getOffset(), pager.perPage);
    });
  };

  useEffect(() => {
    const fetchActiveTab = async () => {
      const storedTab = await GetActiveTab(route); // Add a backend function to retrieve the active tab
      setActiveTab(storedTab || tabs[0]); // Fallback to the first tab if no stored tab exists
    };

    fetchActiveTab();
  }, [route, tabs]);

  useEffect(() => {
    const fetchHeaderStates = async () => {
      const newHeaderShows: Record<string, boolean> = {};
      for (const tab of tabs) {
        const isShowing = await IsShowing(route, tab);
        newHeaderShows[tab] = isShowing;
      }
      setHeaderShows(newHeaderShows);
    };

    fetchHeaderStates();
  }, [route, tabs]);

  useEffect(() => {
    const handleAccordion = (msg: messages.MessageMsg) => {
      const cmp = route === "" ? "project" : route;
      const tab = msg.string2 || ""; // Use msg.string2 as the tab if provided, or fallback to an empty string.
      if (tab && cmp === msg.string1) {
        IsShowing(cmp, tab).then((onOff) => {
          SetShowing(cmp, tab, !onOff).then(() => {
            setHeaderShows((prev) => ({
              ...prev,
              [tab]: !onOff,
            }));
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
    activeTab,
    setActiveTab,
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
