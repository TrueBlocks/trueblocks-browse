import React, { useState, useEffect } from "react";
import { types, messages } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns } from "./NamesTable";
import { View, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { NamePage, GetNamesCnt } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";
import { GroupDefinition, DataTable, Pager } from "@components";

export function NamesView() {
  const [summaryItem, setSummaryItem] = useState<types.NameContainer>({} as types.NameContainer);
  const [count, setCount] = useState<number>(0);
  const pager = useKeyboardPaging("names", count, [], 15);

  useEffect(() => {
    const fetch = async (currentItem: number, itemsPerPage: number) => {
      NamePage(currentItem, itemsPerPage).then((item: types.NameContainer) => {
        if (item) {
          GetNamesCnt().then((cnt: number) => setCount(cnt));
          setSummaryItem(item);
        }
      });
    };
    fetch(pager.curItem, pager.perPage);

    const handleRefresh = () => {
      fetch(pager.curItem, pager.perPage);
    };

    var { Message } = messages;
    EventsOn(Message.DAEMON, handleRefresh);
    return () => {
      EventsOff(Message.DAEMON);
    };
  }, [pager.curItem, pager.perPage]);

  const table = useReactTable({
    data: summaryItem.names || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <FormTable data={summaryItem} definition={createNameForm(table, pager)} />
    </View>
  );
}

type theInstance = InstanceType<typeof types.NameContainer>;
function createNameForm(table: any, pager: Pager): GroupDefinition<theInstance>[] {
  return [
    {
      title: "Name Data",
      colSpan: 6,
      fields: [
        { label: "nItems", type: "int", accessor: "nItems" },
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
          component: <DataTable<types.Name> table={table} loading={false} pagerName="names" />,
        },
      ],
    },
  ];
}
