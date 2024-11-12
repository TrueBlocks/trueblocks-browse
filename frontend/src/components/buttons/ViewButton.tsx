import { IconLink } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { base } from "@gocode/models";
import { useAppState } from "@state";

// ViewButton opens the history page for a given address.
export const ViewButton = ({ value, ...props }: ButtonProps) => {
  const { loadAddress } = useAppState();
  const handleClick = () => {
    loadAddress(value as base.Address);
  };

  return <BaseButton {...props} tip="View" onClick={handleClick} icon={<IconLink />} />;
};
