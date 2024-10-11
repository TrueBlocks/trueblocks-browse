import { useState, useEffect } from "react";
import { AppShell, Stack, Flex } from "@mantine/core";
import { Aside, Header, Navbar, Routes, AppStatus, ViewStatus } from "@components";
import { GetSessionVal, SetSessionVal } from "@gocode/app/App";
import { messages } from "@gocode/models";
import { EventsOn, EventsOff } from "@runtime";
import { AppStateProvider } from "@state";
// import classes from "@/App.module.css";
import classes from "./components/view/View.module.css";

function App() {
  const [showHelp, setShowHelp] = useState<boolean>(false);

  useEffect(() => {
    GetSessionVal("help").then((value) => {
      setShowHelp(value === "true");
    });

    const toggleHelp = () => {
      setShowHelp((prevShowHelp) => {
        const newShowHelp = !prevShowHelp;
        SetSessionVal("help", `${newShowHelp ? "true" : "false"}`);
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
        <AppShell.Main style={{ height: "100vh", display: "flex", flexDirection: "column" }}>
          <div
            style={{
              flexGrow: 1,
              overflowY: "auto",
              backgroundColor: "red",
              padding: "1em",
            }}
          >
            <Routes />
          </div>
          <div
            className={classes.viewStatus}
            style={{
              height: "2rem",
              backgroundColor: "transparent",
              position: "sticky",
              bottom: 0,
            }}
          >
            <ViewStatus />
          </div>
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
