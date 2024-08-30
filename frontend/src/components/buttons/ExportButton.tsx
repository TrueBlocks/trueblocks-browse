import React from "react";
import { Button } from "@mantine/core";
import { ExportToCsv } from "@gocode/app/App";
import { IconFileExport } from "@tabler/icons-react";
import { ButtonProps } from "@components";

export interface ExportButtonProps extends ButtonProps {}

export const ExportButton = ({ value, noText, size, onClick }: ExportButtonProps) => {
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
