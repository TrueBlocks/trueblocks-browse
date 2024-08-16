import React, { useState, useEffect } from "react";
import { types, messages } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns, createForm } from ".";
import { View, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetManifest } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";
import { useAppState } from "@state";

export function ManifestView() {
  const { manifests, setManifests } = useAppState();
  const pager = useKeyboardPaging(manifests.nItems, [], 15);

  useEffect(() => {
    const fetch = async (currentItem: number, itemsPerPage: number) => {
      GetManifest(currentItem, itemsPerPage).then((item: types.ManifestContainer) => {
        if (item) {
          setManifests(item);
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
    data: manifests.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <View>
      <FormTable data={manifests} definition={createForm(table, pager)} />
    </View>
  );
}
