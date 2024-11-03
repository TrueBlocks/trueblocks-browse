import { Stack, Group, Text } from "@mantine/core";
import { IconCheck } from "@tabler/icons-react";

// TODO: Implement this
export function DaemonStatus() {
  return (
    <Stack justify="flex-start" align="flex-start">
      <NodeStatus name="Erigon" />
      <NodeStatus name="Prysm" />
      <NodeStatus name="Ipfs" />
      <NodeStatus name="Scraper" />
      <NodeStatus name="Monitor" />
    </Stack>
  );
}

const NodeStatus = ({ name }: { name: string }) => (
  <Group>
    <IconCheck />
    <Text size="xs">{name}</Text>
  </Group>
);
