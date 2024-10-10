import { IconLink } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { SetSessionVal } from "@gocode/app/App";
import { messages } from "@gocode/models";
import { EventsEmit } from "@runtime";

export const ViewButton = ({ value, ...props }: ButtonProps) => {
  const icon = <IconLink />;

  const handleClick = () => {
    SetSessionVal("route", `/history/${value}`);
    EventsEmit(messages.Message.NAVIGATE, {
      route: `/history/${value}`,
    });
  };

  return <BaseButton tip="View" onClick={handleClick} leftSection={icon} {...props} />;
};
