import React, { useState, useEffect, ReactNode } from "react";
import { types } from "@gocode/models";
import { Title, Stack } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { abiColumns, createAbiForm } from ".";
import classes from "@/App.module.css";
import { View, ViewStatus, ViewTitle, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetAbis, GetAbisCnt } from "@gocode/app/App";

export function AbisView() {
  const [count, setCount] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(false);
  const [loaded, setLoaded] = useState<boolean>(false);
  const [items, setItems] = useState<types.AbiSummary>({} as types.AbiSummary);
  const [chunks, setChunks] = useState<types.Abi[]>([]);
  const { curItem, perPage } = useKeyboardPaging<types.Abi>(chunks, count, [], 15);

  useEffect(() => {
    if (loaded && !loading) {
      const fetch = async (currentItem: number, itemsPerPage: number) => {
        GetAbis(currentItem, itemsPerPage).then((abis: types.AbiSummary) => {
          setItems(abis);
          setChunks(abis.chunks || []);
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
    data: items.chunks || [], // Pass the chunks array or an empty array if undefined
    columns: abiColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <Stack className={classes.mainContent}>
        <ViewTitle />
        <FormTable data={items} definition={createAbiForm(table)} />;{" "}
      </Stack>
      <ViewStatus />
    </View>
  );
}
