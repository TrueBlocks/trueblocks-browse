import { useState, useEffect } from "react";
import { AppShell } from "@mantine/core";
import { Aside, Header, Navbar, Routes, AppStatus } from "@components";
import { GetLast, SetLast } from "@gocode/app/App";
import { messages } from "@gocode/models";
import { EventsOn, EventsOff } from "@runtime";
import { AppStateProvider } from "@state";
import classes from "@/App.module.css";

function App() {
  const [showHelp, setShowHelp] = useState<boolean>(false);

  useEffect(() => {
    GetLast("help").then((value) => {
      setShowHelp(value === "true");
    });

    const toggleHelp = () => {
      setShowHelp((prevShowHelp) => {
        const newShowHelp = !prevShowHelp;
        SetLast("help", `${newShowHelp ? "true" : "false"}`);
        return newShowHelp;
      });
    };

    const { Message } = messages;
    EventsOn(Message.TOGGLEHELP, toggleHelp);
    return () => {
      EventsOff(Message.TOGGLEHELP);
    };
  }, []);

  return (
    <AppStateProvider>
      <AppShell
        header={{ height: "3rem" }}
        navbar={{ collapsed: { desktop: false }, width: "10rem", breakpoint: 0 }}
        aside={{ collapsed: { desktop: !showHelp }, width: "20rem", breakpoint: 0 }}
        footer={{ height: "2rem" }}
      >
        <AppShell.Header>
          <Header title="TrueBlocks Browse" />
        </AppShell.Header>
        <AppShell.Navbar>
          <Navbar />
        </AppShell.Navbar>
        <AppShell.Main className={classes.mainContent}>
          <Routes />
        </AppShell.Main>
        <AppShell.Aside>
          <Aside />
        </AppShell.Aside>
        <AppShell.Footer>
          <AppStatus />
        </AppShell.Footer>
      </AppShell>
    </AppStateProvider>
  );
}

export default App;
