import { useState, useCallback, useEffect } from "react";
import { Button, Group, Stack, TextInput } from "@mantine/core";
import { ExploreButton, ViewButton, PopupProps, CopyButton, DalleButton, GoogleButton } from "@components";

export interface AddressPopupProps extends PopupProps {
  address: string;
  name: string;
}

export const AddressPopup = ({ name, address, onSubmit, onCopy, onClose }: AddressPopupProps) => {
  const [inputValue, setInputValue] = useState(name === address ? "" : name || "");

  const submitForm = useCallback(
    (e: React.FormEvent) => {
      e.preventDefault();
      onSubmit?.(inputValue);
      onClose();
    },
    [inputValue, onSubmit, onClose]
  );

  useEffect(() => {
    const handleClickOutside = () => {
      onClose();
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [onClose]);

  const size = "sm";
  return (
    <div>
      <form onSubmit={submitForm}>
        <Stack bg="white">
          <TextInput
            placeholder="name this address"
            value={inputValue}
            onChange={(event) => setInputValue(event.currentTarget.value)}
            autoFocus
          />
          <Group>
            <ExploreButton value={address} onClose={onClose} />
            <DalleButton value={address} onClose={onClose} />
            <GoogleButton value={address} onClose={onClose} />
            <ViewButton value={address} onClose={onClose} />
            <Button size={size} type="submit">
              Save
            </Button>
            <CopyButton value={address} onClick={onCopy} onClose={onClose} />
          </Group>
        </Stack>
      </form>
    </div>
  );
};

AddressPopup.displayName = "AddressPopup";
