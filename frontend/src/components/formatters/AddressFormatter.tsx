import React, { useEffect, useState } from "react";
import { TextProps } from "@mantine/core";
import { AddrToName } from "@gocode/app/App";
import { base } from "@gocode/models";
import { Formatter } from ".";
import { useAppState } from "@state";

export const AddressFormatter = ({ addressIn, size = "md" }: { addressIn: base.Address; size?: TextProps["size"] }) => {
  const { address } = useAppState();
  const [formattedAddress, setFormattedAddress] = useState<string>("");

  useEffect(() => {
    const formatAddress = async () => {
      const name = await AddrToName(addressIn);
      if (name && name.length > 0) {
        setFormattedAddress(name);
      } else {
        setFormattedAddress(addressIn as unknown as string);
      }
    };
    formatAddress();
  }, [addressIn]);

  if (addressIn === address) {
    return <Formatter size={size} type="text" value={formattedAddress} />;
  }

  return <Formatter size={size} type="text" value={formattedAddress} />;
};
