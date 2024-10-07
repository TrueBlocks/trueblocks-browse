import { useState, useEffect } from "react";
import { ActionIcon, Button } from "@mantine/core";
import { IconExternalLink } from "@tabler/icons-react";
import { ButtonProps } from "@components";
import { GetChainInfo } from "@gocode/app/App";
import { types } from "@gocode/models";
import { BrowserOpenURL } from "@runtime";
import { useAppState } from "@state";

export interface ExploreButtonProps extends Omit<ButtonProps, "size"> {
  endpoint: string;
}

export const ExploreButton = ({ endpoint, value, noText, onClick }: ExploreButtonProps) => {
  const { chain } = useAppState();
  const [chainInfo, setChainInfo] = useState<types.Chain>({} as types.Chain);

  useEffect(() => {
    GetChainInfo(chain).then((info) => {
      setChainInfo(info);
    });
  }, [chain]);

  const handleClick = () => {
    const url = `${chainInfo.remoteExplorer}/${endpoint}/${value}`.replace(/\/\//g, "/");
    BrowserOpenURL(url);
    if (onClick) {
      onClick();
    }
  };

  const size = "sm";
  if (noText) {
    return (
      <ActionIcon size={size} onClick={handleClick} title="Explore">
        <IconExternalLink />
      </ActionIcon>
    );
  }

  return (
    <Button size={size} onClick={handleClick} leftSection={<IconExternalLink />}>
      {"Explore"}
    </Button>
  );
};
