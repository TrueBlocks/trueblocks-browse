import { Group } from "@mantine/core";
import { FieldGroup } from "@components";

type ButtonTrayProps = {
  buttonGroup: FieldGroup<any> | null;
};

export const ButtonTray = ({ buttonGroup }: ButtonTrayProps) => {
  if (!buttonGroup || !buttonGroup.buttons) {
    return null;
  }

  return (
    <Group justify="flex-end">
      {buttonGroup.buttons.map((button, bbIndex) => (
        <div key={bbIndex}>{button}</div>
      ))}
      <div>{/*spacer*/}</div>
    </Group>
  );
};
