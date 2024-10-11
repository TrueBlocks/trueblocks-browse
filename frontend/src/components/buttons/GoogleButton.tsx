import { ExploreButton, ButtonProps } from "@components";

// GoogleButton opens a browser window to search google for a given address, but
// it only shows results that are not blockchain explorers. Useful to study
// an address unrelated to its transactional history.
export const GoogleButton = ({ value, ...props }: ButtonProps) => {
  return <ExploreButton {...props} type="google" value={value} />;
};
