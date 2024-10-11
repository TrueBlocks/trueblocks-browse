import { useState, useEffect } from "react";
import { Image } from "@mantine/core";
import { GetExploreUrl } from "@gocode/app/App";
import { BrowserOpenURL } from "@runtime";
import { FormatterProps } from "./Formatter";

export interface DalleImageProps extends Omit<FormatterProps, "type"> {
  height?: number;
}

export function DalleImage({ value, height = 125 }: DalleImageProps) {
  const [url, setUrl] = useState(value);

  useEffect(() => {
    GetExploreUrl(value as string, false, true).then((url) => {
      setUrl(url);
    });
  }, [value]);

  // Handler to open the URL using Wails' browser
  const handleImageClick = () => {
    if (url) {
      BrowserOpenURL(url);
    }
  };

  const s =
    height === 125
      ? { cursor: "pointer", marginTop: -10, marginBottom: -10 }
      : { cursor: "pointer", marginTop: 0, marginBottom: 0 };
  return <Image onClick={handleImageClick} style={s} src={url} alt={url} height={height} fit={"contain"} />;
}
