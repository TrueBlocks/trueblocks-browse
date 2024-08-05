import React from "react";
import { Title } from "@mantine/core";
import { useLocation } from "wouter";

export function ViewTitle(): JSX.Element {
  const [location] = useLocation();
  const baseRoute = location.split("/")[1] || "";
  const viewName = baseRoute === "" ? "Home View" : `${baseRoute.charAt(0).toUpperCase()}${baseRoute.slice(1)} View`;
  return <Title order={3}>{viewName}</Title>;
}
