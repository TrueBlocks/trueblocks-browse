import { ActionIcon, Button, Group, Popover, Stack, TextInput } from "@mantine/core";
import { BrowserOpenURL, ClipboardSetText } from "@runtime";
import { IconCopy, IconEdit, IconExternalLink } from "@tabler/icons-react";
import React, { useState, forwardRef, useCallback } from "react";

export function DataPopover({ children, editor }: { children: React.ReactNode, editor: React.ReactNode }) {
  return (
    <>
      {editor
        ?(
          <Popover withArrow width="target">
            <Popover.Target>
              <div>
                {children}
              </div>
            </Popover.Target>
            <Popover.Dropdown>
              {editor}
            </Popover.Dropdown>
          </Popover>
        )
        : children
      }
    </>
  );
}

type DataTableEditor = {
  value?: () => any
  onSubmit?: (value: string) => void
}
export const DataTableStringEditor = forwardRef<HTMLDivElement, DataTableEditor>(({ value, onSubmit }, ref) => {
  const [inputValue, setInputValue] = useState(String(value?.() || ""));
  const submitForm = useCallback((e: React.FormEvent) => {
    e.preventDefault();
    onSubmit?.(inputValue);
  }, [inputValue]);
  const copy = useCallback(() => {
    ClipboardSetText(inputValue);
  }, []);

  return (
    <div ref={ref}>
      <form onSubmit={submitForm}>
        <Stack>
          <TextInput
            value={inputValue}
            onChange={(event) => setInputValue(event.currentTarget.value)}
            autoFocus
          />
          <Group>
            <ActionIcon variant="outline" onClick={copy} title="Copy to clipboard">
              <IconCopy />
            </ActionIcon>
            <Button type="submit">Save</Button>
          </Group>
        </Stack>
      </form>
    </div>
  );
});

type DataTableViewer = {
  address: () => string
};
export const DataTableViewInEtherscan = forwardRef<HTMLDivElement, DataTableViewer>(({ address }, ref) => {
  const copy = useCallback(() => {
    ClipboardSetText(address());
  }, []);
  return (
    <div ref={ref}>
      <Group>
        <Button onClick={() => BrowserOpenURL(`https://etherscan.io/address/${address()}`)} leftSection={<IconExternalLink />}>
          View on Etherscan
        </Button>
        <ActionIcon variant="outline" onClick={copy} title="Copy to clipboard">
          <IconCopy />
        </ActionIcon>
      </Group>
    </div>
  );
});