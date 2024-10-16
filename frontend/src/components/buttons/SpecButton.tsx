import { ExploreButton, ButtonProps } from "@components";

// SpecButton opens the Unchained Index specification.
export const SpecButton = ({ value, ...props }: ButtonProps) => {
  return <ExploreButton {...props} type="url" value={value} />;
};
