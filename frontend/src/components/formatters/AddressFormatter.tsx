import React, { useEffect, useState } from "react";
import { AddrToName } from "@gocode/app/App";
import { base } from "@gocode/models";
import { Formatter } from ".";
import { useAppState } from "@state";

export const AddressFormatter = ({ addressIn }: { addressIn: base.Address }) => {
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
    return <Formatter type="text" value={formattedAddress} />;
  }

  return <Formatter type="text" value={formattedAddress} />;
};
