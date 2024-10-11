import { forwardRef } from "react";
import { Group } from "@mantine/core";
import { ExploreButton, CopyButton, PopupProps } from "@components";

export interface AppearancePopupProps extends PopupProps {
  hash: string;
}

export const AppearancePopup = forwardRef<HTMLDivElement, AppearancePopupProps>(({ hash, onCopy, onClose }, ref) => {
  return (
    <Group>
      <ExploreButton value={hash} onClose={onClose} />
      <CopyButton value={hash} onClick={onCopy} onClose={onClose} />
    </Group>
  );
});

AppearancePopup.displayName = "AddressPopup";
