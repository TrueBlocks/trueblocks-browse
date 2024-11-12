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
    <Group
      style={{
        marginLeft: "auto",
        display: "flex",
        gap: "4px",
      }}
    >
      {buttonGroup.buttons.map((b, index) => (
        <div key={index}>{b}</div>
      ))}
    </Group>
  );
};
