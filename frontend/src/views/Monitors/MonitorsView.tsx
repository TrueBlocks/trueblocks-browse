import React, { useState, useEffect } from "react";
import classes from "@/App.module.css";
import { GetMonitorsPage, GetMonitorsCnt } from "@gocode/app/App";
import { app } from "@gocode/models";
import { getCoreRowModel, useReactTable, Table } from "@tanstack/react-table";
import { Stack, Title } from "@mantine/core";
import { monitorColumns } from "./MonitorTable";
import { DataTable } from "@components";
import { View, ViewStatus } from "@components";
import { useKeyboardPaging } from "@hooks";

// Find: NewViews
export function MonitorsView() {
  const [count, setCount] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(false);
  const [loaded, setLoaded] = useState<boolean>(false);
  const [items, setItems] = useState<app.MonitorEx[]>([]);
  const { curItem, perPage } = useKeyboardPaging<app.MonitorEx>(items, count);

  useEffect(() => {
    if (loaded && !loading) {
      const fetch = async (currentItem: number, itemsPerPage: number) => {
        const newItems = await GetMonitorsPage(currentItem, itemsPerPage);
        setItems(newItems);
      };
      fetch(curItem, perPage);
    }
  }, [count, curItem, perPage, loaded, loading]);

  useEffect(() => {
    setLoading(true);
    const fetch = async () => {
      const cnt = await GetMonitorsCnt();
      setCount(cnt);
      setLoaded(true);
    };
    fetch().finally(() => setLoading(false));
  }, []);

  const table = useReactTable({
    data: items,
    columns: monitorColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <Stack className={classes.mainContent}>
        <Title order={3}>
          Monitors: showing record {curItem + 1}-{curItem + 1 + perPage - 1} of {count}
        </Title>
        <DataTable<app.MonitorEx> table={table} loading={loading} />
      </Stack>
      <ViewStatus />
    </View>
  );
}
