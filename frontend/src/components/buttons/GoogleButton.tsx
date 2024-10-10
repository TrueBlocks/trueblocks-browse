import { ExploreButton, ButtonProps } from "@components";

export const GoogleButton = ({ value, ...props }: ButtonProps) => {
  return <ExploreButton type="google" value={value} {...props} />;
};
