import React from "react";
import { IconCircleCheck } from "@tabler/icons-react";
import { Text, TextProps } from "@mantine/core";
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
  | "hash"
  | "error";

export const Formatter = ({ type, value, size = "md" }: { type: knownTypes; value: any; size?: TextProps["size"] }) => {
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
      return <Text size={size}>{formatFloat(v)}</Text>;
    case "bytes":
      return <Text size={size}>{formatBytes(v)}</Text>;
    case "int":
      return <Text size={size}>{formatInteger(v)}</Text>;
    case "address":
      return <AddressFormatter addressIn={value as base.Address} />;
    case "date":
      return <Text size={size}>{useDateTime(v)}</Text>;
    case "boolean":
      var fill = value ? "green" : "red";
      return <IconCircleCheck size={16} color="white" fill={fill} />;
    case "check":
      return value ? <IconCircleCheck size={16} color="white" fill="green" /> : <></>;
    case "error":
      return <Text size={size}>{value ? <IconCircleCheck size={16} color="white" fill="red" /> : <></>}</Text>;
    default:
      return <Text size={size}>{value}</Text>;
  }
};
