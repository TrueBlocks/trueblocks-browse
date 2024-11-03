import { useEffect } from "react";
import { Route as WouterRoute, Switch, useLocation } from "wouter";
import { GetRoute } from "@gocode/app/App";
import { routeItems, RouteItem } from "@layout";
import { useAppState } from "@state";

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
