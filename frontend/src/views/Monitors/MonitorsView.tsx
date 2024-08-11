import React, { useState, useEffect, ReactNode } from "react";
import { types } from "@gocode/models";
import { Stack } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns, createForm } from ".";
import classes from "@/App.module.css";
import { View, ViewStatus, ViewTitle, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetMonitors, GetMonitorsCnt } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";

export function MonitorsView() {
  const [count, setCount] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(false);
  const [loaded, setLoaded] = useState<boolean>(false);
  const [items, setItems] = useState<types.SummaryMonitor>({} as types.SummaryMonitor);
  const [monitors, setMonitors] = useState<types.Monitor[]>([]);
  const [refresh, setRefresh] = useState<boolean>(false);
  const { curItem, perPage } = useKeyboardPaging<types.Monitor>(monitors, count, [], 15);

  useEffect(() => {
    if (loaded && !loading) {
      const fetch = async (currentItem: number, itemsPerPage: number) => {
        GetMonitors(currentItem, itemsPerPage).then((items: types.SummaryMonitor) => {
          setItems(items);
          setMonitors(items.monitors || []);
        });
      };
      fetch(curItem, perPage);
      setRefresh(false);
    }
  }, [count, curItem, perPage, loaded, loading, refresh]);

  useEffect(() => {
    const handleRefresh = () => {
      setRefresh(true);
    };

    EventsOn("DAEMON", handleRefresh);
    return () => {
      EventsOff("DAEMON");
    };
  }, []);

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
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <Stack className={classes.mainContent}>
        <ViewTitle />
        <FormTable data={items} definition={createForm(table, curItem, count, perPage)} />
      </Stack>
      <ViewStatus />
    </View>
  );
}
