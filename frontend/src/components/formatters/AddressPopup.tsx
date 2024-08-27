import React, { useState, forwardRef, useCallback, useEffect, useRef } from "react";
import { ActionIcon, Button, Group, Stack, TextInput } from "@mantine/core";
import { ClipboardSetText } from "@runtime";
import { IconCopy } from "@tabler/icons-react";
import { ExploreButton, MonitorButton } from ".";
import { useHotkeys } from "react-hotkeys-hook";

type PopupProps = {
  name?: string;
  address: string;
  onSubmit?: (value: string) => void;
  onClose?: () => void;
};

export const AddressPopup = forwardRef<HTMLDivElement, PopupProps>(
  ({ name, address, onSubmit, onClose = () => {} }, ref) => {
    const [inputValue, setInputValue] = useState(name === address ? "" : name || "");

    const submitForm = useCallback(
      (e: React.FormEvent) => {
        e.preventDefault();
        onSubmit?.(inputValue);
        onClose(); // Close the popup after submitting
      },
      [inputValue, onSubmit, onClose]
    );

    const copy = useCallback(() => {
      ClipboardSetText(inputValue).then(() => {
        onClose(); // Close the popup after copying
      });
    }, [inputValue, onClose]);

    const handleButtonClick = useCallback(() => {
      onClose(); // Close the popup when either button is clicked
    }, [onClose]);

    // Close the popup when the escape key is pressed
    useHotkeys("escape", (event) => {
      event.preventDefault(); // Prevent the escape key from bubbling up
      event.stopPropagation(); // Stop the escape key event from propagating further
      onClose(); // Close the popup
    });

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
              <ExploreButton address={address} onClick={handleButtonClick} />
              <MonitorButton address={address} onClick={handleButtonClick} />
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
  }
);
