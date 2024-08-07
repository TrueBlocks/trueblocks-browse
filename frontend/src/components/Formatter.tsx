import React, { ReactNode, useEffect, useState } from "react";
import { IconCircleCheck } from "@tabler/icons-react";
import { AddrToName } from "@gocode/app/App";
import { base } from "@gocode/models";
import { useDateTime } from "@hooks";

export type knownTypes = "text" | "float" | "int" | "bytes" | "date" | "boolean" | "check" | "address" | "hash";

export const Formatter: React.FC<{ type: knownTypes; value: any }> = ({ type, value }) => {
  const formatInteger = (number: number): ReactNode => {
    const n = new Intl.NumberFormat(navigator.language).format(number);
    return <>{n}</>;
  };

  const formatFloat = (number: number): ReactNode => {
    const n = number?.toFixed(4);
    return <>{n}</>;
  };

  const formatBytes = (bytes: number): ReactNode => {
    if (bytes === 0) return <>0 Bytes</>;
    const k = 1024;
    const sizes = ["bytes", "Kb", "Mb", "Gb", "Tb", "Pb"];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    const formattedValue = (bytes / Math.pow(k, i)).toLocaleString("en-US", {
      minimumFractionDigits: 1,
      maximumFractionDigits: 1,
    });
    return <>{`${formattedValue} ${sizes[i]}`}</>;
  };

  var v = value as number;
  switch (type) {
    case "float":
      return formatFloat(v);
    case "bytes":
      return formatBytes(v);
    case "int":
      return formatInteger(v);
    case "address":
      return <FormatAddressComponent address={value as base.Address} />;
    case "date":
      return useDateTime(v);
    case "boolean":
      var fill = value ? "green" : "red";
      return <IconCircleCheck size={20} color="white" fill={fill} />;
    case "check":
      return <>{value ? <IconCircleCheck size={20} color="white" fill="green" /> : <></>}</>;
    default:
      return value;
  }
};

const FormatAddressComponent = ({ address }: { address: base.Address }) => {
  const [formattedAddress, setFormattedAddress] = useState<string>("");
  useEffect(() => {
    const formatAddress = async () => {
      const name = await AddrToName(address);
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
