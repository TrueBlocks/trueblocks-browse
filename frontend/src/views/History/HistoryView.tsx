// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { useMemo } from "react";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm, DebugState, UnderConstruction } from "@components";
import { useNoops, useUtils } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { HistoryTableDef, HistoryFormDef } from ".";
// EXISTING_CODE

export const HistoryView = () => {
  const { history, fetchHistory } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  // EXISTING_CODE
  const { info } = useAppState();
  const { ShortenAddr } = useUtils();
  const addrStr = useMemo(
    () => (info?.address ? ShortenAddr(info.address.toString()) : ""),
    [ShortenAddr, info?.address]
  );
  // EXISTING_CODE

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
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={history.nItems}
      fetchFn={fetchHistory}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState n={history.lastUpdate} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE

//-------------------------------------------------------------------
// Template variables:
// class:         History
// lower:         history
// routeLabel:    History
// routeLower:    history
