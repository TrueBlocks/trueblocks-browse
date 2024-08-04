import React from "react";
import { types } from "@gocode/models";
import { GroupDefinition, DataTable } from "@components";

export type AbiInstance = InstanceType<typeof types.AbiSummary>;

export function createAbiForm(table: any): GroupDefinition<AbiInstance>[] {
  return [
    {
      title: "Abi Data",
      colSpan: 6,
      fields: [
        { label: "nFunctions", type: "int", accessor: "nFunctions" },
        { label: "nEvents", type: "int", accessor: "nEvents" },
        { label: "fileSize", type: "bytes", accessor: "fileSize" },
      ],
    },
    {
      title: "Future Use",
      colSpan: 6,
      fields: [],
    },
    {
      title: "Files",
      fields: [],
      customComponents: [
        {
          component: <DataTable<types.ChunkRecord> table={table} loading={false} />,
        },
      ],
    },
  ];
}
