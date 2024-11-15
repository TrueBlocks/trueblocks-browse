import { useEffect } from "react";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm, DebugState } from "@components";
import { messages } from "@gocode/models";
import { useNoops, useUtils } from "@hooks";
import { EventsOff, EventsOn } from "@runtime";
import { useAppState, useViewState, ViewStateProvider } from "@state";
import { HistoryTableDef, HistoryFormDef } from ".";

export const HistoryView = () => {
  const { modifyNoop } = useNoops();
  const { ShortenAddr } = useUtils();
  const { info, history, fetchHistory } = useAppState();

  const table = useReactTable({
    data: history.items || [],
    columns: HistoryTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "history";
  const addrStr = ShortenAddr(info.address.toString());
  const tabs = [addrStr];
  const forms: ViewForm = {
    [addrStr]: <FormTable data={history} groups={HistoryFormDef(info.address, table)} />,
  };

  return (
    <ViewStateProvider route={route} nItems={history.nItems} fetchFn={fetchHistory} modifyFn={modifyNoop}>
      <DebugState n={history.lastUpdate} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};
