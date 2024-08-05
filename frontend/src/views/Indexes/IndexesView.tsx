import React, { useState, useEffect, ReactNode } from "react";
import { types } from "@gocode/models";
import { Title, Stack } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { indexColumns, createIndexForm } from ".";
import classes from "@/App.module.css";
import { View, ViewStatus, ViewTitle, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetIndex, GetIndexCnt } from "@gocode/app/App";

export function IndexesView() {
  const [count, setCount] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(false);
  const [loaded, setLoaded] = useState<boolean>(false);
  const [items, setItems] = useState<types.IndexSummary>({} as types.IndexSummary);
  const [chunks, setChunks] = useState<types.ChunkStats[]>([]);
  const { curItem, perPage } = useKeyboardPaging<types.ChunkStats>(chunks, count, [], 15);

  useEffect(() => {
    if (loaded && !loading) {
      const fetch = async (currentItem: number, itemsPerPage: number) => {
        GetIndex(currentItem, itemsPerPage).then((index: types.IndexSummary) => {
          setItems(index);
          setChunks(index.chunks || []);
        });
      };
      fetch(curItem, perPage);
    }
  }, [count, curItem, perPage, loaded, loading]);

  useEffect(() => {
    setLoading(true);
    const fetch = async () => {
      const cnt = await GetIndexCnt();
      setCount(cnt);
      setLoaded(true);
    };
    fetch().finally(() => setLoading(false));
  }, []);

  const table = useReactTable({
    data: items.chunks || [], // Pass the chunks array or an empty array if undefined
    columns: indexColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <Stack className={classes.mainContent}>
        <ViewTitle />
        <FormTable data={items} definition={createIndexForm(table)} />;{" "}
      </Stack>
      <ViewStatus />
    </View>
  );
}
