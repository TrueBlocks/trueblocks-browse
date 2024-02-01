import React from "react";
import { Flex, Space, Text } from "@mantine/core";
import { IconCheck } from "@tabler/icons-react";

function GlobalFooter() {
  return (
    // Status indicators mock
    <Flex justify="center" align="center">
      <IconCheck />
      <Text size="xs">Node</Text>
      <Space />
      <IconCheck />
      <Text size="xs">Scraper</Text>
    </Flex>
  );
}

export default GlobalFooter;
