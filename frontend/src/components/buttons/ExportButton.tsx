import { IconFileExport } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { ExportAddress } from "@gocode/app/App";
import { base } from "@gocode/models";

// ExportButton exports an address's history to a .csv file.
export const ExportButton = ({ value, ...props }: ButtonProps) => {
  const handleClick = () => {
    ExportAddress(value as base.Address);
  };

  return <BaseButton {...props} tip="Export" onClick={handleClick} icon={<IconFileExport />} />;
};
