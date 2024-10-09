import { useState, useEffect } from "react";
import { ActionIcon, Button } from "@mantine/core";
import { IconExternalLink, IconBrandGoogle, IconBrandOpenai } from "@tabler/icons-react";
import { ButtonProps } from "@components";
import { GetExploreUrl } from "@gocode/app/App";
import { BrowserOpenURL } from "@runtime";

export enum UrlType {
  Google = "google",
  Dalle = "dalle",
  Regular = "regular",
}

interface ExploreButtonProps extends Omit<ButtonProps, "size"> {
  urlType?: UrlType; // Add the UrlType prop
}

export const ExploreButton = ({ value, noText, urlType = UrlType.Regular, onClick }: ExploreButtonProps) => {
  const [url, setUrl] = useState("");
  const [icon, setIcon] = useState(<IconExternalLink />);
  const [text, setText] = useState("Explore");

  useEffect(() => {
    const isGoogle = urlType === UrlType.Google;
    const isDalle = urlType === UrlType.Dalle;
    if (isGoogle) {
      setIcon(<IconBrandGoogle />);
      setText("Google");
    } else if (isDalle) {
      setIcon(<IconBrandOpenai />);
      setText("Dalle");
    }
    GetExploreUrl(value as string, isGoogle, isDalle).then((url) => {
      url = url.replace("/simple/", "/five-tone-postal-protozoa/");
      url = url.replace("http://", "https://");
      setUrl(url);
    });
  }, [value, urlType]);

  const handleClick = () => {
    BrowserOpenURL(url);
    if (onClick) {
      onClick();
    }
  };

  const size = "sm";
  if (noText) {
    return (
      <ActionIcon size={size} onClick={handleClick} title={text}>
        {icon}
      </ActionIcon>
    );
  }

  return (
    <Button size={size} onClick={handleClick} leftSection={icon}>
      {text}
    </Button>
  );
};

export const DalleButton = ({ value, noText, onClick }: ExploreButtonProps) => {
  return <ExploreButton value={value} noText={noText} urlType={UrlType.Dalle} onClick={onClick} />;
};

export const GoogleButton = ({ value, noText, onClick }: ExploreButtonProps) => {
  return <ExploreButton value={value} noText={noText} urlType={UrlType.Google} onClick={onClick} />;
};
