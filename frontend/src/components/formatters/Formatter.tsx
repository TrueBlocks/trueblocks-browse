import React from "react";
import { IconCircleCheck } from "@tabler/icons-react";
import { TextProps } from "@mantine/core";
import { useDateTime, useToEther } from "@hooks";
import { useAppState } from "@state";
import classes from "./Formatter.module.css";
import { getDebugColor } from ".";
import { AddressEditor, DateFormatter, TagFormatter, TextFormatter, EdMode } from "@components";

export type knownType =
  | "address-editor"
  | "address-address-only"
  | "address-name-only"
  | "address-line1"
  | "address-line2"
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
  | "path"
  | "range"
  | "tag"
  | "text"
  | "time"
  | "timestamp"
  | "url";

export type FormatterProps = {
  type: knownType;
  value: any;
  value2?: any;
  className?: string;
  size?: TextProps["size"];
};

export const Formatter = ({ type, value, value2, className, size = "md" }: FormatterProps) => {
  const { address } = useAppState();

  var n = value as number;
  var bi = value as bigint;
  const isCurrent = value === address;
  const cn = getDebugColor(type) || (isCurrent ? classes.bold : className);

  switch (type) {
    case "boolean":
      return <IconCircleCheck size={16} color="white" fill={value ? "green" : "red"} />;
    case "check":
      return value ? <IconCircleCheck size={16} color="white" fill="green" /> : <></>;
    case "error":
      return value ? <IconCircleCheck size={16} color="white" fill="red" /> : <></>;
    case "tag":
      return <TagFormatter value={value} size={size} className={cn} />;
    case "ether":
      value = useToEther(bi);
      break;
    case "timestamp":
      value = useDateTime(n);
    case "date":
      value = value?.replace("T", " ");
      if ((value?.match(/ /g)?.length ?? 0) > 0) {
        // at least one space
        return <DateFormatter value={value} size={size} className={cn} />;
      }
      // else, render using TextFormatter
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
    case "hash":
    case "path":
    case "range":
    case "text":
    case "url":
      break;
    case "address-editor":
      return <AddressEditor type={type} className={cn} size={size} value={value} value2={value2} mode={EdMode.All} />;
    case "address-address-only":
      return (
        <AddressEditor type={type} className={cn} size={size} value={value} value2={value2} mode={EdMode.Address} />
      );
    case "address-name-only":
      return <AddressEditor type={type} className={cn} size={size} value={value} value2={value2} mode={EdMode.Name} />;
    case "address-line1":
      return <TextFormatter value={value} size={size} type={type} className={cn} />;
    case "address-line2":
      return <TextFormatter value={value} size="xs" type={type} className={cn} />;
    default:
      value = "UNKNOWN FORMATTER TYPE";
  }

  return <TextFormatter value={value} size={size} type={type} className={cn} />;
};

const formatInteger = (number: number): string => {
  return number === 0 ? "-" : new Intl.NumberFormat(navigator.language).format(number);
};

const formatFloat = (number: number): string => {
  return number?.toFixed(4);
};

const formatBytes = (bytes: number): string => {
  if (bytes === 0) return "-";
  const k = 1024;
  const sizes = ["b", "Kb", "Mb", "Gb", "Tb", "Pb"];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  const formattedValue = (bytes / Math.pow(k, i)).toLocaleString("en-US", {
    minimumFractionDigits: 1,
    maximumFractionDigits: 1,
  });
  return `${formattedValue} ${sizes[i]}`;
};
