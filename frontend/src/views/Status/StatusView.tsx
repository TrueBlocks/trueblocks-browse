import React, { useState, useEffect } from "react";
import { types, messages } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns, createForm } from ".";
import { View2, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetStatus } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";
import { useAppState } from "@state";

export function StatusView() {
  const { status, setStatus } = useAppState();
  const [count, setCount] = useState<number>(0);
  const pager = useKeyboardPaging(count, [], 10);

  useEffect(() => {
    const fetch = async (currentItem: number, itemsPerPage: number) => {
      GetStatus(currentItem, itemsPerPage).then((item: types.StatusContainer) => {
        if (item) {
          setCount(item.nItems);
          setStatus(item);
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
    data: status.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View2>
      <FormTable data={status} definition={createForm(table, pager)} />
    </View2>
  );
}
