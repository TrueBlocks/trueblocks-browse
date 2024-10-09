import { useState, useEffect } from "react";
import { Image } from "@mantine/core";
import { GetExploreUrl } from "@gocode/app/App";
import { FormatterProps } from "./Formatter";

export interface DalleImageProps extends Omit<FormatterProps, "type"> {
  height?: number;
}

export function DalleImage({ value, height = 200 }: DalleImageProps) {
  const [url, setUrl] = useState(value);

  useEffect(() => {
    GetExploreUrl(value as string, false, true).then((url) => {
      url = url.replace("/simple/", "/five-tone-postal-protozoa/");
      url = url.replace("http://", "https://");
      setUrl(url);
    });
  }, [value]);

  return <Image src={url} alt={url} h={height} fit={"contain"} />;
}
