import { useState, useEffect } from "react";
import { ActionIcon } from "@mantine/core";
import { IconTrash, IconTrashX, IconArrowBackUp } from "@tabler/icons-react";
import { ButtonProps } from "@components";
import { ModifyName } from "@gocode/app/App";
import { base } from "@gocode/models";
import { useAppState, useViewState } from "@state";

export interface CrudButtonProps extends ButtonProps {
  isDeleted: boolean;
}

//  crudOperation(route: Route, selected: number, op: string): void;

export const CrudButton = ({ value, size, isDeleted, onClick }: CrudButtonProps) => {
  const [address, setAddress] = useState<base.Address>(value as unknown as base.Address);
  const { fetchNames, crudOperation } = useAppState();
  const { pager } = useViewState();
  const { selected } = pager;

  useEffect(() => {
    setAddress(value as unknown as base.Address);
  }, [value]);

  const handleDelete = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    ModifyName("delete", address).then(() => {});
    fetchNames(pager.getOffset(), pager.perPage);
    crudOperation("names", selected, "delete");
    if (onClick) {
      onClick();
    }
  };

  const handleUndelete = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    ModifyName("undelete", address).then(() => {});
    fetchNames(pager.getOffset(), pager.perPage);
    crudOperation("names", selected, "undelete");
    if (onClick) {
      onClick();
    }
  };

  const handleRemove = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    ModifyName("remove", address).then(() => {});
    fetchNames(pager.getOffset(), pager.perPage);
    // setRecord(0);
    // setRecord(selected);
    crudOperation("names", selected, "remove");
    if (onClick) {
      onClick();
    }
  };

  if (isDeleted) {
    return (
      <>
        {selected}
        <ActionIcon c="red" size={size} variant="outline" onClick={handleUndelete} title="Delete">
          <IconArrowBackUp />
        </ActionIcon>
        <ActionIcon c="red" size={size} variant="outline" onClick={handleRemove} title="Delete">
          <IconTrashX />
        </ActionIcon>
      </>
    );
  }

  return (
    <>
      {selected}
      <ActionIcon size={size} variant="outline" onClick={handleDelete} title="Delete">
        <IconTrash />
      </ActionIcon>
    </>
  );
};
