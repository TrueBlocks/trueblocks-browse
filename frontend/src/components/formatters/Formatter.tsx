import React from "react";
import { IconCircleCheck } from "@tabler/icons-react";
import { TextProps } from "@mantine/core";
import { base } from "@gocode/models";
import { useDateTime, useToEther } from "@hooks";
import { AddressFormatter, getDebugColor, Popup, NamePopup, AddressPopup } from ".";
import { useAppState } from "@state";
import classes from "./Formatter.module.css";
import { TextFormatter } from "./TextFormatter";

export type knownTypes =
  | "address-and-name"
  | "address-address-only"
  | "address-name-only"
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
  const isCurrent = value === address;
  const cn = getDebugColor(type) || (isCurrent ? classes.bold : className);


  switch (type) {
    case "boolean":
      return <IconCircleCheck size={16} color="white" fill={value ? "green" : "red"} />;
    case "check":
      return value ? <IconCircleCheck size={16} color="white" fill="green" /> : <></>;
    case "error":
      return value ? <IconCircleCheck size={16} color="white" fill="red" /> : <></>;
    case "address-and-name":
      return <AddressFormatter className={cn} size={size} addressIn={value as base.Address} />;
    case "address-address-only":
      return (
        <Popup editor={<AddressPopup address={value} />}>
          <TextFormatter value={value} size={size} type={type} className={cn} />
        </Popup>
      );
    case "address-name-only":
      return (
        <Popup editor={<NamePopup name={value} onSubmit={(newValue: string) => console.log(newValue)} />}>
          <TextFormatter value={value} size={size} type={type} className={cn} />
        </Popup>
      );
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

  return <TextFormatter value={value} size={size} type={type} className={cn} />
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
