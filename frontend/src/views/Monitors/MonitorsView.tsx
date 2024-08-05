import React, { useState, useEffect, ReactNode } from "react";
import { types } from "@gocode/models";
import { Title, Stack } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { monitorColumns, createMonitorForm } from ".";
import classes from "@/App.module.css";
import { View, ViewStatus, ViewTitle, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetMonitors, GetMonitorsCnt } from "@gocode/app/App";

export function MonitorsView() {
  const [count, setCount] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(false);
  const [loaded, setLoaded] = useState<boolean>(false);
  const [items, setItems] = useState<types.SummaryMonitor>({} as types.SummaryMonitor);
  const [monitors, setMonitors] = useState<types.Monitor[]>([]);
  const { curItem, perPage } = useKeyboardPaging<types.Monitor>(monitors, count, [], 15);

  useEffect(() => {
    if (loaded && !loading) {
      const fetch = async (currentItem: number, itemsPerPage: number) => {
        GetMonitors(currentItem, itemsPerPage).then((monitors: types.SummaryMonitor) => {
          setItems(monitors);
          setMonitors(monitors.monitors || []);
        });
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
    data: items.monitors || [], // Pass the monitors array or an empty array if undefined
    columns: monitorColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <Stack className={classes.mainContent}>
        <ViewTitle />
        <FormTable data={items} definition={createMonitorForm(table)} />;{" "}
      </Stack>
      <ViewStatus />
    </View>
  );
}
