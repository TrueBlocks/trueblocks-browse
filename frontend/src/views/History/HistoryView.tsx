import React, { useState, useEffect } from "react";
import { useParams } from "wouter";
import classes from "@/App.module.css";
import { GetHistoryPage, GetHistoryCnt } from "@gocode/app/App";
import { app } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { Stack, Title } from "@mantine/core";
import { txColumns } from "./HistoryTable";
import { EditableSelect, View, ViewStatus } from "@components";
import { useKeyboardPaging } from "@hooks";
import { DataTable } from "@components";

// Find: NewViews
export function HistoryView() {
  const [address, setAddress] = useState<string>("trueblocks.eth");
  const [count, setCount] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(false);
  const [loaded, setLoaded] = useState<boolean>(false);
  const [items, setItems] = useState<app.TransactionEx[]>([]);
  const { curItem, perPage } = useKeyboardPaging<app.TransactionEx>(items, count, [address]);

  const params = useParams();
  const addr = params.address;

  useEffect(() => {
    if (loaded && !loading) {
      const fetch = async (addr: string, currentItem: number, itemsPerPage: number) => {
        const newItems = await GetHistoryPage(addr, currentItem, itemsPerPage);
        setItems(newItems);
      };
      fetch(address, curItem, perPage);
    }
  }, [count, curItem, perPage, address, loaded, loading]);

  useEffect(() => {
    setLoading(true);
    try {
      const fetch = async (addr: string) => {
        const cnt = await GetHistoryCnt(addr);
        setCount(cnt);
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
    data: items,
    columns: txColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <Stack className={classes.mainContent}>
        <Title order={3}>
          History of {address}: showing record {curItem + 1}-{curItem + 1 + perPage - 1} of {count}
        </Title>
        <EditableSelect
          value={address}
          onChange={(value) => setAddress(value)}
          label="Select or enter an address or ENS name"
          placeholder="Enter or select an address"
        />
        <DataTable<app.TransactionEx> table={table} loading={loading} />
      </Stack>
      <ViewStatus />
    </View>
  );
}
