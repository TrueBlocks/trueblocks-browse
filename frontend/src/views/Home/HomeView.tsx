import React from "react";
import { Stack, Text, Title } from "@mantine/core";
import classes from "@/App.module.css";
import { View, ViewStatus, ViewTitle } from "@components";

export function HomeView() {
  return (
    <View>
      <Stack className={classes.mainContent}>
        <ViewTitle />
        <Text>Home View Content</Text>
      </Stack>
      <ViewStatus />
    </View>
  );
}
