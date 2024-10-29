// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
import {
  // Icons
  ProjectIcon,
  HistoryIcon,
  MonitorsIcon,
  NamesIcon,
  AbisIcon,
  IndexesIcon,
  ManifestsIcon,
  StatusIcon,
  SettingsIcon,
  DaemonsIcon,
  WizardIcon,
  // Views
  ProjectView,
  HistoryView,
  MonitorsView,
  NamesView,
  AbisView,
  IndexesView,
  ManifestsView,
  StatusView,
  SettingsView,
  DaemonsView,
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
    order: 10,
    route: expandRoute("history"),
    label: "History",
    icon: HistoryIcon,
    component: HistoryView,
  },
  {
    order: 20,
    route: expandRoute("monitors"),
    label: "Monitors",
    icon: MonitorsIcon,
    component: MonitorsView,
  },
  {
    order: 30,
    route: expandRoute("names"),
    label: "Names",
    icon: NamesIcon,
    component: NamesView,
  },
  {
    order: 40,
    route: expandRoute("abis"),
    label: "Abis",
    icon: AbisIcon,
    component: AbisView,
  },
  {
    order: 50,
    route: expandRoute("indexes"),
    label: "Indexes",
    icon: IndexesIcon,
    component: IndexesView,
  },
  {
    order: 60,
    route: expandRoute("manifests"),
    label: "Manifest",
    icon: ManifestIcon,
    component: ManifestsView,
  },
  {
    order: 70,
    route: expandRoute("status"),
    label: "Status",
    icon: StatusIcon,
    component: StatusView,
  },
  {
    order: 80,
    route: expandRoute("settings"),
    label: "Settings",
    icon: SettingsIcon,
    component: SettingsView,
  },
  {
    order: 90,
    route: expandRoute("daemons"),
    label: "Daemons",
    icon: DaemonsIcon,
    component: DaemonsView,
  },
  {
    order: 100,
    route: expandRoute("wizard"),
    label: "Wizard",
    icon: WizardIcon,
    component: WizardView,
  },
  {
    order: 0,
    route: expandRoute(""),
    label: "Project",
    icon: ProjectIcon,
    component: ProjectView,
  },
];

//-----
// projects --> 0000-project
// historys --> 1010-history
// monitors --> 1020-monitors
// names --> 1030-names
// abis --> 1040-abis
// indexes --> 1050-indexes
// manifests --> 1060-manifests
// status --> 1070-status
// settings --> 1080-settings
// daemons --> 1090-daemons
// sessions --> 1100-session
// configs --> 2110-config
// wizards --> 2120-wizard
//-----
