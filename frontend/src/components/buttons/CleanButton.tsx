import { ExploreButton, ButtonProps } from "@components";

// CleanButton cleans a database
export const CleanButton = ({ value, ...props }: ButtonProps) => {
  return <ExploreButton {...props} type="url" value={value} />;
};
