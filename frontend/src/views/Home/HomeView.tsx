import React from "react";
import { Stack, Text, Title } from "@mantine/core";
import classes from "@/App.module.css";
import { View, ViewStatus } from "@/components/view";

function HomeView() {
  return (
    <View>
      <Title order={3}>Home View Title</Title>
      <Stack className={classes.mainContent}>
        <Text>Home View Content</Text>
      </Stack>
      <ViewStatus>Status / Progress</ViewStatus>
    </View>
  );
}

export default HomeView;
