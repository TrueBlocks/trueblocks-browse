import React, { ReactNode, useMemo } from "react";
import { useEnvironment } from "@hooks";
import { Stack, Text, TextProps } from "@mantine/core";
import { knownTypes } from "./Formatter";

type TextFormatterProps = {
  className?: string,
  type: knownTypes,
  value: any,
  size?: TextProps["size"];
}

export function TextFormatter({ className, type, value, size }: TextFormatterProps) {
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