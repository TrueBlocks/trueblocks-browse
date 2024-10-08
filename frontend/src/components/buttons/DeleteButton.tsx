import { useState, useEffect } from "react";
import { ActionIcon, Group } from "@mantine/core";
import { IconTrash, IconTrashX, IconArrowBackUp } from "@tabler/icons-react";
import { ButtonProps } from "@components";
import { base, app } from "@gocode/models";
import { useViewState } from "@state";

export interface DeleteButtonProps extends Omit<Omit<ButtonProps, "onClick">, "size"> {
  isDeleted: boolean;
}

export const DeleteButton = ({ value, isDeleted }: DeleteButtonProps) => {
  const [address, setAddress] = useState<base.Address>(value as unknown as base.Address);
  const { fetchFn, modifyFn, pager } = useViewState();

  useEffect(() => {
    setAddress(value as unknown as base.Address);
  }, [value]);

  const handleDelete = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    const op = "delete";
    const modData = app.ModifyData.createFrom({
      operation: op,
      address: address,
      value: "",
    });
    modifyFn(modData).then(() => {
      fetchFn(pager.getOffset(), pager.perPage);
    });
  };

  const handleUndelete = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    const op = "undelete";
    const modData = app.ModifyData.createFrom({
      operation: op,
      address: address,
      value: "",
    });
    modifyFn(modData).then(() => {
      fetchFn(pager.getOffset(), pager.perPage);
    });
  };

  const handleRemove = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    const op = "remove";
    const modData = app.ModifyData.createFrom({
      operation: op,
      address: address,
      value: "",
    });
    modifyFn(modData).then(() => {
      fetchFn(pager.getOffset(), pager.perPage);
    });
  };

  const size = "sm";
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
