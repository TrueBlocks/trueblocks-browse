import { IconLink } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { GoToAddress } from "@gocode/app/App";
import { base } from "@gocode/models";

// ViewButton opens the history page for a given address.
export const ViewButton = ({ value, ...props }: ButtonProps) => {
  const address = value as base.Address;

  const handleClick = () => {
    GoToAddress(address).then(() => {});
  };

  return <BaseButton {...props} tip="View" onClick={handleClick} icon={<IconLink />} />;
};
