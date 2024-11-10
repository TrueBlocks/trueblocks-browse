import { Text } from "@mantine/core";
import { useAppState } from "@state";

export const Footer = () => {
  const { info, chain } = useAppState();
  const { status } = useAppState();

  const fn = () => {
    if (info.filename) {
      if (info.dirty) {
        return <i>{info.filename} (dirty)</i>;
      }
      return <>{info.filename}</>;
    }
    return <>{"no file loaded"}</>;
  };

  return (
    <Text size={"sm"}>
      {`${status.clientVersion} / ${chain} / not loaded / ${status.lastUpdate} file: `}
      {fn()}
    </Text>
  );
};
