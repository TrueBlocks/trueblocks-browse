import React, { ReactNode } from "react";
import { Stack } from "@mantine/core";
import { ViewTitle, ViewStatus } from "@components";
import classes from "@/App.module.css";
import { ViewStateProvider } from "@state";
import { Route } from "@/Routes";

export function View({
  route,
  nItems = -1,
  children,
}: {
  route: Route;
  nItems?: number;
  children: ReactNode;
}) {
  return (
    <ViewStateProvider route={route} nItems={nItems}>
      <Stack className={classes.mainContent}>
        <ViewTitle />
        {children}
      </Stack>
      <ViewStatus />
    </ViewStateProvider>
  );
}
