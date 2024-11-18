import { useMemo } from "react";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm, DebugState, UnderConstruction } from "@components";
import { useNoops, useUtils } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { HistoryTableDef, HistoryFormDef } from ".";

export const HistoryView = () => {
  const { modifyNoop } = useNoops();
  const { ShortenAddr } = useUtils();
  const { info, history, fetchHistory } = useAppState();

  const addrStr = useMemo(
    () => (info?.address ? ShortenAddr(info.address.toString()) : ""),
    [ShortenAddr, info?.address]
  );

  const table = useReactTable({
    data: history.items || [],
    columns: HistoryTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "history";
  const tabs = [addrStr, "logs", "statements", "neighbors", "charts", "contracts", "erc20", "erc721", "poaps"];
  const forms: ViewForm = {
    [addrStr]: <FormTable data={history} groups={HistoryFormDef(table, info.address)} />,
    logs: <UnderConstruction message="Coming soon...a list of all logs" />,
    statements: <UnderConstruction message="Coming soon...accounting statements" />,
    neighbors: <UnderConstruction message="Coming soon...all addresses you've interacted with" />,
    charts: <UnderConstruction message="Coming soon...balance history charts, activity charts, etc." />,
    contracts: <UnderConstruction message="Coming soon...all smart contracts you've interacted with" />,
    erc20: <UnderConstruction message="Coming soon...all Tokens you own" />,
    erc721: <UnderConstruction message="Coming soon...all NFTs you own" />,
    poaps: <UnderConstruction message="Coming soon...all POAPs you own" />,
  };

  return (
    <ViewStateProvider route={route} nItems={history.nItems} fetchFn={fetchHistory} modifyFn={modifyNoop}>
      <DebugState n={history.lastUpdate} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};
