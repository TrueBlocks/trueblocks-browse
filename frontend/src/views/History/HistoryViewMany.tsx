// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { useCallback } from "react";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { useNoops } from "@hooks";
import { DebugState, TabItem, View, ViewForm } from "../../components";
import { useAppState, ViewStateProvider } from "../../state";
import { BalancesFormDef, BalancesTableDef } from "../Balances";
import { ChartsFormDef, ChartsTableDef } from "../Charts";
import { IncomingFormDef, IncomingTableDef } from "../Incoming";
import { InternalsFormDef, InternalsTableDef } from "../Internals";
import { LogsFormDef, LogsTableDef } from "../Logs";
import { NeighborsFormDef, NeighborsTableDef } from "../Neighbors";
import { OutgoingFormDef, OutgoingTableDef } from "../Outgoing";
import { ReceiptsFormDef, ReceiptsTableDef } from "../Receipts";
import { StatementsFormDef, StatementsTableDef } from "../Statements";
import { TracesFormDef, TracesTableDef } from "../Traces";

// EXISTING_CODE

export const HistoryView = () => {
  const {
    balances,
    fetchBalances,
    incoming,
    fetchIncoming,
    outgoing,
    fetchOutgoing,
    internals,
    fetchInternals,
    charts,
    fetchCharts,
    logs,
    fetchLogs,
    statements,
    fetchStatements,
    neighbors,
    fetchNeighbors,
    traces,
    fetchTraces,
    receipts,
    fetchReceipts,
  } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  const fetchHistory = useCallback(
    (currentItem: number, itemsPerPage: number) => {
      fetchBalances(currentItem, itemsPerPage);
      fetchIncoming(currentItem, itemsPerPage);
      fetchOutgoing(currentItem, itemsPerPage);
      fetchInternals(currentItem, itemsPerPage);
      fetchCharts(currentItem, itemsPerPage);
      fetchLogs(currentItem, itemsPerPage);
      fetchStatements(currentItem, itemsPerPage);
      fetchNeighbors(currentItem, itemsPerPage);
      fetchTraces(currentItem, itemsPerPage);
      fetchReceipts(currentItem, itemsPerPage);
    },
    [
      fetchBalances,
      fetchIncoming,
      fetchOutgoing,
      fetchInternals,
      fetchCharts,
      fetchLogs,
      fetchStatements,
      fetchNeighbors,
      fetchTraces,
      fetchReceipts,
    ]
  );

  const balancesTable = useReactTable({
    data: balances?.items || [],
    columns: BalancesTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const incomingTable = useReactTable({
    data: incoming?.items || [],
    columns: IncomingTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const outgoingTable = useReactTable({
    data: outgoing?.items || [],
    columns: OutgoingTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const internalsTable = useReactTable({
    data: internals?.items || [],
    columns: InternalsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const chartsTable = useReactTable({
    data: charts?.items || [],
    columns: ChartsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const logsTable = useReactTable({
    data: logs?.items || [],
    columns: LogsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const statementsTable = useReactTable({
    data: statements?.items || [],
    columns: StatementsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const neighborsTable = useReactTable({
    data: neighbors?.items || [],
    columns: NeighborsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const tracesTable = useReactTable({
    data: traces?.items || [],
    columns: TracesTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const receiptsTable = useReactTable({
    data: receipts?.items || [],
    columns: ReceiptsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const tabItems: ViewForm = {
    balances: <TabItem data={balances} groups={BalancesFormDef(balancesTable)} />,
    incoming: <TabItem data={incoming} groups={IncomingFormDef(incomingTable)} />,
    outgoing: <TabItem data={outgoing} groups={OutgoingFormDef(outgoingTable)} />,
    internals: <TabItem data={internals} groups={InternalsFormDef(internalsTable)} />,
    charts: <TabItem data={charts} groups={ChartsFormDef(chartsTable)} />,
    logs: <TabItem data={logs} groups={LogsFormDef(logsTable)} />,
    statements: <TabItem data={statements} groups={StatementsFormDef(statementsTable)} />,
    neighbors: <TabItem data={neighbors} groups={NeighborsFormDef(neighborsTable)} />,
    traces: <TabItem data={traces} groups={TracesFormDef(tracesTable)} />,
    receipts: <TabItem data={receipts} groups={ReceiptsFormDef(receiptsTable)} />,
  };

  // if (!(status?.items?.length > 0)) {
  //   return <>Loading...</>;
  // }

  return (
    <ViewStateProvider
      // do not remove - delint
      nItems={balances.nItems}
      fetchFn={fetchHistory}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState
        u={[
          balances.updater,
          incoming.updater,
          outgoing.updater,
          internals.updater,
          charts.updater,
          logs.updater,
          statements.updater,
          neighbors.updater,
          traces.updater,
          receipts.updater,
        ]}
      />
      <View tabItems={tabItems} searchable />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
