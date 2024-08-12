import React, { useState, useEffect } from "react";
import { useParams } from "wouter";
import { types } from "@gocode/models";
import { Stack } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns, createForm } from ".";
import classes from "@/App.module.css";
import { View, ViewStatus, ViewTitle, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetHistory, GetHistoryCnt } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";

export function HistoryView() {
  const [address, setAddress] = useState<string>("trueblocks.eth");
  const [loading, setLoading] = useState<boolean>(false);
  const [loaded, setLoaded] = useState<boolean>(false);
  const [summaryItem, setSummaryItem] = useState<types.SummaryTransaction>({} as types.SummaryTransaction);
  const [count, setCount] = useState<number>(0);
  const pager = useKeyboardPaging(count, [address], 15);

  const params = useParams();
  const addr = params.address;

  useEffect(() => {
    if (loaded && !loading) {
      const fetch = async (addr: string, currentItem: number, itemsPerPage: number) => {
        GetHistory(addr, currentItem, itemsPerPage).then((item: types.SummaryTransaction) => {
          setSummaryItem(item);
        });
      };
      fetch(address, pager.curItem, pager.perPage);
    }
  }, [count, pager.curItem, pager.perPage, loaded, loading, address]);

  useEffect(() => {
    setLoading(true);
    try {
      const fetch = async (addr: string) => {
        const cnt = await GetHistoryCnt(addr);
        setCount(cnt);
        setLoaded(true);
      };
      fetch(address);
      setLoaded(true);
    } finally {
      setLoading(false);
    }
  }, [address]);

  useEffect(() => {
    if (addr && addr !== "" && addr !== ":address") {
      setAddress(addr);
    } else {
      setAddress("trueblocks.eth");
    }
  }, [addr]);

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
