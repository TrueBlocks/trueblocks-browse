// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
import {
  ProjectIcon,
  DashboardIcon,
  HistoryIcon,
  MonitorsIcon,
  NamesIcon,
  AbisIcon,
  IndexesIcon,
  ManifestsIcon,
  StatusIcon,
  SettingsIcon,
  DaemonsIcon,
  SessionIcon,
  ConfigIcon,
  WizardIcon,
  // Views
  DashboardView,
  HistoryView,
  MonitorsView,
  NamesView,
  AbisView,
  IndexesView,
  ManifestsView,
  StatusView,
  SettingsView,
  DaemonsView,
  SessionView,
  ConfigView,
  WizardView,
} from "@views";

export type Route =
  | ""
  | "history"
  | "monitors"
  | "names"
  | "abis"
  | "indexes"
  | "manifests"
  | "status"
  | "settings"
  | "daemons"
  | "session"
  | "config"
  | "wizard";

export type RouteItem = {
  order: number;
  route: string;
  label: string;
  icon: JSX.Element;
  component: React.ComponentType;
};

const expandRoute = (r: Route): string => {
  if (r === "history") {
    return "/" + r + "/:address";
  }

  return "/" + r;
};

export const routeItems: RouteItem[] = [
  {
    order: 0,
    route: expandRoute(""),
    label: "Dashboard",
    icon: DashboardIcon,
    component: DashboardView,
  },
  {
    order: 1010,
    route: expandRoute("history"),
    label: "History",
    icon: HistoryIcon,
    component: HistoryView,
  },
  {
    order: 1020,
    route: expandRoute("monitors"),
    label: "Monitors",
    icon: MonitorsIcon,
    component: MonitorsView,
  },
  {
    order: 1030,
    route: expandRoute("names"),
    label: "Names",
    icon: NamesIcon,
    component: NamesView,
  },
  {
    order: 1040,
    route: expandRoute("abis"),
    label: "Abis",
    icon: AbisIcon,
    component: AbisView,
  },
  {
    order: 1050,
    route: expandRoute("indexes"),
    label: "Indexes",
    icon: IndexesIcon,
    component: IndexesView,
  },
  {
    order: 1060,
    route: expandRoute("manifests"),
    label: "Manifests",
    icon: ManifestsIcon,
    component: ManifestsView,
  },
  {
    order: 1070,
    route: expandRoute("status"),
    label: "Status",
    icon: StatusIcon,
    component: StatusView,
  },
  {
    order: 1080,
    route: expandRoute("settings"),
    label: "Settings",
    icon: SettingsIcon,
    component: SettingsView,
  },
  {
    order: 1090,
    route: expandRoute("daemons"),
    label: "Daemons",
    icon: DaemonsIcon,
    component: DaemonsView,
  },
  {
    order: 1100,
    route: expandRoute("session"),
    label: "Session",
    icon: SessionIcon,
    component: SessionView,
  },
  {
    order: 1110,
    route: expandRoute("config"),
    label: "Config",
    icon: ConfigIcon,
    component: ConfigView,
  },
  {
    order: 1120,
    route: expandRoute("wizard"),
    label: "Wizard",
    icon: WizardIcon,
    component: WizardView,
  },
];
