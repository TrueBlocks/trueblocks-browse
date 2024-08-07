// from https://viem.sh/docs/utilities/formatUnits
export function useToEther(value: bigint) {
  let display = value.toString();
  const negative = display.startsWith("-");
  if (negative) display = display.slice(1);
  display = display.padStart(18, "0");

  let [integer, fraction] = [display.slice(0, display.length - 18), display.slice(display.length - 18)];

  // Ensure the fraction has exactly 5 digits
  fraction = fraction.slice(0, 5).padEnd(5, "0");

  let v = `${negative ? "-" : ""}${integer || "0"}.${fraction}`;
  if (v === "0.00000") return "-";
  return v;
}
