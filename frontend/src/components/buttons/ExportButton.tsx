import { Button } from "@mantine/core";
import { IconFileExport } from "@tabler/icons-react";
import { ButtonProps } from "@components";
import { ExportToCsv } from "@gocode/app/App";

export const ExportButton = ({ value, noText, size, onClick }: ButtonProps) => {
  const handleClick = () => {
    ExportToCsv(value);
    if (onClick) {
      onClick();
    }
  };

  return (
    <Button size={size} onClick={handleClick} leftSection={<IconFileExport />}>
      {noText ? null : "Export"}
    </Button>
  );
};
