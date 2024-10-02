import { Progress, Text } from "@mantine/core";
import { FormatterProps } from "@components";

export const LoadProgress = ({ value, value2 }: Omit<FormatterProps, "type">) => {
  const n2 = value2 as number;
  const loadedItems = Math.round((value / 100) * n2); // Calculate X

  if (value > 98) {
    return (
      <>
        <Progress value={value} color="yellow" />
        <Text size="xs" mt="xs">
          {loadedItems} of {n2}
        </Text>
      </>
    );
  }

  return (
    <>
      <Progress value={value} color="yellow" striped />
      <Text size="xs" mt="xs">
        {loadedItems} of {n2}
      </Text>
    </>
  );
};
