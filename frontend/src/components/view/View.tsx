import React, { ReactNode } from "react";
import { Stack } from "@mantine/core";
import { ViewTitle, ViewStatus } from "@components";
import classes from "@/App.module.css";

export function View({ title, children }: { title?: string; children: ReactNode }) {
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
