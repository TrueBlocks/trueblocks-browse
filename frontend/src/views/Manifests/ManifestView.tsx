import React, { useState, useEffect } from "react";
import { types, messages } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns, createForm } from ".";
import { View, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { ManifestPage, GetManifestCnt } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";

export function ManifestView() {
  const [summaryItem, setSummaryItem] = useState<types.ManifestContainer>({} as types.ManifestContainer);
  const [count, setCount] = useState<number>(0);
  const pager = useKeyboardPaging("manifest", count, [], 15);

  useEffect(() => {
    const fetch = async (currentItem: number, itemsPerPage: number) => {
      ManifestPage(currentItem, itemsPerPage).then((item: types.ManifestContainer) => {
        if (item) {
          GetManifestCnt().then((cnt: number) => setCount(cnt));
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
