import React from "react";
import { Center, Stack, Title } from "@mantine/core";
import GlobalFooter from "./GlobalFooter";
import GlobalMenu from "./GlobalMenu";

function GlobalNavbar() {
  return (
    <Stack h={"100%"} justify="space-between">
      <div>
        <Center>
          {/* This will be replaced by a logo */}
          <Title order={1}>Browse</Title>
        </Center>
        {/* Main menu */}
        <GlobalMenu />
      </div>
      {/* Navbar footer which can host status indicators */}
      <GlobalFooter />
    </Stack>
  );
}

export default GlobalNavbar;
