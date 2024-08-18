import React from "react";
import { Button } from "@mantine/core";
import { SetLast } from "@gocode/app/App";
import { EventsEmit } from "@runtime";
import { IconLink } from "@tabler/icons-react";
import { messages } from "@gocode/models";

export const MonitorButton = ({ address }: { address: string }) => {
  return (
    <Button
      size={"xs"}
      onClick={(e) => {
        e.preventDefault();
        SetLast("route", `/history/${address}`);
        EventsEmit(messages.Message.NAVIGATE, {
          route: `/history/${address}`,
        });
      }}
      leftSection={<IconLink />}
    >
      View
    </Button>
  );
};
