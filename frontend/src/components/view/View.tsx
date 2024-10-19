import { ReactNode } from "react";
import { Stack } from "@mantine/core";
import { ViewTitle } from "@components";
import classes from "./View.module.css";

export const View = ({ children }: { children: ReactNode }) => {
  return (
    <Stack className={classes.viewContainer}>
      <ViewTitle />
      {children}
    </Stack>
  );
};
