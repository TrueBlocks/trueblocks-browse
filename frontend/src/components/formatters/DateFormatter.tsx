import React from "react";
import { TextFormatter, FormatterProps } from ".";

export function DateFormatter({ className, value, size }: Omit<FormatterProps, "type">) {
  const parts = value.split(" ");
  return (
    <>
      <TextFormatter value={parts[0]} size={size} type="date" className={className} />
      <i>
        <TextFormatter value={parts[1] + " UTC"} size="xs" type="time" className={className} />
      </i>
    </>
  );
}
