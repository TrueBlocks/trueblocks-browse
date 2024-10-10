import { useState, useEffect } from "react";
import { IconBrandOpenai } from "@tabler/icons-react";
import { BaseButton, ButtonProps } from "@components";
import { GetExploreUrl } from "@gocode/app/App";
import { BrowserOpenURL } from "@runtime";

export const DalleButton = ({ value, ...props }: ButtonProps) => {
  const [url, setUrl] = useState("");
  const icon = <IconBrandOpenai />;

  useEffect(() => {
    GetExploreUrl(value as string, false, true).then((url) => {
      setUrl(url);
    });
  }, [value]);

  const handleClick = () => {
    BrowserOpenURL(url);
  };

  return <BaseButton tip="Dalle" onClick={handleClick} leftSection={icon} {...props} />;
};
