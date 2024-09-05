import { Text, Group } from "@mantine/core";
import { Formatter } from "@components";
import { useAppState } from "@state";

export function IndexStatus() {
  const { indexes, meta } = useAppState();

  if (!indexes.items) {
    return <Text size="sm">loading indexes...</Text>;
  }

  return (
    <Group justify={"space-between"}>
      <Text size="sm">unchained index: </Text>
      <Formatter size="sm" type="int" value={indexes.nItems} />
      {" / "}
      <Formatter size="sm" type="int" value={meta.client} />
      {" / "}
      <Formatter size="sm" type="int" value={meta.finalized - meta.client} />
      {" / "}
      <Formatter size="sm" type="int" value={meta.ripe - meta.client} />{" "}
    </Group>
  );
}
