import { useState, useEffect } from "react";
import { Image } from "@mantine/core";
import { GetExploreUrl } from "@gocode/app/App";
import { BrowserOpenURL } from "@runtime";
import { FormatterProps } from "./Formatter";

export interface DalleImageProps extends Omit<FormatterProps, "type"> {
  height?: number;
}

export function DalleImage({ value, height = 175 }: DalleImageProps) {
  const [url, setUrl] = useState(value);

  useEffect(() => {
    GetExploreUrl(value as string, false, true).then((url) => {
      url = url.replace("/simple/", "/five-tone-postal-protozoa/");
      url = url.replace("http://", "https://");
      setUrl(url);
    });
  }, [value]);

  // Handler to open the URL using Wails' browser
  const handleImageClick = () => {
    if (url) {
      BrowserOpenURL(url);
    }
  };

  return (
    <Image
      onClick={handleImageClick}
      style={{ cursor: "pointer" }}
      src={url}
      alt={url}
      height={height}
      fit={"contain"}
    />
  );
}
