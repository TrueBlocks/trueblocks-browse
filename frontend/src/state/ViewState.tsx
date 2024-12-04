import { useState, createContext, useEffect, useContext, ReactNode } from "react";
import { Pager } from "@components";
import { IsHeaderOn, SetHeaderOn, SetFilter, GetActiveTab, GetFilter } from "@gocode/app/App";
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
  tabs: string[];
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
  const [activeTab, setActiveTab] = useState<string>("");
  const [headerShows, setHeaderShows] = useState<Record<string, boolean>>({});
  const [filter, setFilter] = useState<string>("");
  // TODO: `lines` used to be different for `session` and `names`, but
  // TODO: those are now tabs of SettingsView and SharingView. This points
  // TODO: to the fact that we need per-tab state. There are other places
  // TODO: where we need per-tab state as well -- activeTab, headerToggle, etc.
  const lines = 10;
  const pager = useKeyboardPaging(nItems, lines, onEnter);

  // - tabs -----------------------------------------------------------
  useEffect(() => {
    GetActiveTab(route).then((tab) => {
      setActiveTab(tab || tabs[0]);
    });
  }, [route, tabs, setActiveTab]);

  // - tabs -----------------------------------------------------------
  useEffect(() => {
    const handleSwitchTab = (msg: messages.MessageMsg) => {
      const { string1 } = msg;

      const currentIndex = tabs.indexOf(activeTab);
      let newTab = activeTab;

      switch (string1) {
        case "next":
          newTab = tabs[(currentIndex + 1) % tabs.length];
          break;
        case "prev":
          newTab = tabs[(currentIndex - 1 + tabs.length) % tabs.length];
          break;
        default:
          break;
      }

      if (newTab !== activeTab) {
        setActiveTab(newTab);
      }
    };

    EventsOn(messages.Message.SWITCHTAB, handleSwitchTab);
    return () => {
      EventsOff(messages.Message.SWITCHTAB);
    };
  }, [tabs, activeTab, setActiveTab]);

  // - headers -----------------------------------------------------------
  const handleCollapse = (tab: string, newState: string | null) => {
    const isShowing = newState === "header";
    SetHeaderOn(route, tab, isShowing).then(() => {
      setHeaderShows((prev) => ({
        ...prev,
        [tab]: isShowing,
      }));
    });
  };

  // - headers -----------------------------------------------------------
  useEffect(() => {
    const fetchHeaderStates = async () => {
      const newHeaderShows: Record<string, boolean> = {};
      for (const tab of tabs) {
        const isShowing = await IsHeaderOn(route, tab);
        newHeaderShows[tab] = isShowing;
      }
      setHeaderShows(newHeaderShows);
    };

    fetchHeaderStates();
  }, [route, tabs]);

  // - headers -----------------------------------------------------------
  useEffect(() => {
    const handleAccordion = (msg: messages.MessageMsg) => {
      const cmp = route === "" ? "project" : route;
      const tab = msg.string2 || ""; // Use msg.string2 as the tab if provided, or fallback to an empty string.
      if (tab && cmp === msg.string1) {
        IsHeaderOn(cmp, tab).then((onOff) => {
          SetHeaderOn(cmp, tab, !onOff).then(() => {
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

  // - filter -----------------------------------------------------------
  useEffect(() => {
    GetFilter(route).then((filterData) => {
      setFilter(filterData.criteria);
    });
  }, [route]);

  // - filter -----------------------------------------------------------
  const updateFilter = (criteria: string) => {
    setFilter(criteria);
    SetFilter(route, criteria).then(() => {
      fetchFn(pager.getOffset(), pager.perPage);
    });
  };

  // - state -----------------------------------------------------------
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
    tabs,
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
