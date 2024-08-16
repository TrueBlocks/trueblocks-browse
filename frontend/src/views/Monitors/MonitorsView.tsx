import React, { useState, useEffect } from "react";
import { types, messages } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns, createForm } from ".";
import { View, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetMonitors } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";
import { useAppState } from "@state";

export function MonitorsView() {
  const { monitors, setMonitors } = useAppState();
  const pager = useKeyboardPaging(monitors.nItems, [], 15);

  useEffect(() => {
    const fetch = async (currentItem: number, itemsPerPage: number) => {
      GetMonitors(currentItem, itemsPerPage).then((item: types.MonitorContainer) => {
        if (item) {
          setMonitors(item);
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
    data: monitors.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <FormTable data={monitors} definition={createForm(table, pager)} />
    </View>
  );
}
