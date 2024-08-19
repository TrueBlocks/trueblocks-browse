import React, { useState, useEffect } from "react";
import { types, messages } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns } from "./AbisTable";
import { View, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { AbiPage, GetAbisCnt } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";
import { GroupDefinition, DataTable, Pager } from "@components";

export function AbisView() {
  const [summaryItem, setSummaryItem] = useState<types.AbiContainer>({} as types.AbiContainer);
  const [count, setCount] = useState<number>(0);
  const pager = useKeyboardPaging("abis", count, [], 15);

  useEffect(() => {
    const fetch = async (currentItem: number, itemsPerPage: number) => {
      AbiPage(currentItem, itemsPerPage).then((item: types.AbiContainer) => {
        if (item) {
          GetAbisCnt().then((cnt: number) => setCount(cnt));
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
    data: summaryItem.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <FormTable data={summaryItem} definition={createAbisForm(table, pager)} />
    </View>
  );
}

type theInstance = InstanceType<typeof types.AbiContainer>;
function createAbisForm(table: any, pager: Pager): GroupDefinition<theInstance>[] {
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
