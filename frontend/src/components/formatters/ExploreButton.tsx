import React, { forwardRef } from "react";
import { Button } from "@mantine/core";
import { BrowserOpenURL } from "@runtime";
import { IconExternalLink } from "@tabler/icons-react";
import { PopupProps } from ".";

export const ExploreButton = ({ address }: { address: string }) => {
  return (
    <Button
      size={"xs"}
      onClick={() => BrowserOpenURL(`https://etherscan.io/address/${address}`)}
      leftSection={<IconExternalLink />}
    >
      Explore
    </Button>
  );
};
