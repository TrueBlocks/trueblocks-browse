import { ExploreButton, ButtonProps } from "@components";

export const DalleButton = ({ value, ...props }: ButtonProps) => {
  return <ExploreButton {...props} type="dalle" value={value} />;
};
