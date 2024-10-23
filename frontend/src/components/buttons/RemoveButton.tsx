import { Group } from "@mantine/core";
import { IconTrashX } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { base } from "@gocode/models";
import { useModifyFn } from "@hooks";

export const RemoveButton = ({ value, ...props }: ButtonProps) => {
  const address = value as base.Address;
  const { removeItem } = useModifyFn(address);
  return (
    <Group justify="flex-end">
      <BaseButton {...props} tip="Remove" onClick={removeItem} icon={<IconTrashX />} />
    </Group>
  );
};
