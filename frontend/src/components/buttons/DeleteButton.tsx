import { Group } from "@mantine/core";
import { IconTrash, IconTrashX, IconArrowBackUp } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { base, app } from "@gocode/models";
import { useViewState } from "@state";

export interface DeleteButtonProps extends ButtonProps {
  isDeleted: boolean;
}

export const DeleteButton = ({ value, isDeleted, ...props }: DeleteButtonProps) => {
  const givenAddress = value as base.Address;

  const { fetchFn, modifyFn, pager } = useViewState();

  const handleDelete = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    modifyFn(
      app.ModifyData.createFrom({
        operation: "delete",
        address: givenAddress,
      })
    ).then(() => {
      fetchFn(pager.getOffset(), pager.perPage);
    });
  };

  const handleUndelete = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    modifyFn(
      app.ModifyData.createFrom({
        operation: "undelete",
        address: givenAddress,
      })
    ).then(() => {
      fetchFn(pager.getOffset(), pager.perPage);
    });
  };

  const handleRemove = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    modifyFn(
      app.ModifyData.createFrom({
        operation: "remove",
        address: givenAddress,
      })
    ).then(() => {
      fetchFn(pager.getOffset(), pager.perPage);
    });
  };

  if (isDeleted) {
    return (
      <Group justify="flex-end">
        <BaseButton tip="Remove" onClick={handleRemove} leftSection={<IconTrashX />} {...props} />
        <BaseButton tip="Undelete" onClick={handleUndelete} leftSection={<IconArrowBackUp />} {...props} />
      </Group>
    );
  }

  return (
    <Group justify="flex-end">
      <BaseButton tip="Delete" onClick={handleDelete} leftSection={<IconTrash />} {...props} />
    </Group>
  );
};
