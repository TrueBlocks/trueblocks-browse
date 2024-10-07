import { ActionIcon } from "@mantine/core";
import { IconCopy } from "@tabler/icons-react";

// TODO: Should use ButtonProps
export const CopyButton = ({ onClick }: { onClick: () => void }) => {
  const size = "sm";
  return (
    <ActionIcon size={size} onClick={onClick} title="Copy to clipboard">
      <IconCopy />
    </ActionIcon>
  );
};
