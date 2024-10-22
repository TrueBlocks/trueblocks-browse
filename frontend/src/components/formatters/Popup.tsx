import { ReactNode, useState } from "react";
import { Popover } from "@mantine/core";

export type PopupProps = {
  onClose: () => void;
  onSubmit: (value: string) => void;
};

export function Popup({ children, editor }: { children: ReactNode; editor: ReactNode }) {
  const [opened, setOpened] = useState(false);
  const toggleOpened = () => setOpened((o) => !o);

  return (
    <Popover opened={opened} onChange={setOpened} withArrow>
      <Popover.Target>
        <div onClick={toggleOpened}>{children}</div>
      </Popover.Target>
      {editor && <Popover.Dropdown bg="white">{editor}</Popover.Dropdown>}
    </Popover>
  );
}
