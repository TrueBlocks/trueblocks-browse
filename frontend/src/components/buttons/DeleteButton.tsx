import { useState, useEffect } from "react";
import { ActionIcon, Group } from "@mantine/core";
import { IconTrash, IconTrashX, IconArrowBackUp } from "@tabler/icons-react";
import { ButtonProps } from "@components";
import { base, app } from "@gocode/models";
import { useAppState, useViewState } from "@state";

export interface DeleteButtonProps extends ButtonProps {
  isDeleted: boolean;
}

export const DeleteButton = ({ value, size, isDeleted, onClick }: DeleteButtonProps) => {
  const [address, setAddress] = useState<base.Address>(value as unknown as base.Address);
  const { deleteOperation } = useAppState();
  const { route, fetchFn, modifyFn, pager } = useViewState();
  const { selected } = pager;

  useEffect(() => {
    setAddress(value as unknown as base.Address);
  }, [value]);

  const handleDelete = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    const op = "delete";
    deleteOperation(route, selected, op);
    if (modifyFn) {
      const modData = app.ModifyData.createFrom({
        operation: op,
        address: address,
        value: "",
      });
      modifyFn(modData).then(() => {
        fetchFn(pager.getOffset(), pager.perPage);
      });
    }
    if (onClick) {
      onClick();
    }
  };

  const handleUndelete = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    const op = "undelete";
    deleteOperation(route, selected, op);
    if (modifyFn) {
      const modData = app.ModifyData.createFrom({
        operation: op,
        address: address,
        value: "",
      });
      modifyFn(modData).then(() => {
        fetchFn(pager.getOffset(), pager.perPage);
      });
    }
    if (onClick) {
      onClick();
    }
  };

  const handleRemove = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    const op = "remove";
    deleteOperation(route, selected, op);
    if (modifyFn) {
      const modData = app.ModifyData.createFrom({
        operation: op,
        address: address,
        value: "",
      });
      modifyFn(modData).then(() => {
        fetchFn(pager.getOffset(), pager.perPage);
      });
    }
    if (onClick) {
      onClick();
    }
  };

  if (isDeleted) {
    return (
      <Group justify="flex-end">
        <ActionIcon size={size} onClick={handleRemove} title="Remove">
          <IconTrashX />
        </ActionIcon>
        <ActionIcon size={size} onClick={handleUndelete} title="Undelete">
          <IconArrowBackUp />
        </ActionIcon>
      </Group>
    );
  }

  return (
    <Group justify="flex-end">
      <ActionIcon size={size} onClick={handleDelete} title="Delete">
        <IconTrash />
      </ActionIcon>
    </Group>
  );
};
