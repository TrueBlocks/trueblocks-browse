import React from "react";
import { types } from "@gocode/models";
import { GroupDefinition, DataTable } from "@components";
import { Pagination } from "@mantine/core";

export type theInstance = InstanceType<typeof types.SummaryStatus>;

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
      title: "System Data",
      colSpan: 7,
      fields: [
        { label: "trueblocks", accessor: "version" },
        { label: "client", accessor: "clientVersion" },
        { label: "isArchive", type: "boolean", accessor: "isArchive" },
        { label: "isTracing", type: "boolean", accessor: "isTracing" },
      ],
    },
    {
      title: "API Keys",
      colSpan: 5,
      fields: [
        { label: "hasEsKey", type: "boolean", accessor: "hasEsKey" },
        { label: "hasPinKey", type: "boolean", accessor: "hasPinKey" },
        { label: "rpcProvider", accessor: "rpcProvider" },
      ],
    },
    {
      title: "Configuration Paths",
      colSpan: 7,
      fields: [
        { label: "rootConfig", accessor: "rootConfig" },
        { label: "chainConfig", accessor: "chainConfig" },
        { label: "indexPath", accessor: "indexPath" },
        { label: "cachePath", accessor: "cachePath" },
      ],
    },
    {
      title: "Statistics",
      colSpan: 5,
      fields: [
        { label: "latestCached", accessor: "latestUpdate" },
        { label: "nFiles", type: "int", accessor: "nFiles" },
        { label: "nFolders", type: "int", accessor: "nFolders" },
        { label: "sizeInBytes", type: "bytes", accessor: "nBytes" },
      ],
    },
    {
      title: "Caches",
      fields: [],
      components: [
        {
          component: (
            <>
              <DataTable<types.CacheItem> table={table} loading={false} />
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
