import React, { ReactNode } from "react";
import { Text } from "@mantine/core";

export function ViewStatus({ children }: { children: ReactNode }) {
  return <Text size="xs">{children}</Text>;
}
