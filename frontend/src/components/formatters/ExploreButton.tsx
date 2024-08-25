import React from "react";
import { Button } from "@mantine/core";
import { BrowserOpenURL } from "@runtime";
import { IconExternalLink } from "@tabler/icons-react";

export type AddressButtonProps = {
  address: string;
  onClick?: () => void;
};

export const ExploreButton = ({ address, onClick }: AddressButtonProps) => {
  const handleClick = () => {
    BrowserOpenURL(`https://etherscan.io/address/${address}`);
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
