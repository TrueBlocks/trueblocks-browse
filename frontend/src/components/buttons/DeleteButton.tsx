import { Group } from "@mantine/core";
import { IconTrash, IconTrashX, IconArrowBackUp } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { base } from "@gocode/models";
import { useModifyFn } from "@hooks";

export interface DeleteButtonProps extends ButtonProps {
  isDeleted: boolean;
}

export const DeleteButton = ({ value, isDeleted, ...props }: DeleteButtonProps) => {
  const { deleteItem, undeleteItem, removeItem } = useModifyFn(value as base.Address);

  if (isDeleted) {
    return (
      <Group justify="flex-end">
        <BaseButton {...props} tip="Remove" onClick={removeItem} leftSection={<IconTrashX />} />
        <BaseButton {...props} tip="Undelete" onClick={undeleteItem} leftSection={<IconArrowBackUp />} />
      </Group>
    );
  }

  return (
    <Group justify="flex-end">
      <BaseButton {...props} tip="Delete" onClick={deleteItem} leftSection={<IconTrash />} />
    </Group>
  );
};
