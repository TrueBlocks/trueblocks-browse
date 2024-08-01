import React from "react";
import classes from "@/App.module.css";
import { Stack, Title } from "@mantine/core";
import { View, ViewStatus, DataTable } from "@components";

export function AbisView() {
  return (
    <View>
      <Stack className={classes.mainContent}>
        <Title order={3}>
          Abis
          {/* : showing record {curItem + 1}-{curItem + 1 + perPage - 1} of {count} */}
        </Title>
        {/* <DataTable<types.MonitorEx> table={table} loading={loading} /> */}
        Abis View
      </Stack>
      <ViewStatus />
    </View>
  );
}
