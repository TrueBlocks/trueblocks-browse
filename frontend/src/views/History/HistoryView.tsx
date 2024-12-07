// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm, DebugState } from "@components";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { HistoryTableDef, HistoryFormDef } from ".";
// EXISTING_CODE

export const HistoryView = () => {
  const { history, fetchHistory } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  let customTabs: string[] = [];
  // eslint-disable-next-line prefer-const
  let customForms: Record<string, JSX.Element> = {};
  // EXISTING_CODE
  customTabs = [
    "balances",
    "incoming",
    "outgoing",
    "internal",
    "charts",
    "logs",
    "statements",
    "neighbors",
    "traces",
    "receipts",
  ];
  customForms["balances"] = <div>This is the balances tab</div>;
  customForms["incoming"] = <div>This is the incoming tab</div>;
  customForms["outgoing"] = <div>This is the outgoing tab</div>;
  customForms["internal"] = <div>This is the internal tab</div>;
  customForms["charts"] = <div>This is the charts tab</div>;
  customForms["logs"] = <div>This is the logs tab</div>;
  customForms["statements"] = <div>This is the statements tab</div>;
  customForms["neighbors"] = <div>This is the neighbors tab</div>;
  customForms["traces"] = <div>This is the traces tab</div>;
  customForms["receipts"] = <div>This is the receipts tab</div>;
  // EXISTING_CODE

  const table = useReactTable({
    data: history?.items || [],
    columns: HistoryTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const tabs = ["history", ...(customTabs || [])];
  const forms: ViewForm = {
    history: <FormTable data={history} groups={HistoryFormDef(table)} />,
    ...customForms,
  };

  return (
    <ViewStateProvider
      // do not remove - delint
      nItems={history.nItems}
      fetchFn={fetchHistory}
      onEnter={handleEnter}
      modifyFn={handleModify}
      tabs={tabs}
    >
      <DebugState u={[history.updater]} />
      <View forms={forms} searchable />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
