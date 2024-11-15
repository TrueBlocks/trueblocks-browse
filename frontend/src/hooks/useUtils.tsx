import { useCallback } from "react";

export const useUtils = () => {
  const ShortenAddr = useCallback(
    (val: string) => (val.length > 14 ? `${val.slice(0, 8)}...${val.slice(-6)}` : val),
    []
  );

  return { ShortenAddr };
};
