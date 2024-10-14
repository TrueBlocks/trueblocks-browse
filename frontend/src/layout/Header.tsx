import { Group, Title, Text } from "@mantine/core";
import { IndexStatus, ChainSelector } from "@components";

export const Header = ({ title }: { title: string }) => {
  return (
    <Group w={"100%"} justify="space-between">
      <Title order={1}>{title}</Title>
      <ChainSelector />
      <IndexStatus />
    </Group>
  );
};
