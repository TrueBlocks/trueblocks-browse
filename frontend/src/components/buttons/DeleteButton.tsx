import { Group } from "@mantine/core";
import { IconTrash } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { base } from "@gocode/models";
import { useModifyFn } from "@hooks";

export const DeleteButton = ({ value, ...props }: ButtonProps) => {
  const address = value as base.Address;
  const { deleteItem } = useModifyFn(address);
  return (
    <Group justify="flex-end">
      <BaseButton {...props} tip="Delete" onClick={deleteItem} icon={<IconTrash />} />
    </Group>
  );
};
