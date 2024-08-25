import React, { useState, forwardRef, useCallback } from "react";
import { ActionIcon, Button, Group, Stack, TextInput } from "@mantine/core";
import { ClipboardSetText } from "@runtime";
import { IconCopy } from "@tabler/icons-react";
import { ExploreButton, MonitorButton } from ".";

type PopupProps = {
  name?: string;
  address: string;
  onSubmit?: (value: string) => void;
};

export const AddressPopup = forwardRef<HTMLDivElement, PopupProps>(({ name, address, onSubmit }, ref) => {
  const [inputValue, setInputValue] = useState(name === address ? "" : name || "");
  const submitForm = useCallback(
    (e: React.FormEvent) => {
      e.preventDefault();
      onSubmit?.(inputValue);
    },
    [inputValue]
  );

  const copy = useCallback(() => {
    ClipboardSetText(inputValue);
  }, []);

  return (
    <div ref={ref}>
      <form onSubmit={submitForm}>
        <Stack>
          <TextInput
            placeholder="name this address"
            value={inputValue}
            onChange={(event) => setInputValue(event.currentTarget.value)}
            autoFocus
          />
          <Group>
            <ExploreButton address={address} />
            <MonitorButton address={address} />
            <Button size="xs" type="submit">
              Save
            </Button>
            <ActionIcon variant="outline" onClick={copy} title="Copy to clipboard">
              <IconCopy />
            </ActionIcon>
          </Group>
        </Stack>
      </form>
    </div>
  );
});
