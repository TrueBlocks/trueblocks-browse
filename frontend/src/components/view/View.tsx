import { ReactNode } from "react";
import { Stack } from "@mantine/core";
import { ViewTitle } from "@components";
import { useViewState } from "@state";
import classes from "./View.module.css";

export const View = ({ children }: { children: ReactNode }) => {
  const { collapsed } = useViewState();
  return (
    <Stack className={classes.viewContainer}>
      <div>{`collapsed: ${collapsed}`}</div>
      <ViewTitle />
      {children}
    </Stack>
  );
};
