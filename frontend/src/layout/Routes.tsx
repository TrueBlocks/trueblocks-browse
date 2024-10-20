import { useEffect } from "react";
import { Route as WouterRoute, Switch, useLocation } from "wouter";
import { GetRoute } from "@gocode/app/App";
import { useAppState } from "@state";
import {
  // Icons
  ProjectIcon,
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
  ProjectView,
  HistoryView,
  NamesView,
  MonitorsView,
  IndexesView,
  ManifestsView,
  AbisView,
  DaemonsView,
  StatusView,
  SettingsView,
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

function expandRoute(r: Route): string {
  if (r === "history") {
    return "/" + r + "/:address";
  }

  return "/" + r;
}

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
    icon: MonitorIcon,
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

export const Routes = () => {
  const [, setLocation] = useLocation();
  const { isConfigured } = useAppState();

  useEffect(() => {
    GetRoute().then((route) => {
      setLocation(route);
    });
  }, [setLocation]);

  const routes = routeItems
    .filter((item: RouteItem) => (isConfigured ? true : item.route === "/wizard"))
    .sort((a, b) => a.order - b.order);

  return (
    <Switch>
      {routes.map((item) => (
        <WouterRoute key={item.route} path={item.route}>
          <item.component />
        </WouterRoute>
      ))}
    </Switch>
  );
};