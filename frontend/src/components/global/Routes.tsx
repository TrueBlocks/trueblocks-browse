import React from "react";
import { Route, Switch } from "wouter";
import classes from "@/App.module.css";
import BlocksView from "@/views/Blocks/BlocksView";
import HomeView from "@/views/Home/HomeView";
import NamesView from "@/views/Names/NamesView";
import SettingsView from "@/views/Settings/SettingsView";

function Routes() {
  return (
    <div className={classes.mainContent}>
      <Switch>
        <Route path="/blocks">
          <BlocksView />
        </Route>
        <Route path="/names">
          <NamesView />
        </Route>
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
