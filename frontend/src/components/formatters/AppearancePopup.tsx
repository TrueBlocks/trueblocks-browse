import { Group } from "@mantine/core";
import { ExploreButton, CopyButton, PopupProps } from "@components";

export interface AppearancePopupProps extends PopupProps {
  hash: string;
}

export const AppearancePopup = ({ hash, onClose }: AppearancePopupProps) => {
  return (
    <Group bg="white">
      <ExploreButton value={hash} onClose={onClose} />
      <CopyButton value={hash} onClose={onClose} />
    </Group>
  );
};

AppearancePopup.displayName = "AppearancePopup";
