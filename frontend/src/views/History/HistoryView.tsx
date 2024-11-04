import { useEffect } from "react";
import { Text } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { useParams } from "wouter";
import { View, FormTable, ViewForm } from "@components";
import { base } from "@gocode/models";
import { useNoops, useRenderCounter, useUtils } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { HistoryTableDef, HistoryFormDef } from ".";

export const HistoryView = () => {
  const { modifyNoop } = useNoops();
  const { ShortenAddr } = useUtils();
  const { setAddress, history, fetchHistory } = useAppState();
  const renderCount = useRenderCounter();

  const address = useParams().address as unknown as base.Address;
  useEffect(() => {
    setAddress(address);
  }, [address, setAddress]);

  const table = useReactTable({
    data: history.items || [],
    columns: HistoryTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "history";
  const addrStr = ShortenAddr(address.toString());
  const tabs = [addrStr];
  const forms: ViewForm = {
    [addrStr]: <FormTable data={history} groups={HistoryFormDef(address, table)} />,
  };
  return (
    <ViewStateProvider route={route} nItems={history.nItems} fetchFn={fetchHistory} modifyFn={modifyNoop}>
      <Text>Render count: {renderCount}</Text>
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};
