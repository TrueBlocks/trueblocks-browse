// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
import {
  // Icons
  ProjectIcon,
  HistoryIcon,
  MonitorsIcon,
  SharingIcon,
  UnchainedIcon,
  SettingsIcon,
  DaemonsIcon,
  WizardIcon,
  // Views
  ProjectView,
  HistoryView,
  MonitorsView,
  SharingView,
  UnchainedView,
  SettingsView,
  DaemonsView,
  WizardView,
} from "@views";

export type Route = "" | "history" | "monitors" | "sharing" | "unchained" | "settings" | "daemons" | "wizard";

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
    label: "Project",
    icon: ProjectIcon,
    component: ProjectView,
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
    order: 1045,
    route: expandRoute("sharing"),
    label: "Sharing",
    icon: SharingIcon,
    component: SharingView,
  },
  {
    order: 1065,
    route: expandRoute("unchained"),
    label: "Unchained",
    icon: UnchainedIcon,
    component: UnchainedView,
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
    order: 1120,
    route: expandRoute("wizard"),
    label: "Wizard",
    icon: WizardIcon,
    component: WizardView,
  },
];
