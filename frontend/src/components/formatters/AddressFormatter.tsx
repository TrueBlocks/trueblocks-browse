import React, { useEffect, useState } from "react";
import { TextProps } from "@mantine/core";
import { AddrToName } from "@gocode/app/App";
import { base } from "@gocode/models";
import { Formatter, FormatterProps, knownType } from ".";
import { useAppState } from "@state";

export const AddressFormatter = ({ value, className, size = "md" }: FormatterProps) => {
  const { address } = useAppState();
  const [formattedAddress, setFormattedAddress] = useState<string>("");
  const [type, setType] = useState<knownType>("text");

  useEffect(() => {
    const formatAddress = async () => {
      AddrToName(value).then((name) => {
        setFormattedAddress(name?.length > 0 ? name : value);
        setType(name?.length > 0 ? "address-name-only" : "address-address-only");
      });
    };
    formatAddress();
  }, [value]);

  if (value === address) {
    return <Formatter className={className} size={size} type={type} value={formattedAddress} />;
  }

  return <Formatter className={className} size={size} type={type} value={formattedAddress} />;
};
