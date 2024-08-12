import React from "react";
import { types } from "@gocode/models";
import { GroupDefinition, DataTable, Pager } from "@components";

export type theInstance = InstanceType<typeof types.SummaryMonitor>;

export function createForm(table: any, pager: Pager): GroupDefinition<theInstance>[] {
  return [
    {
      title: "Monitor Data",
      colSpan: 6,
      fields: [
        { label: "nMonitors", type: "int", accessor: "nMonitors" },
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
          component: (
            <>
              <DataTable<types.Monitor> table={table} loading={false} pager={pager} />
            </>
          ),
        },
      ],
    },
  ];
}
