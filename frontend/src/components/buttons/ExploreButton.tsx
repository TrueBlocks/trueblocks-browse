import { useState, useEffect } from "react";
import { Button } from "@mantine/core";
import { IconExternalLink } from "@tabler/icons-react";
import { ButtonProps } from "@components";
import { GetChainInfo } from "@gocode/app/App";
import { types } from "@gocode/models";
import { BrowserOpenURL } from "@runtime";
import { useAppState } from "@state";

export interface ExploreButtonProps extends ButtonProps {
  endpoint: string;
}

export const ExploreButton = ({ endpoint, value, noText, size, onClick }: ExploreButtonProps) => {
  const { chain } = useAppState();
  const [chainInfo, setChainInfo] = useState<types.Chain>({} as types.Chain);

  useEffect(() => {
    GetChainInfo(chain).then((info) => {
      setChainInfo(info);
    });
  }, [chain]);

  const handleClick = () => {
    BrowserOpenURL(`${chainInfo.remoteExplorer}/${endpoint}/${value}`);
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
