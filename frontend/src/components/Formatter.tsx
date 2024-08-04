import React, { ReactNode } from "react";
import { IconCircleCheck } from "@tabler/icons-react";

export type knownTypes = "text" | "float" | "int" | "bytes" | "date" | "boolean" | "check" | "address" | "hash";

export const Formatter: React.FC<{ type: knownTypes; value: number }> = ({ type, value }) => {
  const formatInteger = (number: number): React.ReactNode => {
    const n = new Intl.NumberFormat(navigator.language).format(number);
    return <>{n}</>;
  };

  const formatFloat = (number: number): React.ReactNode => {
    const n = number.toFixed(4);
    return <>{n}</>;
  };

  const formatBytes = (bytes: number): React.ReactNode => {
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

  switch (type) {
    case "float":
      return formatFloat(value);
    case "bytes":
      return formatBytes(value);
    case "int":
      return formatInteger(value);
    case "boolean":
      var fill = value ? "green" : "red";
      return <IconCircleCheck size={20} color="white" fill={fill} />;
    case "check":
      return <>{value ? <IconCircleCheck size={20} color="white" fill="green" /> : <></>}</>;
    default:
      return value;
  }
};
