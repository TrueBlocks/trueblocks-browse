import React, { ReactNode } from "react";

import { Popover } from "@mantine/core";

export function Popup({ children, editor }: { children: ReactNode; editor: ReactNode }) {
  return (
    <Popover withArrow>
      <Popover.Target>
        <div>{children}</div>
      </Popover.Target>
      <Popover.Dropdown>{editor}</Popover.Dropdown>
    </Popover>
  );
}
