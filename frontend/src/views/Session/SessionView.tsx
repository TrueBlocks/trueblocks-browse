import { View, ViewForm, FormTable } from "@components";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { SessionFormDef } from "./SessionFormDef";

export const SessionView = () => {
  const { session, fetchSession } = useAppState();
  const { modifyNoop } = useNoops();

  const route = "session";
  const tabs = ["session"];
  const forms: ViewForm = {
    session: <FormTable data={session} groups={SessionFormDef()} />,
  };
  return (
    <ViewStateProvider route={route} fetchFn={fetchSession} modifyFn={modifyNoop}>
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};
