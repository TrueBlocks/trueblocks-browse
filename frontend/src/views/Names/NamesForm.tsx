import React from "react";
import { types } from "@gocode/models";
import { GroupDefinition, DataTable } from "@components";

export type NameInstance = InstanceType<typeof types.SummaryName>;

export function createNameForm(table: any): GroupDefinition<NameInstance>[] {
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
          component: <DataTable<types.Name> table={table} loading={false} />,
        },
      ],
    },
  ];
}
