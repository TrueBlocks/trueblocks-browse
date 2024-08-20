import React, { useEffect } from "react";
import { useParams } from "wouter";
import { types, base } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns } from "./HistoryTable";
import { View, FormTable, DataTable, GroupDefinition } from "@components";
import { GetLastSub } from "@gocode/app/App";
import { useAppState, ViewStateProvider } from "@state";

export function HistoryView() {
  const { setAddress, history, resetPager } = useAppState();

  var aa = useParams().address;
  useEffect(() => {
    // TODO: This doesn't compile. It uses useKeyboardPaging but can't outside of a
    // TODO: component. It's needed when the address changes to go back to page zero.
    // resetPager("history");
    if (aa === ":address") {
      GetLastSub("/history").then((subRoute) => {
        subRoute = subRoute.replace("/", "");
        setAddress(subRoute as unknown as base.Address);
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
    <ViewStateProvider route="history" nItems={history.nItems}>
      <View>
        <FormTable data={history} definition={createHistoryForm(table)} />
      </View>
    </ViewStateProvider>
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
          component: <DataTable<types.Transaction> table={table} loading={false} />,
        },
      ],
    },
  ];
}
