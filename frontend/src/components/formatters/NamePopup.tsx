import React, { useState, forwardRef, useCallback } from "react";

import { ActionIcon, Button, Group, Stack, TextInput } from "@mantine/core";
import { ClipboardSetText } from "@runtime";
import { IconCopy } from "@tabler/icons-react";

type AddressNameEditorProps = {
  name?: () => any;
  onSubmit?: (value: string) => void;
};

export const NamePopup = forwardRef<HTMLDivElement, AddressNameEditorProps>(({ name, onSubmit }, ref) => {
  const [inputValue, setInputValue] = useState(String(name?.() || ""));
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
          <TextInput value={inputValue} onChange={(event) => setInputValue(event.currentTarget.value)} autoFocus />
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
