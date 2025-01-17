import { useState, useEffect, useMemo } from "react";
import { AppShell } from "@mantine/core";
import { ViewStatus } from "@components";
import { IsShowing, SetShowing, GetAppTitle } from "@gocode/app/App";
import { messages } from "@gocode/models";
import { Header, MenuPanel, ViewContainer, HelpPanel, Footer } from "@layout";
import { EventsOn, EventsOff } from "@runtime";
import { AppStateProvider } from "@state";
import classes from "./App.module.css";

export const App = () => {
  const [showHeader, setShowHeader] = useState<boolean>(true);
  const [showMenu, setShowMenu] = useState<boolean>(true);
  const [showHelp, setShowHelp] = useState<boolean>(true);
  const [showFooter, setShowFooter] = useState<boolean>(true);
  const [title, setTitle] = useState<string>("");

  useEffect(() => {
    GetAppTitle().then((title) => {
      setTitle(title);
    });
  }, []);

  const toggles = useMemo(
    () => [
      { layout: "header", setter: setShowHeader },
      { layout: "menu", setter: setShowMenu },
      { layout: "help", setter: setShowHelp },
      { layout: "footer", setter: setShowFooter },
    ],
    [setShowHeader, setShowMenu, setShowHelp, setShowFooter]
  );

  useEffect(() => {
    toggles.forEach(({ layout, setter }) => {
      IsShowing(layout).then((show) => {
        setter(show);
      });
    });
  }, [toggles]);

  useEffect(() => {
    const handleToggleLayout = (msg: messages.MessageMsg) => {
      const toggle = toggles.find((t) => t.layout === msg.string1);
      if (toggle) {
        toggle.setter((prev) => {
          const show = !prev;
          SetShowing(msg.string1, !prev);
          return show;
        });
      }
    };

    EventsOn(messages.Message.TOGGLELAYOUT, handleToggleLayout);
    return () => {
      EventsOff(messages.Message.TOGGLELAYOUT);
    };
  }, [toggles]);

  return (
    <AppStateProvider>
      <AppShell
        header={{ height: showHeader ? "3rem" : "0" }}
        navbar={{ collapsed: { desktop: !showMenu }, width: "10rem", breakpoint: 0 }}
        aside={{ collapsed: { desktop: !showHelp }, width: "20rem", breakpoint: 0 }}
        footer={{ height: showFooter ? "2rem" : "0" }}
      >
        <AppShell.Header>{showHeader ? <Header title={title} /> : <></>}</AppShell.Header>
        <AppShell.Navbar>
          <MenuPanel />
        </AppShell.Navbar>
        <AppShell.Main className={classes.mainContent}>
          <ViewContainer />
          <ViewStatus />
        </AppShell.Main>
        <AppShell.Aside>
          <HelpPanel />
        </AppShell.Aside>
        <AppShell.Footer>{showFooter ? <Footer /> : <></>}</AppShell.Footer>
      </AppShell>
    </AppStateProvider>
  );
};
