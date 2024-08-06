import React from "react";
import { types } from "@gocode/models";
import { GroupDefinition, DataTable } from "@components";

export type TransactionInstance = InstanceType<typeof types.SummaryTransaction>;

export function createTransactionForm(table: any): GroupDefinition<TransactionInstance>[] {
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
        { label: "nTransactions", type: "int", accessor: "nTransactions" },
        { label: "nEvents", type: "int", accessor: "nEvents" },
        { label: "nTokens", type: "int", accessor: "nTokens" },
        { label: "nErrors", type: "int", accessor: "nErrors" },
      ],
    },
    {
      title: "Files",
      fields: [],
      customComponents: [
        {
          component: <DataTable<types.Transaction> table={table} loading={false} />,
        },
      ],
    },
  ];
}
