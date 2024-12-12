// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { useCallback } from "react";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { DebugState, TabItem, View, ViewForm } from "@components";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
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
// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// import { Table } from "@tanstack/react-table";
// import { ExploreButton, ExportButton, DataTable, FieldGroup, GoogleButton } from "@components";
// import { types } from "@gocode/models";
// import { useAppState } from "@state";

// export const HistoryFormDef = (table: Table<types.Transaction>): FieldGroup<types.HistoryContainer>[] => {
//   const { info } = useAppState();
//   return [
//     {
//       label: "DalleDress",
//       colSpan: 2,
//       fields: [{ label: "", type: "dalle", accessor: "address" }],
//     },
//     {
//       label: "Transaction Data",
//       colSpan: 7,
//       fields: [
//         { label: "name", type: "address-name-only", accessor: "address" },
//         { label: "address", type: "address-address-only", accessor: "address" },
//         { label: "balance", type: "ether", accessor: "balance" },
//       ],
//     },
//     {
//       label: "Transaction Data",
//       colSpan: 3,
//       fields: [
//         { label: "nTransactions", type: "int", accessor: "nItems" },
//         { label: "nLogs", type: "int", accessor: "nLogs" },
//         { label: "nTokens", type: "int", accessor: "nTokens" },
//         { label: "nErrors", type: "int", accessor: "nErrors" },
//       ],
//     },
//     {
//       label: "Buttons",
//       buttons: [
//         <ExploreButton key={"explore"} value={info.address} />,
//         <GoogleButton key={"google"} value={info.address} />,
//         <ExportButton key={"export"} value={info.address} />,
//       ],
//     },
//     {
//       label: "Transactions",
//       collapsable: false,
//       components: [<DataTable<types.Transaction> key={"dataTable"} table={table} loading={false} />],
//     },
//   ];
// };

// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// import { createColumnHelper } from "@tanstack/react-table";
// import { CustomColumnDef, Formatter } from "@components";
// import { types } from "@gocode/models";

// const columnHelper = createColumnHelper<types.Transaction>();

// export const HistoryTableDef: CustomColumnDef<types.Transaction, any>[] = [
//   columnHelper.accessor("blockNumber", {
//     header: () => "TxId",
//     cell: (info) => {
//       const { blockNumber, transactionIndex, hash } = info.row.original;
//       return <Formatter type="appearance" value={`${blockNumber}.${transactionIndex}`} value2={hash} />;
//     },
//     meta: { className: "medium cell" },
//   }),
//   columnHelper.accessor("timestamp", {
//     id: "Timestamp",
//     cell: (info) => <Formatter type="timestamp" value={info.renderValue()} />,
//     meta: { className: "medium cell" },
//   }),
//   columnHelper.accessor("from", {
//     header: () => "From",
//     cell: (info) => <Formatter type="address-editor" value={info.renderValue()} />,
//     meta: { className: "wide cell" },
//   }),
//   columnHelper.accessor("to", {
//     header: () => "To",
//     cell: (info) => <Formatter type="address-editor" value={info.renderValue()} />,
//     meta: { className: "wide cell" },
//   }),
//   columnHelper.accessor("value", {
//     header: () => "Ether",
//     cell: (info) => {
//       const { value, isError } = info.row.original;
//       return <Formatter type="ether" value={value} value2={isError} />;
//     },
//     meta: { className: "medium cell" },
//   }),
//   columnHelper.accessor("gasUsed", {
//     header: () => "Gas",
//     cell: (info) => {
//       const { gasUsed, gasPrice, from } = info.row.original;
//       const gasCost = gasUsed * gasPrice;
//       return <Formatter type="gas" value={gasCost} value2={from} />;
//     },
//     meta: { className: "medium cell" },
//   }),
//   columnHelper.accessor("isError", {
//     header: () => "isError",
//     cell: (info) => <Formatter type="error" value={info.renderValue()} />,
//     meta: { className: "small center cell" },
//   }),
//   columnHelper.accessor("input", {
//     header: () => "Function",
//     cell: (info) => <Formatter type="text" value={info.renderValue()} />,
//     meta: { className: "medium center cell" },
//   }),
// ];

// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
// import { View, FormTable, ViewForm, DebugState } from "@components";
// import { useNoops } from "@hooks";
// import { useAppState, ViewStateProvider } from "@state";
// import { HistoryTableDef, HistoryFormDef } from ".";

// export const HistoryView = () => {
//   const { history, fetchHistory } = useAppState();
//   const { enterNoop, modifyNoop } = useNoops();
//   const handleEnter = enterNoop;
//   const handleModify = modifyNoop;

//   // eslint-disable-next-line prefer-const
//   let customTabs: string[] = [];
//   // eslint-disable-next-line prefer-const
//   let customForms: Record<string, JSX.Element> = {};
//   customTabs = [
//     "balances",
//     "incoming",
//     "outgoing",
//     "internal",
//     "charts",
//     "logs",
//     "statements",
//     "neighbors",
//     "traces",
//     "receipts",
//   ];
//   customForms["balances"] = <div>This is the balances tab</div>;
//   customForms["incoming"] = <div>This is the incoming tab</div>;
//   customForms["outgoing"] = <div>This is the outgoing tab</div>;
//   customForms["internal"] = <div>This is the internal tab</div>;
//   customForms["charts"] = <div>This is the charts tab</div>;
//   customForms["logs"] = <div>This is the logs tab</div>;
//   customForms["statements"] = <div>This is the statements tab</div>;
//   customForms["neighbors"] = <div>This is the neighbors tab</div>;
//   customForms["traces"] = <div>This is the traces tab</div>;
//   customForms["receipts"] = <div>This is the receipts tab</div>;

//   const table = useReactTable({
//     data: history?.items || [],
//     columns: HistoryTableDef,
//     getCoreRowModel: getCoreRowModel(),
//   });

//   const tabs = ["history", ...(customTabs || [])];
//   const forms: ViewForm = {
//     history: <FormTable data={history} groups={HistoryFormDef(table)} />,
//     ...customForms,
//   };

//   return (
//     <ViewStateProvider
//       // do not remove - delint
//       nItems={history.nItems}
//       fetchFn={fetchHistory}
//       onEnter={handleEnter}
//       modifyFn={handleModify}
//       tabs={tabs}
//     >
//       <DebugState u={[history.updater]} />
//       <View forms={forms} searchable />
//     </ViewStateProvider>
//   );
// };

// EXISTING_CODE
