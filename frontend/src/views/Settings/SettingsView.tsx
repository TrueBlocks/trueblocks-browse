import React from "react";
import { Stack, Title, Checkbox, InputLabel } from "@mantine/core";
import classes from "@/App.module.css";
import { View, ViewStatus, ViewTitle } from "@components";

export function SettingsView() {
  return (
    <View>
      <Stack className={classes.mainContent}>
        <ViewTitle />
        <InputLabel>
          <Checkbox label={"A checkbox"} />
        </InputLabel>
      </Stack>
      <ViewStatus />
    </View>
  );
}
