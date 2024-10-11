import { IconLink } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { SetSessionVal, Reload } from "@gocode/app/App";
import { base, messages } from "@gocode/models";
import { EventsEmit } from "@runtime";

// ViewButton opens the history page for a given address.
export const ViewButton = ({ value, ...props }: ButtonProps) => {
  const givenAddress = value as base.Address;
  const icon = <IconLink />;

  const handleClick = () => {
    SetSessionVal("route", `/history/${value}`);
    Reload(givenAddress).then(() => {
      EventsEmit(messages.Message.NAVIGATE, {
        route: `/history/${value}`,
      });
    });
  };

  return <BaseButton {...props} tip="View" onClick={handleClick} leftSection={icon} />;
};
