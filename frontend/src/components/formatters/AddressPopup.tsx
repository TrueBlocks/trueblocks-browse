import { useState, forwardRef, useCallback, useEffect } from "react";
import { Button, Group, Stack, TextInput } from "@mantine/core";
import { ExploreButton, ViewButton, PopupProps, CopyButton, DalleButton, GoogleButton } from "@components";

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

    const size = "sm";
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
              <ExploreButton noText value={address} onClick={handleButtonClick} />
              <DalleButton noText value={address} onClick={handleButtonClick} />
              <GoogleButton noText value={address} onClick={handleButtonClick} />
              <ViewButton value={address} onClick={handleButtonClick} />
              <Button size={size} type="submit">
                Save
              </Button>
              <CopyButton onClick={onCopy} />
            </Group>
          </Stack>
        </form>
      </div>
    );
  }
);

AddressPopup.displayName = "AddressPopup";
