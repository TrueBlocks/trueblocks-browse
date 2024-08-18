import React from "react";
import { types } from "@gocode/models";
import { GroupDefinition, DataTable, Pager } from "@components";

type theInstance = InstanceType<typeof types.MonitorContainer>;

export function createForm(table: any, pager: Pager): GroupDefinition<theInstance>[] {
  return [
    {
      title: "Monitor Data",
      colSpan: 6,
      fields: [
        { label: "nItems", type: "int", accessor: "nItems" },
        { label: "nRecords", type: "int", accessor: "nRecords" },
        { label: "nNamed", type: "int", accessor: "nNamed" },
      ],
    },
    {
      title: "Other",
      colSpan: 6,
      fields: [
        { label: "fileSize", type: "bytes", accessor: "fileSize" },
        { label: "nDeleted", type: "int", accessor: "nDeleted" },
      ],
    },
    {
      title: "Files",
      fields: [],
      components: [
        {
          component: <DataTable<types.Monitor> table={table} loading={false} pager={pager} />,
        },
      ],
    },
  ];
}
