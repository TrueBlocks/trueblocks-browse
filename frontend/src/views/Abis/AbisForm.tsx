import React from "react";
import { types } from "@gocode/models";
import { GroupDefinition, DataTable, Pager } from "@components";

type theInstance = InstanceType<typeof types.AbiContainer>;

export function createForm(table: any, pager: Pager): GroupDefinition<theInstance>[] {
  return [
    {
      title: "Abi Data",
      colSpan: 6,
      fields: [
        { label: "nItems", type: "int", accessor: "nItems" },
        { label: "nFunctions", type: "int", accessor: "nFunctions" },
        { label: "nEvents", type: "int", accessor: "nEvents" },
        { label: "fileSize", type: "bytes", accessor: "fileSize" },
      ],
    },
    {
      title: "Bounds",
      colSpan: 6,
      fields: [
        { label: "largestFile", type: "text", accessor: "largestFile" },
        { label: "mostFunctions", type: "text", accessor: "mostFunctions" },
        { label: "mostEvents", type: "text", accessor: "mostEvents" },
      ],
    },
    {
      title: "Files",
      fields: [],
      components: [
        {
          component: <DataTable<types.Abi> table={table} loading={false} pager={pager} />,
        },
      ],
    },
  ];
}
