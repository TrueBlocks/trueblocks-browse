import { useState, useEffect } from "react";
import { IconExternalLink, IconBrandOpenai, IconBrandGoogle } from "@tabler/icons-react";
import { ButtonProps, BaseButton } from "@components";
import { GetExploreUrl } from "@gocode/app/App";
import { BrowserOpenURL } from "@runtime";

export interface ExploreButtonProps extends ButtonProps {
  type?: "explore" | "google" | "dalle" | "url";
}

// ExploreButton opens a browser window to an explorer. It's also the basis
// for the DalleButton and GoogleButton components.
export const ExploreButton = ({ type = "explore", value, ...props }: ExploreButtonProps) => {
  const [url, setUrl] = useState("");
  const [icon, setIcon] = useState(<IconExternalLink />);
  const [tip, setTip] = useState("Explore");

  useEffect(() => {
    if (type == "url") {
      setTip("Spec");
      setIcon(<IconExternalLink />);
      setUrl(value as string);
    } else {
      const google = type === "google";
      const dalle = type === "dalle";
      GetExploreUrl(value as string, google, dalle).then((url) => {
        setUrl(url);
      });
      switch (type) {
        case "explore":
          setTip("Explore");
          setIcon(<IconExternalLink />);
          break;
        case "google":
          setTip("Google");
          setIcon(<IconBrandGoogle />);
          break;
        case "dalle":
          setTip("Dalle");
          setIcon(<IconBrandOpenai />);
          break;
      }
    }
  }, [value, type]);

  const handleClick = () => {
    BrowserOpenURL(url);
  };

  return <BaseButton {...props} tip={tip} onClick={handleClick} leftSection={icon} />;
};
