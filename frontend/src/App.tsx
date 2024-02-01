import React from "react";
import { AppShell } from "@mantine/core";
import GlobalNavbar from "./components/global/GlobalNavbar";
import Routes from "./components/global/Routes";

function App() {
  return (
    <AppShell navbar={{ width: "15rem", breakpoint: 0 }}>
      <AppShell.Navbar>
        {/* To change menu items, go to 'GlobalMenu' component */}
        <GlobalNavbar />
      </AppShell.Navbar>
      <AppShell.Main>
        {/* To add new views, go to 'Routes' component */}
        <Routes />
      </AppShell.Main>
    </AppShell>
  );
}

export default App;
