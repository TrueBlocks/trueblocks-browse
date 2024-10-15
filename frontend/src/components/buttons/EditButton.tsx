import { Group } from "@mantine/core";
import { IconEdit } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { base } from "@gocode/models";
import { useModifyFn } from "@hooks";

export const EditButton = ({ value, ...props }: ButtonProps) => {
  const address = value as base.Address;
  const { modifyItem } = useModifyFn(address);
  return (
    <Group justify="flex-end">
      <BaseButton {...props} tip="Edit" onClick={modifyItem} leftSection={<IconEdit />} />
    </Group>
  );
};
