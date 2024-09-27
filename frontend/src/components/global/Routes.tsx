import { useEffect } from "react";
import { Route, Switch, useLocation } from "wouter";
import { routeItems, RouteItem } from "@/Routes";
import { GetSessionVal } from "@gocode/app/App";
import { useAppState } from "@state";
import classes from "@/App.module.css";

export const Routes = () => {
  const [, setLocation] = useLocation();
  const { isConfigured } = useAppState();

  useEffect(() => {
    (GetSessionVal("route") || "/").then((route) => {
      setLocation(route);
    });
  }, [setLocation]);

  // item.route !== "/wizard"
  const routes = routeItems
    .filter((item: RouteItem) => (isConfigured ? true : item.route === "/wizard"))
    .sort((a, b) => a.order - b.order);

  return (
    <div className={classes.mainContent}>
      <Switch>
        {routes.map((item) => (
          <Route key={item.route} path={item.route}>
            <item.component />
          </Route>
        ))}
      </Switch>
    </div>
  );
};
