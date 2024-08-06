import React, { useState, useEffect } from "react";
import { useParams } from "wouter";
import { types } from "@gocode/models";
import { Title, Stack } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { transactionColumns, createTransactionForm } from ".";
import classes from "@/App.module.css";
import { View, ViewStatus, ViewTitle, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetHistory, GetHistoryCnt } from "@gocode/app/App";

export function HistoryView() {
  const [address, setAddress] = useState<string>("trueblocks.eth");
  const [count, setCount] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(false);
  const [loaded, setLoaded] = useState<boolean>(false);
  const [items, setItems] = useState<types.SummaryTransaction>({} as types.SummaryTransaction);
  const [txs, setTxs] = useState<types.Transaction[]>([]);
  const { curItem, perPage } = useKeyboardPaging<types.Transaction>(txs, count, [address], 15);

  const params = useParams();
  const addr = params.address;

  useEffect(() => {
    if (loaded && !loading) {
      const fetch = async (addr: string, currentItem: number, itemsPerPage: number) => {
        GetHistory(addr, currentItem, itemsPerPage).then((txSummary: types.SummaryTransaction) => {
          setItems(txSummary);
          setTxs(txSummary.transactions || []);
        });
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
    data: items.transactions || [], // Pass the chunks array or an empty array if undefined
    columns: transactionColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <Stack className={classes.mainContent}>
        <ViewTitle />
        <FormTable data={items} definition={createTransactionForm(table)} />;{" "}
      </Stack>
      <ViewStatus />
    </View>
  );
}
