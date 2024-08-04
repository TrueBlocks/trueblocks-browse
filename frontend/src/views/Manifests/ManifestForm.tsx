import React from "react";
import { types } from "@gocode/models";
import { GroupDefinition, DataTable } from "@components";

export type ManifestInstance = InstanceType<typeof types.ManifestSummary>;

export function createManifestForm(table: any): GroupDefinition<ManifestInstance>[] {
  return [
    {
      title: "Manifest Data",
      colSpan: 6,
      fields: [
        { label: "version", accessor: "version" },
        { label: "chain", accessor: "chain" },
        { label: "specification", accessor: "specification" },
        { label: "latestUpdate", accessor: "latestUpdate" },
      ],
    },
    {
      title: "Statistics",
      colSpan: 6,
      fields: [
        { label: "nBlooms", type: "int", accessor: "nBlooms" },
        { label: "bloomsSize", type: "bytes", accessor: "bloomsSize" },
        { label: "nIndexes", type: "int", accessor: "nIndexes" },
        { label: "indexSize", type: "bytes", accessor: "indexSize" },
      ],
    },
    {
      title: "Chunks",
      fields: [],
      customComponents: [
        {
          component: <DataTable<types.ChunkRecord> table={table} loading={false} />,
        },
      ],
    },
  ];
}
