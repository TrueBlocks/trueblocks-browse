import { IconBooks } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { BrowserOpenURL } from "@runtime";

// PublishButton shares values with others
export const PublishButton = ({ value, ...props }: ButtonProps) => {
  const handleClick = () => {
    BrowserOpenURL(value as unknown as string);
  };
  return <BaseButton {...props} tip={"Publish"} onClick={handleClick} icon={<IconBooks />} />;
};
