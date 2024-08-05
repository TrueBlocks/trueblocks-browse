import React, { useState, useEffect, ReactNode } from "react";
import { types } from "@gocode/models";
import { Title, Stack } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { nameColumns, NameInstance, createNameForm } from ".";
import classes from "@/App.module.css";
import { View, ViewStatus, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetNames, GetNamesCnt } from "@gocode/app/App";

export function NamesView() {
  const [count, setCount] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(false);
  const [loaded, setLoaded] = useState<boolean>(false);
  const [items, setItems] = useState<types.NameSummary>({} as types.NameSummary);
  const [names, setNames] = useState<types.Name[]>([]);
  const { curItem, perPage } = useKeyboardPaging<types.Name>(names, count, [], 15);

  useEffect(() => {
    if (loaded && !loading) {
      const fetch = async (currentItem: number, itemsPerPage: number) => {
        GetNames(currentItem, itemsPerPage).then((names: types.NameSummary) => {
          setItems(names);
          setNames(names.names || []);
        });
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
    data: items.names || [], // Pass the names array or an empty array if undefined
    columns: nameColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <Stack className={classes.mainContent}>
        <Title order={3}>Names View</Title>
        <FormTable data={items} definition={createNameForm(table)} />;{" "}
      </Stack>
      <ViewStatus />
    </View>
  );
}
