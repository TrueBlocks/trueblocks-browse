import React, { forwardRef, useCallback } from "react";
import { ActionIcon } from "@mantine/core";
import { IconCopy } from "@tabler/icons-react";
import { ExploreButton, PopupProps } from "@components";
import { Group } from "@mantine/core";

export interface AppearancePopupProps extends PopupProps {
  hash: string;
}

export const AppearancePopup = forwardRef<HTMLDivElement, AppearancePopupProps>(({ hash, onClose, onCopy }, ref) => {
  const handleButtonClick = useCallback(() => {
    onClose();
  }, [onClose]);

  return (
    <Group>
      <ExploreButton size="sm" endpoint="tx" value={hash} onClick={handleButtonClick} />
      <ActionIcon variant="outline" onClick={onCopy} title="Copy to clipboard">
        <IconCopy />
      </ActionIcon>
    </Group>
  );
});
