import React, { useEffect } from "react";
import { useParams } from "wouter";
import { types, base } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns } from "./HistoryTable";
import { View, FormTable, DataTable, GroupDefinition } from "@components";
import { GetLastSub } from "@gocode/app/App";
import { useAppState } from "@state";

export function HistoryView() {
  const { address, setAddress, history } = useAppState();

  var aa = useParams().address;
  useEffect(() => {
    if (aa === ":address") {
      GetLastSub("/history").then((subRoute) => {
        // console.log("subRoute: ", subRoute);
        // console.log("as address: ", subRoute as unknown as base.Address);
        subRoute = subRoute.replace("/", "");
        return setAddress(subRoute as unknown as base.Address);
      });
    } else {
      setAddress(aa as unknown as base.Address);
    }
  }, [aa]);

  const table = useReactTable({
    data: history.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View route="history" nItems={history.nItems}>
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
        { label: "address", type: "address-address-only", accessor: "address" },
        { label: "name", type: "address-name-only", accessor: "name" },
        { label: "balance", type: "ether", accessor: "balance" },
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
