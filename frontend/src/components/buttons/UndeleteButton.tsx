import { Group } from "@mantine/core";
import { IconArrowBackUp } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { base } from "@gocode/models";
import { useModifyFn } from "@hooks";

export const UndeleteButton = ({ value, ...props }: ButtonProps) => {
  const address = value as base.Address;
  const { undeleteItem } = useModifyFn(address);
  return (
    <Group justify="flex-end">
      <BaseButton {...props} tip="Undelete" onClick={undeleteItem} icon={<IconArrowBackUp />} />
    </Group>
  );
};
