import { Button } from "@mantine/core";
import { IconLink } from "@tabler/icons-react";
import { ButtonProps } from "@components";
import { SetSessionVal } from "@gocode/app/App";
import { messages } from "@gocode/models";
import { EventsEmit } from "@runtime";

export const ViewButton = ({ value, noText, size, onClick }: ButtonProps) => {
  const handleClick = () => {
    SetSessionVal("route", `/history/${value}`);
    EventsEmit(messages.Message.NAVIGATE, {
      route: `/history/${value}`,
    });
    if (onClick) {
      onClick();
    }
  };

  return (
    <Button size={size} onClick={handleClick} leftSection={<IconLink />}>
      {noText ? null : "View"}
    </Button>
  );
};
