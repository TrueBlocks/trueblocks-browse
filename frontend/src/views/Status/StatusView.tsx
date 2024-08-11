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
  const [loading, setLoading] = useState<boolean>(false);
  const [loaded, setLoaded] = useState<boolean>(false);
  const [summaryItem, setSummaryItem] = useState<types.SummaryStatus>({} as types.SummaryStatus);
  const [refresh, setRefresh] = useState<boolean>(false);
  const [count, setCount] = useState<number>(0);
  const pager = useKeyboardPaging(count, [], 10);

  useEffect(() => {
    if (loaded && !loading) {
      const fetch = async (currentItem: number, itemsPerPage: number) => {
        GetStatus(currentItem, itemsPerPage).then((item: types.SummaryStatus) => {
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
      const cnt = await GetStatusCnt();
      setCount(cnt);
      setLoaded(true);
    };
    fetch().finally(() => setLoading(false));
  }, []);

  const table = useReactTable({
    data: summaryItem.caches || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <Stack className={classes.mainContent}>
        <ViewTitle />
        <FormTable data={summaryItem} definition={createForm(table, pager.curItem, count, pager.perPage)} />
      </Stack>
      <ViewStatus />
    </View>
  );
}
