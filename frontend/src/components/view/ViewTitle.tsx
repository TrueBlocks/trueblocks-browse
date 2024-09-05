import { Title } from "@mantine/core";
import { useViewName } from "@hooks";

export function ViewTitle(): JSX.Element {
  return <Title order={3}>{useViewName()}</Title>;
}
