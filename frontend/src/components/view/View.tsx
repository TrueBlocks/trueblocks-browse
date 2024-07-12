import React, { ReactNode } from "react";
import { Stack } from "@mantine/core";
import classes from "@/App.module.css";

export function View(params: { title?: string; children: ReactNode }) {
  return <Stack className={classes.mainContent}>{params.children}</Stack>;
}
