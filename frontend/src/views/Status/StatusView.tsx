import React, { useState, useEffect } from "react";
import { types, messages } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns, createForm } from ".";
import { View, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { StatusPage, GetStatusCnt } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";

export function StatusView() {
  const [summaryItem, setSummaryItem] = useState<types.StatusContainer>({} as types.StatusContainer);
  const [count, setCount] = useState<number>(0);
  const pager = useKeyboardPaging(count, [], 10);

  useEffect(() => {
    const fetch = async (currentItem: number, itemsPerPage: number) => {
      StatusPage(currentItem, itemsPerPage).then((item: types.StatusContainer) => {
        if (item) {
          GetStatusCnt().then((cnt: number) => setCount(cnt));
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
      <FormTable data={summaryItem} definition={createForm(table, pager)} />
    </View>
  );
}
