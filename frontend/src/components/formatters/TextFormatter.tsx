import { Stack, Text } from "@mantine/core";
import { useEnvironment } from "@hooks";
import { FormatterProps } from "./Formatter";

export function TextFormatter({ className, type, value, size }: FormatterProps) {
  const debug = useEnvironment("TB_DEBUG_DISPLAY");

  const inner = (
    <Text className={className} size={size}>
      {value}
    </Text>
  );

  if (debug == "verbose") {
    return (
      <Stack gap={0}>
        {inner}
        <Text size="xs">{type}</Text>
      </Stack>
    );
  }
  return inner;
}
