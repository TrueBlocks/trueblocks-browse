import { Button, ActionIcon } from "@mantine/core";
import { IconLink } from "@tabler/icons-react";
import { ButtonProps } from "@components";
import { SetSessionVal } from "@gocode/app/App";
import { messages } from "@gocode/models";
import { EventsEmit } from "@runtime";

export const ViewButton = ({ value, noText, onClick }: Omit<ButtonProps, "size">) => {
  const handleClick = () => {
    SetSessionVal("route", `/history/${value}`);
    EventsEmit(messages.Message.NAVIGATE, {
      route: `/history/${value}`,
    });
    if (onClick) {
      onClick();
    }
  };

  const size = "sm";
  if (noText) {
    return (
      <ActionIcon size={size} onClick={handleClick} title="View">
        <IconLink />
      </ActionIcon>
    );
  }

  return (
    <Button size={size} onClick={handleClick} leftSection={<IconLink />}>
      {"View"}
    </Button>
  );
};
