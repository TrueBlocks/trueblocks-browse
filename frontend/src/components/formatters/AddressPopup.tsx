import React, { forwardRef, useCallback } from "react";
import { ActionIcon, Button, Group } from "@mantine/core";
import { ClipboardSetText } from "@runtime";
import { IconCopy, IconExternalLink } from "@tabler/icons-react";
import { ExploreButton, MonitorButton } from ".";

type PopupProps = {
  address: string;
};

export const AddressPopup = forwardRef<HTMLDivElement, PopupProps>(({ address }, ref) => {
  const copy = useCallback(() => {
    ClipboardSetText(address);
  }, []);
  return (
    <div ref={ref}>
      <Group>
        <ExploreButton address={address} />
        <MonitorButton address={address} />
        <ActionIcon variant="outline" onClick={copy} title="Copy to clipboard">
          <IconCopy />
        </ActionIcon>
      </Group>
    </div>
  );
});
