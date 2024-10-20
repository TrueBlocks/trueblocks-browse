import { useState, useEffect } from "react";
import { FieldGroup, FormTable, View } from "@components";
import { GetSession, ModifyNoop } from "@gocode/app/App";
import { config } from "@gocode/models";
import { ViewStateProvider } from "@state";

export const SettingsView = () => {
  const [session, setSession] = useState<config.Session | null>(null);
  // const [config, setConfig] = useState<config.ConfigFile | null>(null);

  useEffect(() => {
    GetSession().then((s) => setSession(s));
  }, []);

  if (!session) {
    return <div>Loading...</div>;
  }

  const route = "settings";

  return (
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    <ViewStateProvider route={route} fetchFn={(_unused1: number, _unused2: number) => {}} modifyFn={ModifyNoop}>
      <View>
        <FormTable data={session} groups={createSettingsForm()} />
      </View>
    </ViewStateProvider>
  );
};

const createSettingsForm = (): FieldGroup<config.Session>[] => {
  return [
    {
      legend: "Session Data 1",
      colSpan: 6,
      collapsable: false,
      fields: [
        { label: "chain", type: "text", accessor: "chain" },
        { label: "lastFile", type: "text", accessor: "lastFile" },
        { label: "lastRoute", type: "text", accessor: "lastRoute" },
      ],
    },
    {
      legend: "Session Data 2",
      colSpan: 6,
      collapsable: false,
      fields: [
        // { label: "lastSub", type: "text", accessor: "lastSub" },
        // { label: "window", type: "text", accessor: "window" },
        // { label: "daemons", type: "text", accessor: "lastRoute" },
        // { label: "wizard", type: "text", accessor: "lastRoute" },
        // { label: "toggles", type: "text", accessor: "lastRoute" },
      ],
    },
  ];
};
