import React from "react";
import { Stack, Title, Checkbox, InputLabel } from "@mantine/core";
import classes from "@/App.module.css";
import { View, ViewStatus } from "@components";

export function SettingsView() {
  return (
    <View>
      <Title order={3}>Settings View Header</Title>
      <Stack className={classes.mainContent}>
        <InputLabel>
          <Checkbox label={"A checkbox"} />
        </InputLabel>
      </Stack>
      <ViewStatus />
    </View>
  );
}
