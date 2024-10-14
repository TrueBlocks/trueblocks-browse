import { Stack } from "@mantine/core";
import { Menu, DaemonStatus } from "@components";

export const MenuPanel = () => {
  return (
    <Stack h={"100%"} justify="space-between">
      <Menu />
      <DaemonStatus />
    </Stack>
  );
}
