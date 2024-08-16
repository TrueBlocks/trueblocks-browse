import React, { useState, useEffect } from "react";
import { types, messages } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns, createForm } from ".";
import { View, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetNames } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";
import { useAppState } from "@state";

export function NamesView() {
  const { names, setNames } = useAppState();
  const pager = useKeyboardPaging(names.nItems, [], 15);

  useEffect(() => {
    const fetch = async (currentItem: number, itemsPerPage: number) => {
      GetNames(currentItem, itemsPerPage).then((item: types.NameContainer) => {
        if (item) {
          setNames(item);
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
    data: names.names || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <FormTable data={names} definition={createForm(table, pager)} />
    </View>
  );
}
