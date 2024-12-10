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

export type Route = "project" | "history" | "monitors" | "sharing" | "unchained" | "settings" | "daemons" | "wizard";

export type RouteItem = {
  order: number;
  route: string;
  label: string;
  icon: JSX.Element;
  component: React.ComponentType;
};

export const routeItems: RouteItem[] = [
  {
    order: 0,
    route: "project",
    label: "Project",
    icon: ProjectIcon,
    component: ProjectView,
  },
  {
    order: 1010,
    route: "history",
    label: "History",
    icon: HistoryIcon,
    component: HistoryView,
  },
  {
    order: 1025,
    route: "monitors",
    label: "Monitors",
    icon: MonitorsIcon,
    component: MonitorsView,
  },
  {
    order: 1050,
    route: "sharing",
    label: "Sharing",
    icon: SharingIcon,
    component: SharingView,
  },
  {
    order: 1065,
    route: "unchained",
    label: "Unchained",
    icon: UnchainedIcon,
    component: UnchainedView,
  },
  {
    order: 1080,
    route: "settings",
    label: "Settings",
    icon: SettingsIcon,
    component: SettingsView,
  },
  {
    order: 1090,
    route: "daemons",
    label: "Daemons",
    icon: DaemonsIcon,
    component: DaemonsView,
  },
  {
    order: 1120,
    route: "wizard",
    label: "Wizard",
    icon: WizardIcon,
    component: WizardView,
  },
];
