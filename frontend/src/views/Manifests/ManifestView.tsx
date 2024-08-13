import React, { useState, useEffect } from "react";
import { types } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns, createForm } from ".";
import { View2, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { GetManifest, GetManifestCnt } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";

export function ManifestView() {
  const [loading, setLoading] = useState<boolean>(false);
  const [loaded, setLoaded] = useState<boolean>(false);
  const [summaryItem, setSummaryItem] = useState<types.ManifestContainer>({} as types.ManifestContainer);
  const [refresh, setRefresh] = useState<boolean>(false);
  const [count, setCount] = useState<number>(0);
  const pager = useKeyboardPaging(count, [], 15);

  useEffect(() => {
    if (loaded && !loading) {
      const fetch = async (currentItem: number, itemsPerPage: number) => {
        GetManifest(currentItem, itemsPerPage).then((item: types.ManifestContainer) => {
          setSummaryItem(item);
        });
      };
      fetch(pager.curItem, pager.perPage);
      setRefresh(false);
    }
  }, [count, pager.curItem, pager.perPage, loaded, loading]);

  useEffect(() => {
    const handleRefresh = () => {
      setRefresh(true);
    };

    EventsOn("DAEMON", handleRefresh);
    return () => {
      EventsOff("DAEMON");
    };
  }, []);

  useEffect(() => {
    setLoading(true);
    const fetch = async () => {
      const cnt = await GetManifestCnt();
      setCount(cnt);
      setLoaded(true);
    };
    fetch().finally(() => setLoading(false));
  }, []);

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
