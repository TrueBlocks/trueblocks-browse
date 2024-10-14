import { Title } from "@mantine/core";
import { useViewName } from "@hooks";

export const ViewTitle = () => {
  return <Title order={3}>{useViewName()}</Title>;
};
