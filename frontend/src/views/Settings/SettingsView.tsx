import { useState, useEffect } from "react";
import { FieldGroup, FormTable, View } from "@components";
import { GetSession } from "@gocode/app/App";
import { config } from "@gocode/models";
import { ViewStateProvider } from "@state";
import { useNoops } from "../../hooks";

export const SettingsView = () => {
  const { fetchNoop, modifyNoop } = useNoops();
  // TODO BOGUS: The settings state should be in the AppState
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
    <ViewStateProvider route={route} fetchFn={fetchNoop} modifyFn={modifyNoop}>
      <View>
        <FormTable data={session} groups={createSettingsForm()} />
      </View>
    </ViewStateProvider>
  );
};

const createSettingsForm = (): FieldGroup<config.Session>[] => {
  return [
    {
      label: "Session Data 1",
      colSpan: 6,
      collapsable: false,
      fields: [
        { label: "chain", type: "text", accessor: "chain" },
        { label: "lastFile", type: "text", accessor: "lastFile" },
        { label: "lastRoute", type: "text", accessor: "lastRoute" },
      ],
    },
    {
      label: "Session Data 2",
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
