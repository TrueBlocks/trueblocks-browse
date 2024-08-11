import React from "react";
import { types } from "@gocode/models";
import { GroupDefinition, DataTable, Paginator } from "@components";

export type theInstance = InstanceType<typeof types.SummaryAbis>;

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
      title: "Abi Data",
      colSpan: 6,
      fields: [
        { label: "nAbis", type: "int", accessor: "nAbis" },
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
          component: (
            <>
              <DataTable<types.Abi> table={table} loading={false} />
              <Paginator pageNumber={pageNumber} totalPages={totalPages} />
            </>
          ),
        },
      ],
    },
  ];
}
