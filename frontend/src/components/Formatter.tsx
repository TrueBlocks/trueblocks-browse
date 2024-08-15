import React, { ReactNode } from "react";
import { IconCircleCheck } from "@tabler/icons-react";
import { base } from "@gocode/models";
import { useDateTime } from "@hooks";
import { AddressFormatter } from "./AddressFormatter";

export type knownTypes =
  | "text"
  | "float"
  | "int"
  | "bytes"
  | "date"
  | "boolean"
  | "check"
  | "address"
  | "name"
  | "hash"
  | "error";

export const Formatter: React.FC<{ type: knownTypes; value: any; subType?: string }> = ({ type, value, subType }) => {
  const formatInteger = (number: number): string => {
    return new Intl.NumberFormat(navigator.language).format(number);
  };

  const formatFloat = (number: number): string => {
    return number?.toFixed(4);
  };

  const formatBytes = (bytes: number): string => {
    if (bytes === 0) return "0 Bytes";
    const k = 1024;
    const sizes = ["bytes", "Kb", "Mb", "Gb", "Tb", "Pb"];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    const formattedValue = (bytes / Math.pow(k, i)).toLocaleString("en-US", {
      minimumFractionDigits: 1,
      maximumFractionDigits: 1,
    });
    return `${formattedValue} ${sizes[i]}`;
  };

  var v = value as number;
  switch (type) {
    case "float":
      return <>{formatFloat(v)}</>;
    case "bytes":
      return <>{formatBytes(v)}</>;
    case "int":
      return <>{formatInteger(v)}</>;
    case "address":
      return <AddressFormatter address={value as base.Address} />;
    case "name":
      return <AddressFormatter showSame={false} address={value as base.Address} />;
    case "date":
      return <>{useDateTime(v)}</>;
    case "boolean":
      var fill = value ? "green" : "red";
      return <IconCircleCheck size={16} color="white" fill={fill} />;
    case "check":
      return <>{value ? <IconCircleCheck size={16} color="white" fill="green" /> : <></>}</>;
    case "error":
      return <>{value ? <IconCircleCheck size={16} color="white" fill="red" /> : <></>}</>;
    default:
      return <>{value}</>;
  }
};
