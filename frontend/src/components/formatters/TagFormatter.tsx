import React from "react";
import { FormatterProps } from "./Formatter"
import { Badge } from "@mantine/core";

export function TagFormatter({ className, value, size }: Omit<FormatterProps, "type">) {
  return (
    <Badge variant="outline" onClick={() => console.log("Tag clicked:", value)} size={size} className={className}>
      {value}
    </Badge>
  );
}