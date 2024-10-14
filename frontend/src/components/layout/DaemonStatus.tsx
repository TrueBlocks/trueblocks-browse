import { Stack, Group, Text } from "@mantine/core";
import { IconCheck } from "@tabler/icons-react";

// TODO: Implement this
export function DaemonStatus() {
  return (
    <Stack justify="flex-start" align="flex-start">
      <Status name="Erigon" />
      <Status name="Prysm" />
      <Status name="Ipfs" />
      <Status name="Scraper" />
      <Status name="Monitor" />
    </Stack>
  );
}

const Status = ({ name }: { name: string }) => (
  <Group>
    <IconCheck />
    <Text size="xs">{name}</Text>
  </Group>
);
