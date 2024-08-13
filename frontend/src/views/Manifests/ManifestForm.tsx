import React from "react";
import { types } from "@gocode/models";
import { GroupDefinition, DataTable, Pager } from "@components";

type theInstance = InstanceType<typeof types.ManifestContainer>;

export function createForm(table: any, pager: Pager): GroupDefinition<theInstance>[] {
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
      components: [
        {
          component: <DataTable<types.ChunkRecord> table={table} loading={false} pager={pager} />,
        },
      ],
    },
  ];
}
