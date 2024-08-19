import React from "react";
import { IconCircleCheck } from "@tabler/icons-react";
import { Text, TextProps } from "@mantine/core";
import { base } from "@gocode/models";
import { useDateTime, useToEther } from "@hooks";
import { AddressFormatter } from "./AddressFormatter";

export type knownTypes =
  | "address-name"
  | "address-only"
  | "appearance"
  | "boolean"
  | "bytes"
  | "check"
  | "date"
  | "error"
  | "ether"
  | "float"
  | "hash"
  | "int"
  | "name-only"
  | "path"
  | "range"
  | "text"
  | "timestamp"
  | "url";

type FormatterProps = {
  type: knownTypes;
  size?: TextProps["size"];
  value: any;
  value2?: any;
};

export const Formatter = ({ type, size = "md", value, value2 = null }: FormatterProps) => {
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
    case "address-name":
      return <AddressFormatter size={size} addressIn={value as base.Address} />;
    case "address-only":
      return <Text size={size}>{value}</Text>;
    case "appearance":
      return <Text size={size}>{value}</Text>;
    case "boolean":
      var fill = value ? "green" : "red";
      return <IconCircleCheck size={16} color="white" fill={fill} />;
    case "bytes":
      return <Text size={size}>{formatBytes(v)}</Text>;
    case "check":
      return value ? <IconCircleCheck size={16} color="white" fill="green" /> : <></>;
    case "date":
      return <Text size={size}>{value}</Text>;
    case "error":
      return <Text size={size}>{value ? <IconCircleCheck size={16} color="white" fill="red" /> : <></>}</Text>;
    case "ether":
      return <Text size={size}>{useToEther(value as bigint)}</Text>;
    case "float":
      return <Text size={size}>{formatFloat(v)}</Text>;
    case "hash":
      return <Text size={size}>{value}</Text>;
    case "int":
      if (v === 0) {
        return <Text size={size}>{"-"}</Text>;
      } else {
        return <Text size={size}>{formatInteger(v)}</Text>;
      }
    case "name-only":
      return <Formatter type="text" value={value} />;
    case "path":
      return <Text size={size}>{value}</Text>;
    case "range":
      return <Text size={size}>{value}</Text>;
    case "text":
      return <Text size={size}>{value}</Text>;
    case "timestamp":
      return <Text size={size}>{useDateTime(v)}</Text>;
    case "url":
      return <Text size={size}>{value}</Text>;
    default:
      return <Text size={size}>UNKNOWN FORMATTER TYPE</Text>;
  }
};
