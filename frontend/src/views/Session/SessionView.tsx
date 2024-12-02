// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, ViewForm, FormTable, DebugState } from "@components";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { SessionFormDef, SessionTableDef } from ".";
// EXISTING_CODE

export const SessionView = () => {
  const { session, fetchSession } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  // EXISTING_CODE
  // EXISTING_CODE

  const table = useReactTable({
    data: session?.items || [],
    columns: SessionTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "session";
  const tabs = ["session"];
  const forms: ViewForm = {
    session: <FormTable data={session} groups={SessionFormDef(table)} />,
  };

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={session.nItems}
      fetchFn={fetchSession}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState u={[session.updater]} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
