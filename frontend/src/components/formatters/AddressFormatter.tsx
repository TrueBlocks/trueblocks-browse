import React, { useEffect, useState } from "react";
import { AddrToName } from "@gocode/app/App";
import { base } from "@gocode/models";
import { Formatter } from ".";

export const AddressFormatter = ({ address }: { address: base.Address }) => {
  const [formattedAddress, setFormattedAddress] = useState<string>("");
  useEffect(() => {
    const formatAddress = async () => {
      const name = await AddrToName(address);
      const isHex = /^0x[0-9A-Fa-fx]+$/.test(name);
      if (name && name.length > 0) {
        setFormattedAddress(name);
      } else {
        setFormattedAddress(address as unknown as string);
      }
    };
    formatAddress();
  }, [address]);

  return <Formatter type="text" value={formattedAddress} />;
};
