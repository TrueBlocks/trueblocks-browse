import React from "react";
import { AppShell, Text } from "@mantine/core";
import Aside from "./components/global/Aside";
import Header from "./components/global/Header";
import Navbar from "./components/global/Navbar";
import Routes from "./components/global/Routes";
import classes from "@/App.module.css";

function App() {
  return (
    <AppShell
      header={{ height: "3rem" }}
      navbar={{ collapsed: { desktop: false }, width: "10rem", breakpoint: 0 }}
      aside={{ collapsed: { desktop: false }, width: "10rem", breakpoint: 0 }}
      footer={{ height: "2rem" }}
    >
      <AppShell.Header>
        <Header title="ApplicationTitle" />
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
