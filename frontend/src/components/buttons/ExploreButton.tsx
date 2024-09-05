import { Button } from "@mantine/core";
import { IconExternalLink } from "@tabler/icons-react";
import { ButtonProps } from "@components";
import { BrowserOpenURL } from "@runtime";

export interface ExploreButtonProps extends ButtonProps {
  endpoint: string;
}

export const ExploreButton = ({ endpoint, value, noText, size, onClick }: ExploreButtonProps) => {
  const handleClick = () => {
    BrowserOpenURL(`https://etherscan.io/${endpoint}/${value}`);
    if (onClick) {
      onClick();
    }
  };

  return (
    <Button size={size} onClick={handleClick} leftSection={<IconExternalLink />}>
      {noText ? null : "Explore"}
    </Button>
  );
};
