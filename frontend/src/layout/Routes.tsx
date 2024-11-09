import { useEffect, useState } from "react";
import { Route as WouterRoute, Switch, useLocation } from "wouter";
import { GetRoute } from "@gocode/app/App";
import { types } from "@gocode/models";
import { routeItems, RouteItem } from "@layout";
import { useAppState } from "@state";

export const Routes = () => {
  const [, setLocation] = useLocation();
  const { wizard } = useAppState();
  const [routes, setRoutes] = useState<RouteItem[]>([]);

  useEffect(() => {
    GetRoute().then((route) => {
      setLocation(route);
    });
  }, [setLocation]);

  useEffect(() => {
    const r = routeItems
      .filter((item: RouteItem) => (wizard.state == types.WizState.FINISHED ? true : item.route === "/wizard"))
      .sort((a, b) => a.order - b.order);
    setRoutes(r);
  }, [wizard.state]);

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
