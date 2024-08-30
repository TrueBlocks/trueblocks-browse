import React from "react";

export type ButtonProps = {
  value: string;
  noText?: boolean;
  size: "xs" | "sm" | "md" | "lg" | "xl";
  onClick?: () => void;
};
