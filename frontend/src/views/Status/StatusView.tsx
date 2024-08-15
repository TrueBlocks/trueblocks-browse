import React, { useState, useEffect } from "react";
import { types } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns, createForm } from ".";
import { View2, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetStatus, GetStatusCnt } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";

export function StatusView() {
  const [summaryItem, setSummaryItem] = useState<types.StatusContainer>({} as types.StatusContainer);
  const [count, setCount] = useState<number>(0);
  const pager = useKeyboardPaging(count, [], 10);

  useEffect(() => {
    const fetch = async (currentItem: number, itemsPerPage: number) => {
      GetStatus(currentItem, itemsPerPage).then((item: types.StatusContainer) => {
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

    EventsOn("DAEMON", handleRefresh);
    return () => {
      EventsOff("DAEMON");
    };
  }, [pager.curItem, pager.perPage]);

  const table = useReactTable({
    data: summaryItem.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View2>
      <FormTable data={summaryItem} definition={createForm(table, pager)} />
    </View2>
  );
}
