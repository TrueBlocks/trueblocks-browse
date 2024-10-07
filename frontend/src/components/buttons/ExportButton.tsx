import { Button, ActionIcon } from "@mantine/core";
import { IconFileExport } from "@tabler/icons-react";
import { ButtonProps } from "@components";
import { ExportToCsv } from "@gocode/app/App";

export const ExportButton = ({ value, noText, onClick }: Omit<ButtonProps, "size">) => {
  const handleClick = () => {
    ExportToCsv(value);
    if (onClick) {
      onClick();
    }
  };

  const size = "sm";
  if (noText) {
    return (
      <ActionIcon size={size} onClick={handleClick} title="Export">
        <IconFileExport />
      </ActionIcon>
    );
  }

  return (
    <Button size={size} onClick={handleClick} leftSection={<IconFileExport />}>
      {"Export"}
    </Button>
  );
};
