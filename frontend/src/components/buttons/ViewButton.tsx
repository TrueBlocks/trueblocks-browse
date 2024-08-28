import React from "react";
import { Button } from "@mantine/core";
import { SetLast } from "@gocode/app/App";
import { EventsEmit } from "@runtime";
import { IconLink } from "@tabler/icons-react";
import { messages } from "@gocode/models";
import { ButtonProps } from "@components";

export interface ViewButtonProps extends ButtonProps {}

export const ViewButton = ({ value, noText, size, onClick }: ViewButtonProps) => {
  const handleClick = () => {
    SetLast("route", `/history/${value}`);
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
