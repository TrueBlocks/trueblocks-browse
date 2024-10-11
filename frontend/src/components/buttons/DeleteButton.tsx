import { Group } from "@mantine/core";
import { IconTrash, IconTrashX, IconArrowBackUp } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { base } from "@gocode/models";
import { useModifyFn } from "@hooks";

export interface DeleteButtonProps extends ButtonProps {
  isDeleted: boolean;
}

// DeleteButton deletes, undeletes, or removes the pointed to items. This component
// always operates on an address, but it picks up its modifyFn and fetchFn from
// the useModifyFn hook. This allows different views (MonitorsView, AbisView, etc.)
// to use delete buttons without having to know the details of how to delete items.
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
