import { useState, useEffect } from "react";
import { Image } from "@mantine/core";
import { GetExploreUrl, LoadDalleImage } from "@gocode/app/App";
import { base } from "@gocode/models";
import { BrowserOpenURL } from "@runtime";
import { FormatterProps } from "./Formatter";

export interface DalleImageProps extends Omit<FormatterProps, "type"> {
  height?: number;
}

export function DalleImage({ value, height = 125 }: DalleImageProps) {
  const [url, setUrl] = useState(value);
  const [localUrl, setLocalUrl] = useState<string | null>(null);

  useEffect(() => {
    LoadDalleImage(value as base.Address).then((found) => {
      if (found) {
        setLocalUrl(`/path/to/local/images/${value}`);
      } else {
        // Otherwise, set it to the remote URL
        GetExploreUrl(value as string, false, true).then((remoteUrl) => {
          setUrl(remoteUrl);
        });
      }
    });
  }, [value]);

  // Handler to open the URL using Wails' browser
  const handleImageClick = () => {
    if (localUrl || url) {
      BrowserOpenURL(localUrl || url);
    }
  };

  const s =
    height === 125
      ? { cursor: "pointer", marginTop: -10, marginBottom: -10 }
      : { cursor: "pointer", marginTop: 0, marginBottom: 0 };

  return (
    <Image
      onClick={handleImageClick}
      style={s}
      src={localUrl || url} // Display local image if available, else fallback to remote
      alt={localUrl || url}
      height={height}
      fit={"contain"}
    />
  );
}
