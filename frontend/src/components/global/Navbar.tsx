import React from "react";
import { Stack } from "@mantine/core";
import { Menu, ServerStatus } from "./";

export function Navbar() {
  return (
    <Stack h={"100%"} justify="space-between">
      <Menu />
      <ServerStatus />
    </Stack>
  );
}
