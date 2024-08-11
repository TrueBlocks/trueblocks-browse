import React from "react";
import { types } from "@gocode/models";
import { GroupDefinition, DataTable } from "@components";
import { Pagination } from "@mantine/core";

export type theInstance = InstanceType<typeof types.SummaryName>;

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
      title: "Name Data",
      colSpan: 6,
      fields: [
        { label: "nNames", type: "int", accessor: "nNames" },
        { label: "nContracts", type: "int", accessor: "nContracts" },
        { label: "nErc20s", type: "int", accessor: "nErc20s" },
        { label: "nErc721s", type: "int", accessor: "nErc721s" },
        { label: "nDeleted", type: "int", accessor: "nDeleted" },
      ],
    },
    {
      title: "Database Parts",
      colSpan: 6,
      fields: [
        { label: "nCustom", type: "int", accessor: "nCustom" },
        { label: "nRegular", type: "int", accessor: "nRegular" },
        { label: "nPrefund", type: "int", accessor: "nPrefund" },
        { label: "nBaddress", type: "int", accessor: "nBaddress" },
      ],
    },
    {
      title: "Names",
      fields: [],
      components: [
        {
          component: (
            <>
              <DataTable<types.Name> table={table} loading={false} />
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
