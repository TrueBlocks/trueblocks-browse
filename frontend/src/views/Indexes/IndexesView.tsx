// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { DebugState, FormTable, View, ViewForm } from "@components";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { IndexesTableDef, IndexesFormDef } from ".";
// EXISTING_CODE

export const IndexesView = () => {
  const { indexes, fetchIndexes } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  // EXISTING_CODE
  // EXISTING_CODE

  const table = useReactTable({
    data: indexes?.items || [],
    columns: IndexesTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "indexes";
  const tabs = ["indexes"];
  const forms: ViewForm = {
    indexes: <FormTable data={indexes} groups={IndexesFormDef(table)} />,
  };

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={indexes.nItems}
      fetchFn={fetchIndexes}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState u={indexes.updater} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
