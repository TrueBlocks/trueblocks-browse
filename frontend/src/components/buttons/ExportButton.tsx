import { IconFileExport } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { ExportToCsv } from "@gocode/app/App";
import { base } from "@gocode/models";

export const ExportButton = ({ value, ...props }: ButtonProps) => {
  const icon = <IconFileExport />;

  const handleClick = () => {
    ExportToCsv(value as base.Address);
  };

  return <BaseButton {...props} tip="Export" onClick={handleClick} leftSection={icon} />;
};
