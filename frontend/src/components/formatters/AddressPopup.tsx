import React, { forwardRef, useCallback } from "react";
import { ActionIcon, Button, Group } from "@mantine/core";
import { ClipboardSetText } from "@runtime";
import { SetLast } from "@gocode/app/App";
import { EventsEmit } from "@runtime";
import { IconCopy, IconExternalLink, IconLink } from "@tabler/icons-react";
import { messages } from "@gocode/models";
import { ExploreButton, MonitorButton } from ".";

type PopupProps = {
  address: () => string;
};

export const AddressPopup = forwardRef<HTMLDivElement, PopupProps>(({ address }, ref) => {
  const copy = useCallback(() => {
    ClipboardSetText(address());
  }, []);

  return (
    <div ref={ref}>
      <Group>
        <ExploreButton address={address()} />
        <MonitorButton address={address()} />
        <ActionIcon variant="outline" onClick={copy} title="Copy to clipboard">
          <IconCopy />
        </ActionIcon>
      </Group>
    </div>
  );
});
