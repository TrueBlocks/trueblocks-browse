// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { View, ViewForm, FormTable, DebugState } from "@components";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { SessionFormDef } from "./SessionFormDef";
// EXISTING_CODE

export const SessionView = () => {
  const { session, fetchSession } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();

  // EXISTING_CODE
  // EXISTING_CODE

  const route = "session";
  const tabs = ["session"];
  const forms: ViewForm = {
    session: <FormTable data={session} groups={SessionFormDef(session)} />,
  };
  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={0}
      fetchFn={fetchSession}
      onEnter={enterNoop}
      modifyFn={modifyNoop}
    >
      <DebugState n={session.lastUpdate} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
