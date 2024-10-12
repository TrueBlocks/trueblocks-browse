import { useEffect } from "react";
import { Route, Switch, useLocation } from "wouter";
import { routeItems, RouteItem } from "@/Routes";
import { GetSessionVal } from "@gocode/app/App";
import { useAppState } from "@state";

export const Routes = () => {
  const [, setLocation] = useLocation();
  const { isConfigured } = useAppState();

  useEffect(() => {
    (GetSessionVal("route") || "/").then((route) => {
      setLocation(route);
    });
  }, [setLocation]);

  const routes = routeItems
    .filter((item: RouteItem) => (isConfigured ? true : item.route === "/wizard"))
    .sort((a, b) => a.order - b.order);

  return (
    <Switch>
      {routes.map((item) => (
        <Route key={item.route} path={item.route}>
          <item.component />
        </Route>
      ))}
    </Switch>
  );
};
