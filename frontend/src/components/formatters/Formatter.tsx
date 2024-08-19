import React from "react";
import { IconCircleCheck } from "@tabler/icons-react";
import { Text, TextProps } from "@mantine/core";
import { base } from "@gocode/models";
import { useDateTime, useToEther } from "@hooks";
import { AddressFormatter } from "./AddressFormatter";
import { useAppState } from "@state";
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
  className?: string;
};

export const Formatter = ({ type, size = "md", className, value, value2 = null }: FormatterProps) => {
  const { address } = useAppState();

  var n = value as number;
  var bi = value as bigint;
  var addr = value as unknown as base.Address;

  switch (type) {
    case "address-name":
      return <AddressFormatter size={size} addressIn={value as base.Address} />;
    case "address-only":
    case "name-only":
      className = addr === address ? classes.bold : className;
      break;
    case "boolean":
      return <IconCircleCheck size={16} color="white" fill={value ? "green" : "red"} />;
    case "check":
      return value ? <IconCircleCheck size={16} color="white" fill="green" /> : <></>;
    case "error":
      return value ? <IconCircleCheck size={16} color="white" fill="red" /> : <></>;
    case "ether":
      value = useToEther(bi);
      break;
    case "timestamp":
      value = useDateTime(n);
      break;
    case "bytes":
      value = formatBytes(n);
      break;
    case "float":
      value = formatFloat(n);
      break;
    case "int":
      value = formatInteger(n);
      break;
    case "appearance":
    case "date":
    case "hash":
    case "path":
    case "range":
    case "text":
    case "url":
      break;
    default:
      value = "UNKNOWN FORMATTER TYPE";
  }

  return (
    <Text className={getDebugColor(type) || className} size={size}>
      {value}
    </Text>
  );
};

const getDebugColor = (type: knownTypes): string | null => {
  var ret: string | null = null;
  switch (type) {
    case "address-name":
      break;
    case "boolean":
      break;
    case "check":
      break;
    case "error":
      break;
    case "ether":
      ret = classes.brown;
      break;
    case "timestamp":
      ret = classes.blue;
      break;
    case "bytes":
      ret = classes.green;
      break;
    case "float":
      ret = classes.red;
      break;
    case "int":
      ret = classes.pink;
      break;
    case "address-only":
      ret = classes.orange;
      break;
    case "appearance":
      ret = classes.lightblue;
      break;
    case "date":
      ret = classes.purple;
      break;
    case "hash":
      ret = classes.lightblue;
      break;
    case "name-only":
      ret = classes.orange;
      break;
    case "path":
      ret = classes.pink;
      break;
    case "range":
      ret = classes.green;
      break;
    case "text":
      ret = classes.green;
      break;
    case "url":
      ret = classes.red;
      break;
    default:
      ret = classes.blue;
      break;
  }
  return null; // ret;
};

const formatInteger = (number: number): string => {
  return number === 0 ? "-" : new Intl.NumberFormat(navigator.language).format(number);
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
