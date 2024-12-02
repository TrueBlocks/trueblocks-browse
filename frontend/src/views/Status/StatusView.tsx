// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm, DebugState } from "@components";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { StatusFormDef, StatusTableDef } from ".";
// EXISTING_CODE

export const StatusView = () => {
  const { status, fetchStatus } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  // EXISTING_CODE
  // EXISTING_CODE

  const table = useReactTable({
    data: status?.items || [],
    columns: StatusTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "status";
  const tabs = ["status"];
  const forms: ViewForm = {
    status: <FormTable data={status} groups={StatusFormDef(table)} />,
  };

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={status.nItems}
      fetchFn={fetchStatus}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState u={[status.updater]} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
