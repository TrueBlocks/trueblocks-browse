import React, { useState, useEffect, ReactNode } from "react";
import { types } from "@gocode/models";
import { Stack } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns, createForm } from ".";
import classes from "@/App.module.css";
import { View, ViewStatus, ViewTitle, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetStatus, GetStatusCnt } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";

export function StatusView() {
  const [count, setCount] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(false);
  const [loaded, setLoaded] = useState<boolean>(false);
  const [items, setItems] = useState<types.SummaryStatus>({} as types.SummaryStatus);
  const [caches, setCaches] = useState<types.CacheItem[]>([]);
  const [refresh, setRefresh] = useState<boolean>(false);
  const { curItem, perPage } = useKeyboardPaging<types.CacheItem>(caches, count, [], 8);

  useEffect(() => {
    if (loaded && !loading) {
      const fetch = async (currentItem: number, itemsPerPage: number) => {
        GetStatus(currentItem, itemsPerPage).then((items: types.SummaryStatus) => {
          setItems(items);
          setCaches(items.caches || []);
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
      const cnt = await GetStatusCnt();
      setCount(cnt);
      setLoaded(true);
    };
    fetch().finally(() => setLoading(false));
  }, []);

  const table = useReactTable({
    data: items.caches || [], // Pass the caches array or an empty array if undefined
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <Stack className={classes.mainContent}>
        <ViewTitle />
        <FormTable data={items} definition={createForm(table)} />;{" "}
      </Stack>
      <ViewStatus />
    </View>
  );
}
