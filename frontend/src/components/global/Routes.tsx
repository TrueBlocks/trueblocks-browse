import React from "react";
import { Route, Switch } from "wouter";
import classes from "./Routes.module.css";
import HomeView from "@/views/Home/HomeView";
import NamesView from "@/views/Names/NamesView";
import SettingsView from "@/views/Settings/SettingsView";

function Routes() {
  return (
    <div className={classes.container}>
      <Switch>
        {/* Settings */}
        <Route path="/names">
          <NamesView />
        </Route>
        <Route path="/settings">
          <SettingsView />
        </Route>

        {/* Default route */}
        <Route>
          <HomeView />
        </Route>
      </Switch>
    </div>
  );
}

export default Routes;
