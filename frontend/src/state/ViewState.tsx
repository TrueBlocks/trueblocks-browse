import { useState, createContext, useEffect, useContext, ReactNode } from "react";
import { Pager } from "@components";
import { IsHeaderOn, SetHeaderOn, SetFilter, GetLastTab, GetFilter, SetLastTab } from "@gocode/app/App";
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
  const lines = 10;
  const pager = useKeyboardPaging(nItems, lines, onEnter);

  // - route/tabs -----------------------------------------------------------
  useEffect(() => {
    GetLastTab(route).then((tab) => {
      setActiveTab(tab || tabs[0]);
    });
  }, [route, tabs, setActiveTab]);

  // - route/tabs -----------------------------------------------------------
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
        SetLastTab(route, newTab);
      }
    };

    EventsOn(messages.Message.SWITCHTAB, handleSwitchTab);
    return () => {
      EventsOff(messages.Message.SWITCHTAB);
    };
  }, [route, tabs, activeTab, setActiveTab]);

  // - headers -----------------------------------------------------------
  const handleCollapse = (route: string, newState: string | null) => {
    const isShowing = newState === "header";
    const key = `${route}-${activeTab}`;
    SetHeaderOn(route, activeTab, isShowing).then(() => {
      setHeaderShows((prev) => ({
        ...prev,
        [key]: isShowing,
      }));
    });
  };

  // - headers -----------------------------------------------------------
  useEffect(() => {
    const fetchHeaderStates = () => {
      tabs.forEach((tab) => {
        const key = `${route}-${tab}`;
        IsHeaderOn(route, tab).then((isShowing) => {
          setHeaderShows((prev) => ({
            ...prev,
            [key]: isShowing,
          }));
        });
      });
    };

    fetchHeaderStates();
  }, [route, tabs]);

  // - headers -----------------------------------------------------------
  useEffect(() => {
    const handleAccordion = (route: string, tab: string, isShowing: boolean) => {
      const key = `${route}-${tab}`;
      setHeaderShows((prev) => ({
        ...prev,
        [key]: isShowing,
      }));
    };

    const { Message } = messages;
    EventsOn(Message.TOGGLEACCORDION, (msg: { string1: string; string2: string; bool: boolean }) => {
      handleAccordion(msg.string1, msg.string2, msg.bool);
    });

    return () => {
      EventsOff(Message.TOGGLEACCORDION);
    };
  }, []);

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
    route,
    nItems,
    headerShows,
    handleCollapse,
    pager,
    fetchFn,
    modifyFn,
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
