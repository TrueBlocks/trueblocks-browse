import React from "react";
import { types } from "@gocode/models";
import { createColumnHelper, ColumnDef } from "@tanstack/react-table";
import { CustomMeta, GroupDefinition, DataTable } from "@components";

export type ManifestInstance = InstanceType<typeof types.ManifestEx>;

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
        { label: "nBlooms", accessor: "nBlooms" },
        { label: "bloomsSize", accessor: "bloomsSize" },
        { label: "nIndexes", accessor: "nIndexes" },
        { label: "indexSize", accessor: "indexSize" },
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
