import { ActionIcon, Button, Group, Popover, Stack, TextInput } from "@mantine/core";
import { ClipboardSetText } from "@runtime";
import { IconCopy, IconEdit } from "@tabler/icons-react";
import React, { useState, forwardRef, useCallback } from "react";

export function DataPopover({ children, editor }: { children: React.ReactNode, editor: React.ReactNode }) {
  return (
    <>
      {editor
        ?(
          <Popover withArrow>
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
  const [edit, setEdit] = useState(false);
  const submitForm = useCallback((e: React.FormEvent) => {
    e.preventDefault();
    onSubmit?.(inputValue);
    setEdit(false);
  }, [inputValue, setEdit]);
  const copy = useCallback(() => {
    ClipboardSetText(inputValue);
  }, []);

  return (
    <div ref={ref}>
      {edit
        ? (
          <form onSubmit={submitForm}>
            <Stack>
              <TextInput
                value={inputValue}
                onChange={(event) => setInputValue(event.currentTarget.value)}
              />
              <Group>
                <Button type="submit">Save</Button>
                <Button type="button" variant="outline" onClick={() => setEdit(false)}>Cancel</Button>
              </Group>
            </Stack>
          </form>
        )
        : (
          <Group>
            <div>{inputValue}</div>
            <ActionIcon onClick={() => setEdit(true)}>
              <IconEdit />
            </ActionIcon>
            <ActionIcon variant="outline" onClick={copy}>
              <IconCopy />
            </ActionIcon>
          </Group>
        )
      }
    </div>
  );
});