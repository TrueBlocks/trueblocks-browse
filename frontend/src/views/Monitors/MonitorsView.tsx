import React, { useState, useEffect } from "react";
import { types, messages } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns } from "./MonitorsTable";
import { View, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { MonitorPage, GetMonitorsCnt } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";
import { GroupDefinition, DataTable, Pager } from "@components";

export function MonitorsView() {
  const [summaryItem, setSummaryItem] = useState<types.MonitorContainer>({} as types.MonitorContainer);
  const [count, setCount] = useState<number>(0);
  const pager = useKeyboardPaging("monitors", count, [], 15);

  useEffect(() => {
    const fetch = async (currentItem: number, itemsPerPage: number) => {
      MonitorPage(currentItem, itemsPerPage).then((item: types.MonitorContainer) => {
        if (item) {
          GetMonitorsCnt().then((cnt: number) => setCount(cnt));
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
      <FormTable data={summaryItem} definition={createMonitorForm(table, pager)} />
    </View>
  );
}

type theInstance = InstanceType<typeof types.MonitorContainer>;
function createMonitorForm(table: any, pager: Pager): GroupDefinition<theInstance>[] {
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
          component: <DataTable<types.Monitor> table={table} loading={false} pagerName="monitors" />,
        },
      ],
    },
  ];
}
