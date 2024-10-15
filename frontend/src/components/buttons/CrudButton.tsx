import { Group } from "@mantine/core";
import { ButtonProps, EditButton, DeleteButton, UndeleteButton, RemoveButton } from "@components";

export interface CrudButtonProps extends ButtonProps {
  isDeleted?: boolean;
  withEdit?: boolean;
}

// CrudButton allows updating, deleting, undeleting, or removing the pointed to
// item. This component always operates on an address, but it picks up its
// modifyFn and fetchFn from the useModifyFn hook. This allows different views
// (MonitorsView, AbisView, etc.) to use the button without having to know the
// details of how the view operates.
export const CrudButton = ({ value, isDeleted = false, withEdit = false, ...props }: CrudButtonProps) => {
  if (isDeleted) {
    return (
      <Group justify="flex-end">
        <RemoveButton value={value} {...props} />
        <UndeleteButton value={value} {...props} />
      </Group>
    );
  }

  return (
    <Group justify="flex-end">
      {withEdit ? <EditButton value={value} {...props} /> : <></>}
      <DeleteButton value={value} {...props} />
    </Group>
  );
};
