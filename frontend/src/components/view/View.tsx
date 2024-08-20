import React, { ReactNode } from "react";
import { Stack } from "@mantine/core";
import { ViewTitle, ViewStatus } from "@components";
import classes from "@/App.module.css";
import { ViewStateProvider } from "@state";
import { Route } from "@/Routes";

export function View({ children }: { children: ReactNode }) {
  return (
    <>
      <Stack className={classes.mainContent}>
        <ViewTitle />
        {children}
      </Stack>
      <ViewStatus />
    </>
  );
}
