import { useState, useEffect, useMemo } from "react";
import { AppShell } from "@mantine/core";
import { Aside, Header, Navbar, ViewContainer, ViewStatus, AppStatus } from "@components";
import { IsShowing, SetShowing } from "@gocode/app/App";
import { messages } from "@gocode/models";
import { EventsOn, EventsOff } from "@runtime";
import { AppStateProvider } from "@state";
import classes from "./App.module.css";

export const App = () => {
  const [showHeader, setShowHeader] = useState<boolean>(true);
  const [showMenu, setShowMenu] = useState<boolean>(true);
  const [showHelp, setShowHelp] = useState<boolean>(true);
  const [showFooter, setShowFooter] = useState<boolean>(true);
  const toggles = useMemo(
    () => [
      { component: "header", setter: setShowHeader },
      { component: "menu", setter: setShowMenu },
      { component: "help", setter: setShowHelp },
      { component: "footer", setter: setShowFooter },
    ],
    [setShowHeader, setShowMenu, setShowHelp, setShowFooter]
  );

  useEffect(() => {
    toggles.forEach(({ component, setter }) => {
      IsShowing(component).then((show) => {
        setter(show);
      });
    });
  }, [toggles]);

  useEffect(() => {
    const handleToggle = (msg: messages.ToggleMsg) => {
      const toggle = toggles.find((t) => t.component === msg.component);
      if (toggle) {
        toggle.setter((prev) => {
          const show = !prev;
          SetShowing(msg.component, show);
          return show;
        });
      }
    };

    EventsOn(messages.Message.TOGGLE, handleToggle);
    return () => {
      EventsOff(messages.Message.TOGGLE);
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
        <AppShell.Header>{showHeader ? <Header title="TrueBlocks Browse" /> : <></>}</AppShell.Header>
        <AppShell.Navbar>
          <Navbar />
        </AppShell.Navbar>
        <AppShell.Main className={classes.mainContent}>
          <ViewContainer />
          <ViewStatus />
        </AppShell.Main>
        <AppShell.Aside>
          <Aside />
        </AppShell.Aside>
        <AppShell.Footer>{showFooter ? <AppStatus /> : <></>}</AppShell.Footer>
      </AppShell>
    </AppStateProvider>
  );
};
