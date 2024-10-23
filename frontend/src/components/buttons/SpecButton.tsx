import { IconBlockquote } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { BrowserOpenURL } from "@runtime";

// SpecButton opens the Unchained Index specification.
export const SpecButton = ({ value, ...props }: ButtonProps) => {
  const handleClick = () => {
    BrowserOpenURL(value as unknown as string);
  };
  return <BaseButton {...props} tip={"Spec"} onClick={handleClick} icon={<IconBlockquote />} />;
};
