import React, { useState, useEffect } from "react";
import classes from "@/App.module.css";
import { GetNamesPage, GetNamesCnt } from "@gocode/app/App";
import { types } from "@gocode/models";
import { getCoreRowModel, useReactTable, Table } from "@tanstack/react-table";
import { Stack, Title } from "@mantine/core";
import { nameColumns } from "./NameTable";
import { DataTable } from "@components";
import { View, ViewStatus } from "@components";
import { useKeyboardPaging } from "@hooks";

// TODO: This should use tabs per type (Regular, Custom, Prefund, Baddress, etc.)
// TODO: Or, at least have tags for each type.
export function NamesView() {
  const [count, setCount] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(false);
  const [loaded, setLoaded] = useState<boolean>(false);
  const [items, setItems] = useState<types.NameEx[]>([]);
  const { curItem, perPage } = useKeyboardPaging<types.NameEx>(items, count);

  useEffect(() => {
    if (loaded && !loading) {
      const fetch = async (currentItem: number, itemsPerPage: number) => {
        const newItems = await GetNamesPage(currentItem, itemsPerPage);
        setItems(newItems);
      };
      fetch(curItem, perPage);
    }
  }, [count, curItem, perPage, loaded, loading]);

  useEffect(() => {
    setLoading(true);
    const fetch = async () => {
      const cnt = await GetNamesCnt();
      setCount(cnt);
      setLoaded(true);
    };
    fetch().finally(() => setLoading(false));
  }, []);

  const table = useReactTable({
    data: items,
    columns: nameColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <Stack className={classes.mainContent}>
        <Title order={3}>
          Names: showing record {curItem + 1}-{curItem + 1 + perPage - 1} of {count}
        </Title>
        <DataTable<types.NameEx> table={table} loading={loading} />
      </Stack>
      <ViewStatus />
    </View>
  );
}
