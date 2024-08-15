import React, { useState, useEffect } from "react";
import { messages } from "@gocode/models";
import { AppShell, Text } from "@mantine/core";
import { Aside, Header, Navbar, Routes } from "@components";
import { EventsOn, EventsOff } from "@runtime";
import { useLocation } from "wouter";
import classes from "@/App.module.css";
import { GetLast, SetLast } from "@gocode/app/App";

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
        <Text size={"sm"}>time / date / currently opened file</Text>
      </AppShell.Footer>
    </AppShell>
  );
}

export default App;
