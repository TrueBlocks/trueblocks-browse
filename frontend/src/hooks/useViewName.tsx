import { useLocation } from "wouter";

export function useViewName(): string {
  const [location] = useLocation();

  const baseRoute = location.split("/")[1] || "";
  const viewName =
    baseRoute === "" ? "Portfolio View" : `${baseRoute.charAt(0).toUpperCase()}${baseRoute.slice(1)} View`;

  return viewName;
}
