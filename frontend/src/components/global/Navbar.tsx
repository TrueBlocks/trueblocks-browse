import React from "react";
import { Stack } from "@mantine/core";
import Menu from "./Menu";
import ServerStatus from "./ServerStatus";

function Navbar() {
  return (
    <Stack h={"100%"} justify="space-between">
      <Menu />
      <ServerStatus />
    </Stack>
  );
}

export default Navbar;
