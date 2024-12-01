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

  // eslint-disable-next-line prefer-const
  let customTabs: string[] = [];
  // eslint-disable-next-line prefer-const
  let customForms: Record<string, JSX.Element> = {};
  // EXISTING_CODE
  customTabs = ["balances", "charts", "logs", "statements", "neighbors", "traces", "receipts"];
  customForms["balances"] = <div>This is a custom tab</div>;
  customForms["charts"] = <div>This is a custom tab</div>;
  customForms["logs"] = <div>This is a custom tab</div>;
  customForms["statements"] = <div>This is a custom tab</div>;
  customForms["neighbors"] = <div>This is a custom tab</div>;
  customForms["traces"] = <div>This is a custom tab</div>;
  customForms["receipts"] = <div>This is a custom tab</div>;
  // EXISTING_CODE

  const table = useReactTable({
    data: history?.items || [],
    columns: HistoryTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "history";
  const tabs = ["history", ...(customTabs || [])];
  const forms: ViewForm = {
    history: <FormTable data={history} groups={HistoryFormDef(table)} />,
    ...customForms,
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
      <DebugState u={[history.updater]} />
      <View tabs={tabs} forms={forms} searchable />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
