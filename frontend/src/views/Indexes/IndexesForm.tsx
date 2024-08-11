import React from "react";
import { types } from "@gocode/models";
import { GroupDefinition, DataTable } from "@components";
import { Pagination } from "@mantine/core";

export type theInstance = InstanceType<typeof types.SummaryIndex>;

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
      title: "Summary Data",
      colSpan: 6,
      fields: [
        { label: "bloomSz", type: "bytes", accessor: "bloomSz" },
        { label: "chunkSz", type: "bytes", accessor: "chunkSz" },
        { label: "nAddrs", type: "int", accessor: "nAddrs" },
        { label: "nApps", type: "int", accessor: "nApps" },
        { label: "nBlocks", type: "int", accessor: "nBlocks" },
        { label: "nBlooms", type: "int", accessor: "nBlooms" },
      ],
    },
    {
      title: "Statistics",
      colSpan: 6,
      fields: [
        { label: "addrsPerBlock", type: "float", accessor: "addrsPerBlock" },
        { label: "appsPerAddr", type: "float", accessor: "appsPerAddr" },
        { label: "appsPerBlock", type: "float", accessor: "appsPerBlock" },
      ],
    },
    {
      title: "Chunks",
      fields: [],
      components: [
        {
          component: (
            <>
              <DataTable<types.ChunkStats> table={table} loading={false} />
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
