import React from "react";
import { Stack, Text, Title } from "@mantine/core";
import classes from "@/App.module.css";
import { View, ViewStatus, ViewTitle } from "@components";

export function HomeView() {
  return (
    <View>
      <ViewTitle />
      <Stack className={classes.mainContent}>
        <Text>Home View Content</Text>
      </Stack>
      <ViewStatus />
    </View>
  );
}
