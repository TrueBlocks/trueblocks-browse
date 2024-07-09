import React from "react";
import { Route, Switch } from "wouter";
import classes from "./Routes.module.css";
import HomeView from "@/views/Home/HomeView";
import SettingsView from "@/views/Settings/SettingsView";

function Routes() {
  return (
    <div className={classes.container}>
      <Switch>
        <Route path="/settings">
          <SettingsView />
        </Route>
        <Route>
          <HomeView />
        </Route>
      </Switch>
    </div>
  );
}

export default Routes;
