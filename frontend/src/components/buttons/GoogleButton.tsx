import { ExploreButton, ButtonProps } from "@components";

export const GoogleButton = ({ value, ...props }: ButtonProps) => {
  return <ExploreButton {...props} type="google" value={value} />;
};
