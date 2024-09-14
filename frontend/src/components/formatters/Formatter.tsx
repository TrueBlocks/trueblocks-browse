import { TextProps } from "@mantine/core";
import { IconCircleCheck } from "@tabler/icons-react";
import {
  AddressFormatter,
  AppearanceFormatter,
  CrudButton,
  DateFormatter,
  TagFormatter,
  TextFormatter,
  EdMode,
} from "@components";
import { base } from "@gocode/models";
import { useAppState } from "@state";
import { GetDebugColor } from ".";

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
  | "crud"
  | "date"
  | "error"
  | "ether"
  | "gas"
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
  value2?: boolean | base.Hash | base.Address | string | undefined;
  className?: string;
  size?: TextProps["size"];
};

export const Formatter = ({ type, value, value2, className, size = "md" }: FormatterProps) => {
  const { address } = useAppState();

  const cn = GetDebugColor(type) || className;
  const n = value as number;
  const bi = value as bigint;
  const bool = value2 as boolean;
  const from = value2 as unknown as base.Address;

  switch (type) {
    case "boolean":
      return <IconCircleCheck size={16} color="white" fill={value ? "green" : "red"} />;
    case "check":
      return value ? <IconCircleCheck size={16} color="white" fill="green" /> : <></>;
    case "error":
      return value ? <IconCircleCheck size={16} color="white" fill="red" /> : <></>;
    case "tag":
      return <TagFormatter value={value} size={size} className={cn} />;
    case "gas":
      value = from !== address ? "-" : formatEther(bi);
      break;
    case "ether":
      value = bool ? "" : formatEther(bi);
      break;
    case "timestamp":
      value = formatDateTime(n);
      break;
    case "date":
      value = value?.replace("T", " ");
      if ((value?.match(/ /g)?.length ?? 0) > 0) {
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
      return <AppearanceFormatter value={value} value2={value2} size={size} className={cn} />;
    case "hash":
    case "path":
    case "range":
    case "text":
    case "url":
      break;
    case "crud":
      return <CrudButton size="xs" value={value} isDeleted={bool} />;
    case "address-editor":
      return (
        <AddressFormatter type={type} className={cn} size={size} value={value} value2={value2} mode={EdMode.All} />
      );
    case "address-address-only":
      return (
        <AddressFormatter type={type} className={cn} size={size} value={value} value2={value2} mode={EdMode.Address} />
      );
    case "address-name-only":
      return (
        <AddressFormatter type={type} className={"cn"} size={size} value={value} value2={value2} mode={EdMode.Name} />
      );
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

const formatEther = (value: bigint | string) => {
  // from https://viem.sh/docs/utilities/formatUnits
  if (typeof value === "string" && value.includes(".")) {
    return value;
  }
  if (!value) return "-";

  let display = value.toString();
  const negative = display.startsWith("-");
  if (negative) display = display.slice(1);
  display = display.padStart(18, "0");

  const integer = display.slice(0, display.length - 18);
  let fraction = display.slice(display.length - 18);
  fraction = fraction.slice(0, 5).padEnd(5, "0");

  const v = `${negative ? "-" : ""}${integer || "0"}.${fraction}`;
  // return display === "000000000000000000" ? "-" : v + " --- " + display;
  if (v === "0.00000") return "-";
  return v;
};

const formatDateTime = (timestamp: number): string => {
  const date = new Date(timestamp * 1000); // Convert timestamp from seconds to milliseconds

  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, "0");
  const day = String(date.getDate()).padStart(2, "0");
  const hours = String(date.getHours()).padStart(2, "0");
  const minutes = String(date.getMinutes()).padStart(2, "0");
  const seconds = String(date.getSeconds()).padStart(2, "0");

  const formatted = `${year}-${month}-${day}T${hours}:${minutes}:${seconds}`;
  return formatted;
};
