import { IconWashTemperature3 } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { BrowserOpenURL } from "@runtime";

// CleanButton cleans a database
export const CleanButton = ({ value, ...props }: ButtonProps) => {
  const handleClick = () => {
    BrowserOpenURL(value as unknown as string);
  };
  return <BaseButton {...props} tip={"Clean"} onClick={handleClick} icon={<IconWashTemperature3 />} />;
};
