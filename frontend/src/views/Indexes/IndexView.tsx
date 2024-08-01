import React from "react";
import classes from "@/App.module.css";
import { Stack, Title } from "@mantine/core";
import { View, ViewStatus, DataTable } from "@components";

export function IndexView() {
  return (
    <View>
      <Stack className={classes.mainContent}>
        <Title order={3}>
          Indexes
          {/* : showing record {curItem + 1}-{curItem + 1 + perPage - 1} of {count} */}
        </Title>
        {/* <DataTable<types.MonitorEx> table={table} loading={loading} /> */}
        Indexes View
      </Stack>
      <ViewStatus />
    </View>
  );
}
