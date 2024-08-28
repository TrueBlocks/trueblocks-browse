import React, { useState, forwardRef, useCallback, useEffect, useRef } from "react";
import { ActionIcon, Button, Group, Stack, TextInput } from "@mantine/core";
import { IconCopy } from "@tabler/icons-react";
import { ExploreButton, ViewButton, PopupProps } from "@components";

export interface AddressPopupProps extends PopupProps {
  address: string;
  name: string;
}

export const AddressPopup = forwardRef<HTMLDivElement, AddressPopupProps>(
  ({ name, address, onSubmit, onCopy, onClose }, ref) => {
    const [inputValue, setInputValue] = useState(name === address ? "" : name || "");

    const submitForm = useCallback(
      (e: React.FormEvent) => {
        e.preventDefault();
        onSubmit?.(inputValue);
        onClose(); // Close the popup after submitting
      },
      [inputValue, onSubmit, onClose]
    );

    const handleButtonClick = useCallback(() => {
      onClose();
    }, [onClose]);

    // Close the popup when clicking outside
    useEffect(() => {
      const handleClickOutside = (event: MouseEvent) => {
        if (ref && "current" in ref && ref.current && !ref.current.contains(event.target as Node)) {
          onClose();
        }
      };

      document.addEventListener("mousedown", handleClickOutside);
      return () => {
        document.removeEventListener("mousedown", handleClickOutside);
      };
    }, [ref, onClose]);

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
              <ExploreButton endpoint="address" value={address} onClick={handleButtonClick} />
              <ViewButton address={address} onClick={handleButtonClick} />
              <Button size="xs" type="submit">
                Save
              </Button>
              <ActionIcon variant="outline" onClick={onCopy} title="Copy to clipboard">
                <IconCopy />
              </ActionIcon>
            </Group>
          </Stack>
        </form>
      </div>
    );
  }
);
