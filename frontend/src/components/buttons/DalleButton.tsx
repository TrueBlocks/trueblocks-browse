import { ExploreButton, ButtonProps } from "@components";

// DalleButton opens up the DalleDress explorer website.
export const DalleButton = ({ value, ...props }: ButtonProps) => {
  return <ExploreButton {...props} type="dalle" value={value} />;
};
