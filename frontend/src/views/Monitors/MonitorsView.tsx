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
  const [loading, setLoading] = useState<boolean>(false);
  const [loaded, setLoaded] = useState<boolean>(false);
  const [summaryItem, setSummaryItem] = useState<types.MonitorContainer>({} as types.MonitorContainer);
  const [refresh, setRefresh] = useState<boolean>(false);
  const [count, setCount] = useState<number>(0);
  const pager = useKeyboardPaging(count, [], 15);

  useEffect(() => {
    if (loaded && !loading) {
      const fetch = async (currentItem: number, itemsPerPage: number) => {
        GetMonitors(currentItem, itemsPerPage).then((item: types.MonitorContainer) => {
          setSummaryItem(item);
        });
      };
      fetch(pager.curItem, pager.perPage);
      setRefresh(false);
    }
  }, [count, pager, loaded, loading, refresh]);

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
    data: summaryItem.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <Stack className={classes.mainContent}>
        <ViewTitle />
        <FormTable data={summaryItem} definition={createForm(table, pager)} />
      </Stack>
      <ViewStatus />
    </View>
  );
}
