import React from "react";
import { IconCircleCheck } from "@tabler/icons-react";
import { Text, TextProps } from "@mantine/core";
import { base } from "@gocode/models";
import { useDateTime, useToEther } from "@hooks";
import { AddressFormatter } from "./AddressFormatter";
import classes from "./Formatter.module.css";

export type knownTypes =
  | "address-name"
  | "address-only"
  | "appearance"
  | "boolean"
  | "bytes"
  | "check"
  | "date"
  | "error"
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
      return <MyText size={size}>{value}</MyText>;
    case "appearance":
      return <MyText size={size}>{value}</MyText>;
    case "boolean":
      var fill = value ? "green" : "red";
      return <IconCircleCheck size={16} color="white" fill={fill} />;
    case "bytes":
      return <MyText size={size}>{formatBytes(v)}</MyText>;
    case "check":
      return value ? <IconCircleCheck size={16} color="white" fill="green" /> : <></>;
    case "date":
      return <MyText size={size}>{useDateTime(v)}</MyText>;
    case "error":
      return <MyText size={size}>{value ? <IconCircleCheck size={16} color="white" fill="red" /> : <></>}</MyText>;
  case "ether":
      return <MyText size={size}>{useToEther(value as bigint)}</MyText>;
    case "float":
      return <MyText size={size}>{formatFloat(v)}</MyText>;
    case "hash":
      return <MyText size={size}>{value}</MyText>;
    case "int":
      if (v === 0) {
        return <MyText size={size}>{"-"}</MyText>;
      } else {
        return <MyText size={size}>{formatInteger(v)}</MyText>;
      }
    case "name-only":
      return <Formatter type="text" value={value} />;
    case "path":
      return <MyText size={size}>{value}</MyText>;
    case "range":
      return <MyText size={size}>{value}</MyText>;
    case "text":
      return <MyText size={size}>{value}</MyText>;
    case "timestamp":
      return <MyText size={size}>{useDateTime(v)}</MyText>;
    case "url":
      return <MyText size={size}>{value}</MyText>;
    default:
      return <MyText size={size}>UNKNOWN FORMATTER TYPE</MyText>;
  }
};

export const MyText = ({ size, children }: { size?: TextProps["size"]; children: React.ReactNode }) => {
  var color = "red";
  return (
    <Text className={classes.black} c={color} size={size}>
      {children}
    </Text>
  );
};
