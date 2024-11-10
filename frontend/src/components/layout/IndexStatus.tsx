import { Text, Group } from "@mantine/core";
import { Formatter } from "@components";
import { useAppState } from "@state";

export function IndexStatus() {
  const { info } = useAppState();
  const { indexes } = useAppState();

  if (!indexes.items) {
    return <Text size="sm">loading indexes...</Text>;
  }

  return (
    <Group justify={"space-between"}>
      <Text size="sm">unchained index: </Text>
      <Formatter size="sm" type="int" value={indexes.nItems} />
      {" / "}
      <Formatter size="sm" type="int" value={info.meta.client} />
      {" / "}
      <Formatter size="sm" type="int" value={info.meta.finalized - info.meta.client} />
      {" / "}
      <Formatter size="sm" type="int" value={info.meta.ripe - info.meta.client} />{" "}
    </Group>
  );
}
