export type ButtonProps = {
  value: string;
  noText?: boolean;
  size: "xs" | "sm" | "md" | "lg" | "xl";
  onClick?: () => void;
};
