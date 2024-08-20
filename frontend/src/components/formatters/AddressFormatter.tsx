import React, { useEffect, useState } from "react";
import { TextProps } from "@mantine/core";
import { AddrToName } from "@gocode/app/App";
import { base } from "@gocode/models";
import { Formatter, knownTypes } from ".";
import { useAppState } from "@state";
import classes from "./Formatter.module.css";

export const AddressFormatter = ({
  addressIn,
  className,
  size = "md",
}: {
  addressIn: base.Address;
  className?: string;
  size?: TextProps["size"];
}) => {
  const { address } = useAppState();
  const [formattedAddress, setFormattedAddress] = useState<string>("");
  const [type, setType] = useState<knownTypes>("text");

  useEffect(() => {
    const formatAddress = async () => {
      AddrToName(addressIn).then((name) => {
        setFormattedAddress(name?.length > 0 ? name : (addressIn as unknown as string));
        setType(name?.length > 0 ? "address-name-only" : "address-address-only");
      });
    };
    formatAddress();
  }, [addressIn]);

  if (addressIn === address) {
    return <Formatter className={className} size={size} type={type} value={formattedAddress} />;
  }

  return <Formatter className={className} size={size} type={type} value={formattedAddress} />;
};
