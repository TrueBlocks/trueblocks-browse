import React, { useState, useEffect } from "react";
import { useParams, useLocation } from "wouter";
import { types, messages } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns } from "./HistoryTable";
import { View, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetLastSub, HistoryPage, GetHistoryCnt } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";
import { GroupDefinition, DataTable, Pager } from "@components";

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
      <FormTable data={summaryItem} definition={createHistoryForm(table, pager)} />
    </View>
  );
}

type theInstance = InstanceType<typeof types.TransactionContainer>;
function createHistoryForm(table: any, pager: Pager): GroupDefinition<theInstance>[] {
  return [
    {
      title: "Transaction Data",
      colSpan: 6,
      fields: [
        { label: "address", type: "text", accessor: "address" },
        { label: "name", type: "text", accessor: "name" },
        { label: "balance", type: "text", accessor: "balance" },
      ],
    },
    {
      title: "Future Use",
      colSpan: 6,
      fields: [
        { label: "nTransactions", type: "int", accessor: "nItems" },
        { label: "nEvents", type: "int", accessor: "nEvents" },
        { label: "nTokens", type: "int", accessor: "nTokens" },
        { label: "nErrors", type: "int", accessor: "nErrors" },
      ],
    },
    {
      title: "Files",
      fields: [],
      components: [
        {
          component: <DataTable<types.Transaction> table={table} loading={false} pagerName="history" />,
        },
      ],
    },
  ];
}
