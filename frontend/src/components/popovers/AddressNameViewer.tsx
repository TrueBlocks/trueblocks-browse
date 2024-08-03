import React, { forwardRef, useCallback } from "react";

import { ActionIcon, Button, Group, Popover, Stack, TextInput } from "@mantine/core";
import { BrowserOpenURL, ClipboardSetText } from "@runtime";
import { IconCopy, IconExternalLink } from "@tabler/icons-react";

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
          onClick={() => BrowserOpenURL(`https://etherscan.io/address/${address()}`)}
          leftSection={<IconExternalLink />}
        >
          View on Etherscan
        </Button>
        <ActionIcon variant="outline" onClick={copy} title="Copy to clipboard">
          <IconCopy />
        </ActionIcon>
      </Group>
    </div>
  );
});
