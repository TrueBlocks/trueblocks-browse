import { base } from "@gocode/models";

export type ButtonProps = {
  value: string | base.Address;
  size: "xs" | "sm" | "md" | "lg" | "xl";
  noText?: boolean;
  onClick?: () => void;
};
