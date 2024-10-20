import { IconSquarePlus } from "@tabler/icons-react";
import { ButtonProps, BaseButton } from "@components";

// AddButton cleans a database
export const AddButton = ({ value, ...props }: ButtonProps) => {
  return <BaseButton {...props} tip={"Add"} value={value} leftSection={<IconSquarePlus />} />;
};
