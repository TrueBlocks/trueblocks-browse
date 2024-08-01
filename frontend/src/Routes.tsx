import React from "react";

// Find: NewViews
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
  ServersIcon,
  SettingsIcon,
  // Views
  HomeView,
  HistoryView,
  NamesView,
  MonitorsView,
  IndexView,
  ManifestView,
  AbisView,
  ServersView,
  StatusView,
  SettingsView,
} from "@views";

// Note:
//  Change with care. The order of the items in this list matters (the last one is the default).
//  The order field is used to sort the menu items.
export const routeItems = [
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
    route: "/indexes",
    label: "Indexes",
    icon: IndexesIcon,
    component: IndexView,
  },
  {
    order: 50,
    route: "/manifest",
    label: "Manifest",
    icon: ManifestIcon,
    component: ManifestView,
  },
  {
    order: 60,
    route: "/abis",
    label: "Abis",
    icon: AbisIcon,
    component: AbisView,
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
    route: "/servers",
    label: "Servers",
    icon: ServersIcon,
    component: ServersView,
  },
  {
    order: 90,
    route: "/settings",
    label: "Settings",
    icon: SettingsIcon,
    component: SettingsView,
  },
  {
    order: 0,
    route: "/",
    label: "Home",
    icon: HomeIcon,
    component: HomeView,
  },
];
