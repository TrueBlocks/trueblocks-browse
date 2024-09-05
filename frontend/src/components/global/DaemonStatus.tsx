import { Flex, Space, Text } from "@mantine/core";
import { IconCheck } from "@tabler/icons-react";

export function DaemonStatus() {
  return (
    <Flex justify="center" align="center">
      <IconCheck />
      <Text size="xs">Node</Text>
      <Space />
      <IconCheck />
      <Text size="xs">Scraper</Text>
    </Flex>
  );
}
