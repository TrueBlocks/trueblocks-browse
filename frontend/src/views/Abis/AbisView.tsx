import React, { useEffect, useState } from "react";
import classes from "@/App.module.css";
import { Stack, Title } from "@mantine/core";
import { View, ViewStatus, DataTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { types } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { abiColumns } from "./AbisTable";
import { GetAbisPage, GetAbisCnt } from "@gocode/app/App";

// Find: NewViews
export function AbisView() {
  const [count, setCount] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(false);
  const [loaded, setLoaded] = useState<boolean>(false);
  const [items, setItems] = useState<types.AbiFile[]>([]);
  const { curItem, perPage } = useKeyboardPaging<types.AbiFile>(items, count);

  useEffect(() => {
    if (loaded && !loading) {
      const fetch = async (currentItem: number, itemsPerPage: number) => {
        GetAbisPage(currentItem, itemsPerPage).then((newItems) => {
          setItems(newItems);
        });
      };
      fetch(curItem, perPage);
    }
  }, [count, curItem, perPage, loaded, loading]);

  useEffect(() => {
    setLoading(true);
    const fetch = async () => {
      const cnt = await GetAbisCnt();
      setCount(cnt);
      setLoaded(true);
    };
    fetch().finally(() => setLoading(false));
  }, []);

  const table = useReactTable({
    data: items,
    columns: abiColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <Stack className={classes.mainContent}>
        <Title order={3}>
          Abis: showing record {curItem + 1}-{curItem + 1 + perPage - 1} of {count}
        </Title>
        <DataTable<types.AbiFile> table={table} loading={loading} />
      </Stack>
      <ViewStatus />
    </View>
  );
}
