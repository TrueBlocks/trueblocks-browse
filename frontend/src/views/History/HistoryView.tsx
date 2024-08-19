import React, { useState, useEffect } from "react";
import { useParams, useLocation } from "wouter";
import { types, messages } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns, createForm } from ".";
import { View, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetLastSub, HistoryPage, GetHistoryCnt } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";

export function HistoryView() {
  const [address, setAddress] = useState<string>("");
  const [summaryItem, setSummaryItem] = useState<types.TransactionContainer>({} as types.TransactionContainer);
  const [count, setCount] = useState<number>(0);
  const [location, _] = useLocation();
  const pager = useKeyboardPaging("history", count, [address], 15);

  useEffect(() => {
    if (address !== "") {
      // console.log("HistoryView::preFetch [" + address + "] [" + addr + "]");
      const fetch = async (addy: string, currentItem: number, itemsPerPage: number) => {
        HistoryPage(addy, currentItem, itemsPerPage).then((item: types.TransactionContainer) => {
          if (item) {
            setSummaryItem(item);
            GetHistoryCnt(addy).then((cnt: number) => {
              setCount(cnt);
            });
          }
        });
      };
      fetch(address, pager.curItem, pager.perPage);
    }
  }, [address, pager.curItem, pager.perPage]);

  let addr = useParams().address;
  useEffect(() => {
    // console.log("HistoryView::addr", addr);
    if (addr === ":address") {
      GetLastSub("/history").then((a) => (addr = a));
    }
    if (addr && addr !== "" && addr !== ":address") {
      // console.log("HistoryView::setAddress", addr);
      setAddress(addr);
    } else {
      // console.log("HistoryView::skipping", addr);
    }
  }, [addr]);

  const table = useReactTable({
    data: summaryItem.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      {/* <div>{`addr: ${addr}`}</div>
      <div>{`address: ${address}`}</div>
      <div>{`location: ${location}`}</div> */}
      <FormTable data={summaryItem} definition={createForm(table, pager)} />
    </View>
  );
}
