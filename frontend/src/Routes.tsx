import React from "react";

// Find: NewViews-React
import {
  // Icons
  HomeIcon,
  HistoryIcon,
  MonitorIcon,
  NamesIcon,
  IndexesIcon,
  ManifestIcon,
  AbisIcon,
  StatusIcon,
  DaemonsIcon,
  SettingsIcon,
  // Views
  HomeView,
  HistoryView,
  NamesView,
  MonitorsView,
  IndexesView,
  ManifestView,
  AbisView,
  DaemonsView,
  StatusView,
  SettingsView,
} from "@views";

export type Route =
  | ""
  | "history"
  | "monitors"
  | "names"
  | "abis"
  | "indexes"
  | "manifest"
  | "status"
  | "settings"
  | "daemons";

export type FullRoute = string;

function asFull(r: Route): FullRoute {
  if (r === "history") {
    return "/" + r + "/:address";
  }

  return "/" + r;
}

type RouteItem = {
  order: number;
  route: FullRoute;
  label: string;
  icon: JSX.Element;
  component: React.ComponentType;
};

export const routeItems: RouteItem[] = [
  {
    order: 10,
    route: asFull("history"),
    label: "History",
    icon: HistoryIcon,
    component: HistoryView,
  },
  {
    order: 20,
    route: asFull("monitors"),
    label: "Monitors",
    icon: MonitorIcon,
    component: MonitorsView,
  },
  {
    order: 30,
    route: asFull("names"),
    label: "Names",
    icon: NamesIcon,
    component: NamesView,
  },
  {
    order: 40,
    route: asFull("abis"),
    label: "Abis",
    icon: AbisIcon,
    component: AbisView,
  },
  {
    order: 50,
    route: asFull("indexes"),
    label: "Indexes",
    icon: IndexesIcon,
    component: IndexesView,
  },
  {
    order: 60,
    route: asFull("manifest"),
    label: "Manifest",
    icon: ManifestIcon,
    component: ManifestView,
  },
  {
    order: 70,
    route: asFull("status"),
    label: "Status",
    icon: StatusIcon,
    component: StatusView,
  },
  {
    order: 80,
    route: asFull("settings"),
    label: "Settings",
    icon: SettingsIcon,
    component: SettingsView,
  },
  {
    order: 90,
    route: asFull("daemons"),
    label: "Daemons",
    icon: DaemonsIcon,
    component: DaemonsView,
  },
  {
    order: 0,
    route: asFull(""),
    label: "Home",
    icon: HomeIcon,
    component: HomeView,
  },
];
