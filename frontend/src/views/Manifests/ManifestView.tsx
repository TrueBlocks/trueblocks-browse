import React, { useState, useEffect } from "react";
import { types, messages } from "@gocode/models";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { tableColumns } from "./ManifestTable";
import { View, FormTable } from "@components";
import { useKeyboardPaging } from "@hooks";
import { ManifestPage, GetManifestCnt } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";
import { GroupDefinition, DataTable, Pager } from "@components";

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
      <FormTable data={summaryItem} definition={createManifestForm(table)} />
    </View>
  );
}

type theInstance = InstanceType<typeof types.ManifestContainer>;
function createManifestForm(table: any): GroupDefinition<theInstance>[] {
  return [
    {
      title: "Manifest Data",
      colSpan: 6,
      fields: [
        { label: "version", type: "text", accessor: "version" },
        { label: "chain", type: "text", accessor: "chain" },
        { label: "specification", type: "hash", accessor: "specification" },
        { label: "latestUpdate", type: "date", accessor: "latestUpdate" },
      ],
    },
    {
      title: "Statistics",
      colSpan: 6,
      fields: [
        { label: "nBlooms", type: "int", accessor: "nBlooms" },
        { label: "bloomsSize", type: "bytes", accessor: "bloomsSize" },
        { label: "nIndexes", type: "int", accessor: "nIndexes" },
        { label: "indexSize", type: "bytes", accessor: "indexSize" },
      ],
    },
    {
      title: "Chunks",
      fields: [],
      components: [
        {
          component: <DataTable<types.ChunkRecord> table={table} loading={false} pagerName="manifest" />,
        },
      ],
    },
  ];
}
