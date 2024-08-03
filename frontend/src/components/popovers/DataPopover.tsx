import React, { ReactNode } from "react";

import { Popover } from "@mantine/core";

export function DataPopover({ children, editor }: { children: ReactNode; editor: ReactNode }) {
  return (
    <>
      {editor ? (
        <Popover withArrow width="target">
          <Popover.Target>
            <div>{children}</div>
          </Popover.Target>
          <Popover.Dropdown>{editor}</Popover.Dropdown>
        </Popover>
      ) : (
        children
      )}
    </>
  );
}
