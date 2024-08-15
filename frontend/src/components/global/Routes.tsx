import React, { useEffect } from "react";
import { Route, Switch, useLocation } from "wouter";
import classes from "@/App.module.css";
import { GetLast } from "@gocode/app/App";
import { routeItems } from "@/Routes";

export const Routes = () => {
  const [, setLocation] = useLocation();

  useEffect(() => {
    (GetLast("route") || "/").then((route) => {
      setLocation(route);
    });
  }, [setLocation]);

  return (
    <div className={classes.mainContent}>
      <Switch>
        {routeItems.map((item) => (
          <Route key={item.route} path={item.route}>
            <item.component />
          </Route>
        ))}
      </Switch>
    </div>
  );
};
