import React, { useState, useEffect } from "react";
import { types } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns, createForm } from ".";
import { View2, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetAbis, GetAbisCnt } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";

export function AbisView() {
  const [summaryItem, setSummaryItem] = useState<types.AbiContainer>({} as types.AbiContainer);
  const [count, setCount] = useState<number>(0);
  const pager = useKeyboardPaging(count, [], 15);

  useEffect(() => {
    const fetch = async (currentItem: number, itemsPerPage: number) => {
      GetAbis(currentItem, itemsPerPage).then((item: types.AbiContainer) => {
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
