import React, { useEffect, useState } from "react";
import classes from "@/App.module.css";
import { Stack, Title } from "@mantine/core";
import { View, ViewStatus, DataTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { types } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { indexColumns } from "./IndexesTable";
import { GetIndexesPage, GetIndexesCnt } from "@gocode/app/App";

// Find: NewViews
export function IndexesView() {
  const [count, setCount] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(false);
  const [loaded, setLoaded] = useState<boolean>(false);
  const [items, setItems] = useState<types.ChunkStats[]>([]);
  const { curItem, perPage } = useKeyboardPaging<types.ChunkStats>(items, count);

  useEffect(() => {
    if (loaded && !loading) {
      const fetch = async (currentItem: number, itemsPerPage: number) => {
        GetIndexesPage(currentItem, itemsPerPage).then((newItems) => {
          setItems(newItems);
        });
      };
      fetch(curItem, perPage);
    }
  }, [count, curItem, perPage, loaded, loading]);

  useEffect(() => {
    setLoading(true);
    const fetch = async () => {
      const cnt = await GetIndexesCnt();
      setCount(cnt);
      setLoaded(true);
    };
    fetch().finally(() => setLoading(false));
  }, []);

  const table = useReactTable({
    data: items,
    columns: indexColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <Stack className={classes.mainContent}>
        <Title order={3}>
          Indexes: showing record {curItem + 1}-{curItem + 1 + perPage - 1} of {count}
        </Title>
        <DataTable<types.ChunkStats> table={table} loading={loading} />
      </Stack>
      <ViewStatus />
    </View>
  );
}
