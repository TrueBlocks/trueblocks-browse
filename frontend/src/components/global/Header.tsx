import React from "react";
import { Group, Title } from "@mantine/core";
import AppStatus from "./AppStatus";

const Header = ({ title }: { title: string }) => {
  return (
    <Group w={"100%"} justify="space-between">
      <Title order={1}>{title}</Title>
      <AppStatus />
    </Group>
  );
};

export default Header;
