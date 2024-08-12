import React from "react";
import { types } from "@gocode/models";
import { GroupDefinition, DataTable, Pager } from "@components";

export type theInstance = InstanceType<typeof types.SummaryIndex>;

export function createForm(table: any, pager: Pager): GroupDefinition<theInstance>[] {
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
              <DataTable<types.ChunkStats> table={table} loading={false} pager={pager} />
            </>
          ),
        },
      ],
    },
  ];
}
