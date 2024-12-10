import { useMemo } from "react";
import { useLocation } from "wouter";
import { useAppState } from "@state";

export const useViewName = (): string => {
  const { route } = useAppState();
  return useMemo(() => {
    return `${route[0].toUpperCase()}${route.slice(1)} View`;
  }, [route, location]);
};
