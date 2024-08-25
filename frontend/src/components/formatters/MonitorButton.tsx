import React from "react";
import { Button } from "@mantine/core";
import { SetLast } from "@gocode/app/App";
import { EventsEmit } from "@runtime";
import { IconLink } from "@tabler/icons-react";
import { messages } from "@gocode/models";
import { AddressButtonProps } from "."

export const MonitorButton = ({ address, onClick }: AddressButtonProps) => {
  const handleClick = () => {
    SetLast("route", `/history/${address}`);
    EventsEmit(messages.Message.NAVIGATE, {
      route: `/history/${address}`,
    });
    if (onClick) {
      onClick();
    }
  };

  return (
    <Button size={"xs"} onClick={handleClick} leftSection={<IconLink />}>
      View
    </Button>
  );
};
