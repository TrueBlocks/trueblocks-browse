import { IconLink } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { LoadAddress } from "@gocode/app/App";

// ViewButton opens the history page for a given address.
export const ViewButton = ({ value, ...props }: ButtonProps) => {
  const addressStr = value as string;
  const handleClick = () => {
    LoadAddress(addressStr).then(() => {});
  };

  return <BaseButton {...props} tip="View" onClick={handleClick} icon={<IconLink />} />;
};
