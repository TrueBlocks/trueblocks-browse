import { base } from "@gocode/models";
import { useAppState } from "../state";

// from https://viem.sh/docs/utilities/formatUnits
export function useToEther(value: bigint | string) {
  // Check if the input is a string that already contains a decimal
  if (typeof value === "string" && value.includes(".")) {
    return value;
  }

  // Handle bigint input as usual
  if (!value) return "-";

  let display = value.toString();
  const negative = display.startsWith("-");
  if (negative) display = display.slice(1);
  display = display.padStart(18, "0");

  const [integer] = [display.slice(0, display.length - 18), display.slice(display.length - 18)];
  let [fraction] = [display.slice(0, display.length - 18), display.slice(display.length - 18)];

  // Ensure the fraction has exactly 5 digits
  fraction = fraction.slice(0, 5).padEnd(5, "0");

  const v = `${negative ? "-" : ""}${integer || "0"}.${fraction}`;
  if (v === "0.00000") return "-";
  return v;
}

export function useToGas(value: bigint, from: base.Address) {
  const { address } = useAppState();
  const gas = useToEther(value);
  return from !== address ? "-" : gas;
}
