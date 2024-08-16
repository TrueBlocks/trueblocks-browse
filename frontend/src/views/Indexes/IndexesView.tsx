import React, { useState, useEffect } from "react";
import { types, messages } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns, createForm } from ".";
import { View, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetIndex } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";
import { useAppState } from "@state";

export function IndexesView() {
  const { indexes, setIndexes } = useAppState();
  const pager = useKeyboardPaging(indexes.nItems, [], 15);

  useEffect(() => {
    const fetch = async (currentItem: number, itemsPerPage: number) => {
      GetIndex(currentItem, itemsPerPage).then((item: types.IndexContainer) => {
        if (item) {
          setIndexes(item);
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
    data: indexes.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <FormTable data={indexes} definition={createForm(table, pager)} />
    </View>
  );
}
