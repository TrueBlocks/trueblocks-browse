import React from "react";
import { types } from "@gocode/models";
import { GroupDefinition, DataTable } from "@components";
import { Pagination } from "@mantine/core";

export type theInstance = InstanceType<typeof types.SummaryTransaction>;

export function createForm(
  table: any,
  firstRecord: number,
  totalRecords: number,
  perPage: number
): GroupDefinition<theInstance>[] {
  const pageNumber = firstRecord < perPage ? 1 : Math.ceil(firstRecord / perPage) + 1;
  const totalPages = Math.ceil(totalRecords / perPage);

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
      components: [
        {
          component: (
            <>
              <DataTable<types.Transaction> table={table} loading={false} />
              <div style={{ display: "flex", justifyContent: "flex-end", marginTop: "1rem" }}>
                <Pagination size="sm" value={pageNumber} total={totalPages} />
              </div>
            </>
          ),
        },
      ],
    },
  ];
}
