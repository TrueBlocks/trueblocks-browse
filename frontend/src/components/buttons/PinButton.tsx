import { IconPin } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { BrowserOpenURL } from "@runtime";

// PinButton pins the index or manifest to the Unchained Index
export const PinButton = ({ value, ...props }: ButtonProps) => {
  const handleClick = () => {
    BrowserOpenURL(value as unknown as string);
  };
  return <BaseButton {...props} tip={"Pin"} onClick={handleClick} icon={<IconPin />} />;
};
