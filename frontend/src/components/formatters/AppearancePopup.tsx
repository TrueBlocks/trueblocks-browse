import { forwardRef, useCallback } from "react";
import { ActionIcon } from "@mantine/core";
import { Group } from "@mantine/core";
import { IconCopy } from "@tabler/icons-react";
import { ExploreButton, PopupProps } from "@components";

export interface AppearancePopupProps extends PopupProps {
  hash: string;
}

export const AppearancePopup = forwardRef<HTMLDivElement, AppearancePopupProps>(({ hash, onClose, onCopy }, ref) => {
  const handleButtonClick = useCallback(() => {
    onClose();
  }, [onClose]);

  return (
    <Group>
      <ExploreButton value={hash} onClick={handleButtonClick} endpoint="tx" />
      <ActionIcon variant="outline" onClick={onCopy} title="Copy to clipboard">
        <IconCopy />
      </ActionIcon>
    </Group>
  );
});

AppearancePopup.displayName = "AddressPopup";
