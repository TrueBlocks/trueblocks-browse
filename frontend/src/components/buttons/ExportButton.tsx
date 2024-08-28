import React from "react";
import { Button } from "@mantine/core";
import { BrowserOpenURL } from "@runtime";
import { IconFileExport } from "@tabler/icons-react";

export type ExportButtonProps = {
  endpoint: string;
  value: string;
  onClick?: () => void;
};

export const ExportButton = ({ endpoint, value, onClick }: ExportButtonProps) => {
  const handleClick = () => {
    BrowserOpenURL(`https://etherscan.io/${endpoint}/${value}`);
    if (onClick) {
      onClick();
    }
  };

  return (
    <Button size={"xs"} onClick={handleClick} leftSection={<IconFileExport />}>
      Export
    </Button>
  );
};
