import { useState, useEffect } from "react";
import { Checkbox, InputLabel } from "@mantine/core";
import { View } from "@components";
import { GetSession, ModifyNoop } from "@gocode/app/App";
import { config } from "@gocode/models";
import { useAppState, ViewStateProvider } from "@state";

export const SettingsView = () => {
  const [session, setSession] = useState<config.Session | null>(null);
  const { status } = useAppState();

  useEffect(() => {
    GetSession().then((s) => setSession(s));
  }, []);

  const route = "settings";
  return (
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    <ViewStateProvider route={route} fetchFn={(_unused1: number, _unused2: number) => {}} modifyFn={ModifyNoop}>
      <View>
        <InputLabel>
          <Checkbox label={"A checkbox"} />
          <pre>{JSON.stringify(session, null, 2)}</pre>
          <pre>{status ? JSON.stringify(status, null, 2) : ""}</pre>
        </InputLabel>
      </View>
    </ViewStateProvider>
  );
};
