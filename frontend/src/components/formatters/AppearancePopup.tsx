import { forwardRef, useCallback } from "react";
import { Group } from "@mantine/core";
import { ExploreButton, CopyButton, PopupProps } from "@components";

export interface AppearancePopupProps extends PopupProps {
  hash: string;
}

export const AppearancePopup = forwardRef<HTMLDivElement, AppearancePopupProps>(({ hash, onCopy, onClose }, ref) => {
  const handleButtonClick = useCallback(() => {
    onClose();
  }, [onClose]);

  return (
    <Group>
      <ExploreButton value={hash} onClick={handleButtonClick} />
      <CopyButton onClick={onCopy} />
    </Group>
  );
});

AppearancePopup.displayName = "AddressPopup";
