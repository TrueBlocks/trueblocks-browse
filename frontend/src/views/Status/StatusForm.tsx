import React from "react";
import { types } from "@gocode/models";
import { GroupDefinition, DataTable } from "@components";

export type StatusInstance = InstanceType<typeof types.StatusEx>;

export function createStatusForm(table: any): GroupDefinition<StatusInstance>[] {
  return [
    {
      title: "System Data",
      colSpan: 7,
      fields: [
        { label: "trueblocks", accessor: "version" },
        { label: "client", accessor: "clientVersion" },
        { label: "isArchive", accessor: "isArchive" },
        { label: "isTracing", accessor: "isTracing" },
      ],
    },
    {
      title: "API Keys",
      colSpan: 5,
      fields: [
        { label: "hasEsKey", accessor: "hasEsKey" },
        { label: "hasPinKey", accessor: "hasPinKey" },
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
        { label: "nFolders", accessor: "nFolders" },
        { label: "nFiles", accessor: "nFiles" },
        { label: "sizeInBytes", accessor: "nBytes" },
      ],
    },
    {
      title: "Caches",
      fields: [],
      customComponents: [
        {
          component: <DataTable<types.CacheItem> table={table} loading={false} />,
        },
      ],
    },
  ];
}
