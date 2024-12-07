import { useLocation } from "wouter";
import { Route } from "@layout";

export function useViewName(): string {
  const [location] = useLocation();

  const baseRoute = location.split("/")[1] || "";
  const viewName = baseRoute === "" ? "Project View" : `${baseRoute.charAt(0).toUpperCase()}${baseRoute.slice(1)} View`;

  return viewName;
}

export function useViewRoute(): Route {
  const [location] = useLocation();
  const baseRoute = location.split("/")[1] || "";
  const route = (baseRoute === "" ? "project" : `${baseRoute}`).toLowerCase();
  return route as Route;
}
