import { ExploreButton, ButtonProps } from "@components";

// PublishButton cleans a database
export const PublishButton = ({ value, ...props }: ButtonProps) => {
  return <ExploreButton {...props} type="url" value={value} />;
};
