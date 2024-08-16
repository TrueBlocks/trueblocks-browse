import React, { useState, useEffect } from "react";
import { types, messages } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns, createForm } from ".";
import { View, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetAbis } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";
import { useAppState } from "@state";

export function AbisView() {
  const { abis, setAbis } = useAppState();
  const pager = useKeyboardPaging(abis.nItems, [], 15);

  useEffect(() => {
    const fetch = async (currentItem: number, itemsPerPage: number) => {
      GetAbis(currentItem, itemsPerPage).then((item: types.AbiContainer) => {
        if (item) {
          setAbis(item);
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
    data: abis.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <FormTable data={abis} definition={createForm(table, pager)} />
    </View>
  );
}
