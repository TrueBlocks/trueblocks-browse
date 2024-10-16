import { ExploreButton, ButtonProps } from "@components";

// PinButton pins the index
export const PinButton = ({ value, ...props }: ButtonProps) => {
  return <ExploreButton {...props} type="url" value={value} />;
};
