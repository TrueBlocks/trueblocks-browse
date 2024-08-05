import React, { useState, useEffect, ReactNode } from "react";
import { types } from "@gocode/models";
import { Title, Stack } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { cacheColumns, createStatusForm } from ".";
import classes from "@/App.module.css";
import { View, ViewStatus, ViewTitle, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetStatus, GetStatusCnt } from "@gocode/app/App";

export function StatusView() {
  const [count, setCount] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(false);
  const [loaded, setLoaded] = useState<boolean>(false);
  const [items, setItems] = useState<types.StatusEx>({} as types.StatusEx);
  const [caches, setCaches] = useState<types.CacheItem[]>([]);
  const { curItem, perPage } = useKeyboardPaging<types.CacheItem>(caches, count, [], 8);

  useEffect(() => {
    if (loaded && !loading) {
      const fetch = async (currentItem: number, itemsPerPage: number) => {
        GetStatus(currentItem, itemsPerPage).then((status: types.StatusEx) => {
          setItems(status);
          setCaches(status.caches || []);
        });
      };
      fetch(curItem, perPage);
    }
  }, [count, curItem, perPage, loaded, loading]);

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
    columns: cacheColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <Stack className={classes.mainContent}>
        <ViewTitle />
        <FormTable data={items} definition={createStatusForm(table)} />;{" "}
      </Stack>
      <ViewStatus />
    </View>
  );
}
