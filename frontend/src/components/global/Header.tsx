import React from "react";
import { Group, Title, Text } from "@mantine/core";
import { IndexStatus } from "./";
import { useLocation } from "wouter";

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
