import React, { useState, useEffect } from "react";
import { useParams } from "wouter";
import { types, base } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns } from "./HistoryTable";
import { View, FormTable, DataTable, GroupDefinition, Pager } from "@components";
import { GetLastSub, HistoryPage } from "@gocode/app/App";
import { useAppState } from "@state";

export function HistoryView() {
  const { address, setAddress, history, setHistory } = useAppState();
  const historyPgr = useAppState().getPager("history");

  let addr = useParams().address;
  useEffect(() => {
    if (addr === ":address") {
      GetLastSub("/history").then((a) => (addr = a));
    }
    if (addr && addr !== "" && addr !== ":address") {
      setAddress(addr as unknown as base.Address);
    }
  }, [addr]);

  const table = useReactTable({
    data: history.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <FormTable data={history} definition={createHistoryForm(table)} />
    </View>
  );
}

type theInstance = InstanceType<typeof types.TransactionContainer>;
function createHistoryForm(table: any): GroupDefinition<theInstance>[] {
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
