import React from "react";
import { Button } from "@mantine/core";
import { BrowserOpenURL } from "@runtime";
import { IconExternalLink } from "@tabler/icons-react";

export type ExploreButtonProps = {
  endpoint: string;
  value: string;
  onClick?: () => void;
};

export const ExploreButton = ({ endpoint, value, onClick }: ExploreButtonProps) => {
  const handleClick = () => {
    BrowserOpenURL(`https://etherscan.io/${endpoint}/${value}`);
    if (onClick) {
      onClick();
    }
  };

  return (
    <Button size={"xs"} onClick={handleClick} leftSection={<IconExternalLink />}>
      Explore
    </Button>
  );
};
