import React, { useState, useEffect } from "react";
import { Checkbox, InputLabel } from "@mantine/core";
import { GetSession } from "@gocode/app/App";
import { config } from "@gocode/models";
import { View } from "@components";
import { useAppState } from "@state";

export function SettingsView() {
  const [session, setSession] = useState<config.Session | null>(null);
  const { status } = useAppState();

  useEffect(() => {
    GetSession().then((s) => setSession(s));
  }, []);

  return (
    <View>
      <InputLabel>
        <Checkbox label={"A checkbox"} />
        <pre>{JSON.stringify(session, null, 2)}</pre>
        <pre>{status ? JSON.stringify(status, null, 2) : ""}</pre>
      </InputLabel>
    </View>
  );
}
