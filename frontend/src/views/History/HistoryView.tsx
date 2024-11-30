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

  // EXISTING_CODE
  // EXISTING_CODE

  const table = useReactTable({
    data: history?.items || [],
    columns: HistoryTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "history";
  const tabs = ["history"];
  const forms: ViewForm = {
    history: <FormTable data={history} groups={HistoryFormDef(table)} />,
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
