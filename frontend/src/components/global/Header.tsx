import { Group, Title, Text } from "@mantine/core";
import { useLocation } from "wouter";
import { IndexStatus } from "./";

export const Header = ({ title }: { title: string }) => {
  const [location] = useLocation();
  return (
    <Group w={"100%"} justify="space-between">
      <Title order={1}>{title}</Title>
      <Text>{`location: ${location}`}</Text>
      <IndexStatus />
    </Group>
  );
};
