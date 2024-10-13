import { IconLink } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { GoToHistory } from "@gocode/app/App";
import { base } from "@gocode/models";

// ViewButton opens the history page for a given address.
export const ViewButton = ({ value, ...props }: ButtonProps) => {
  const address = value as base.Address;
  const icon = <IconLink />;

  const handleClick = () => {
    GoToHistory(address).then(() => {});
  };

  return <BaseButton {...props} tip="View" onClick={handleClick} leftSection={icon} />;
};
