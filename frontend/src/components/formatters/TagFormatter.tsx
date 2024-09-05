import { Badge } from "@mantine/core";
import { FormatterProps } from "./Formatter";

export function TagFormatter({ className, value, size }: Omit<FormatterProps, "type">) {
  return (
    <Badge variant="outline" onClick={() => console.log("Tag clicked:", value)} size={size} className={className}>
      {value}
    </Badge>
  );
}
