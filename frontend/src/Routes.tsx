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
  WizardIcon,
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
  WizardView,
} from "@views";

export type RouteItem = {
  order: number;
  route: string;
  label: string;
  icon: JSX.Element;
  component: React.ComponentType;
};

export const routeItems: RouteItem[] = [
  {
    order: 10,
    route: "/history/:address",
    label: "History",
    icon: HistoryIcon,
    component: HistoryView,
  },
  {
    order: 20,
    route: "/monitors",
    label: "Monitors",
    icon: MonitorIcon,
    component: MonitorsView,
  },
  {
    order: 30,
    route: "/names",
    label: "Names",
    icon: NamesIcon,
    component: NamesView,
  },
  {
    order: 40,
    route: "/abis",
    label: "Abis",
    icon: AbisIcon,
    component: AbisView,
  },
  {
    order: 50,
    route: "/indexes",
    label: "Indexes",
    icon: IndexesIcon,
    component: IndexesView,
  },
  {
    order: 60,
    route: "/manifest",
    label: "Manifest",
    icon: ManifestIcon,
    component: ManifestView,
  },
  {
    order: 70,
    route: "/status",
    label: "Status",
    icon: StatusIcon,
    component: StatusView,
  },
  {
    order: 80,
    route: "/settings",
    label: "Settings",
    icon: SettingsIcon,
    component: SettingsView,
  },
  {
    order: 90,
    route: "/daemons",
    label: "Daemons",
    icon: DaemonsIcon,
    component: DaemonsView,
  },
  {
    order: 100,
    route: "/wizard",
    label: "Wizard",
    icon: WizardIcon,
    component: WizardView,
  },
  {
    order: 0,
    route: "/",
    label: "Home",
    icon: HomeIcon,
    component: HomeView,
  },
];
