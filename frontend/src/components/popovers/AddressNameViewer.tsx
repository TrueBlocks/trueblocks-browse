import React, { forwardRef, useCallback } from "react";

import { ActionIcon, Button, Group, Popover, Stack, TextInput } from "@mantine/core";
import { BrowserOpenURL, ClipboardSetText } from "@runtime";
import { SetLast } from "@gocode/app/App";
import { EventsEmit } from "@runtime";
import { IconCopy, IconExternalLink } from "@tabler/icons-react";
import { messages } from "@gocode/models";

type AddressNameViewerProps = {
  address: () => string;
};

export const AddressNameViewer = forwardRef<HTMLDivElement, AddressNameViewerProps>(({ address }, ref) => {
  const copy = useCallback(() => {
    ClipboardSetText(address());
  }, []);
  return (
    <div ref={ref}>
      <Group>
        <Button
          size={"xs"}
          onClick={() => BrowserOpenURL(`https://etherscan.io/address/${address()}`)}
          leftSection={<IconExternalLink />}
        >
          Explore
        </Button>
        <Button
          size={"xs"}
          onClick={(e) => {
            e.preventDefault();
            SetLast("route", `/history/${address()}`);
            EventsEmit(messages.Message.NAVIGATE, {
              route: `/history/${address()}`,
            });
          }}
          leftSection={<IconExternalLink />}
        >
          Monitor
        </Button>
        <ActionIcon variant="outline" onClick={copy} title="Copy to clipboard">
          <IconCopy />
        </ActionIcon>
      </Group>
    </div>
  );
});
