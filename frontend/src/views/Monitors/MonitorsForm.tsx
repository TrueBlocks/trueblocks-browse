import React from "react";
import { types } from "@gocode/models";
import { GroupDefinition, DataTable } from "@components";
import { Pagination } from "@mantine/core";

export type theInstance = InstanceType<typeof types.SummaryMonitor>;

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
              <DataTable<types.Monitor> table={table} loading={false} />
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
