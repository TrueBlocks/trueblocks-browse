import { ExploreButton, ButtonProps } from "@components";

// ShareButton cleans a database
export const ShareButton = ({ value, ...props }: ButtonProps) => {
  return <ExploreButton {...props} type="url" value={value} />;
};
