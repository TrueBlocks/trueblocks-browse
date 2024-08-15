import React, { useState, useEffect } from "react";
import { Checkbox, InputLabel } from "@mantine/core";
import { GetSession } from "@gocode/app/App";
import { config } from "@gocode/models";
import { View2 } from "@components";

export function SettingsView() {
  const [session, setSession] = useState<config.Session | null>(null);

  useEffect(() => {
    GetSession().then((s) => setSession(s));
  }, []);

  return (
    <View2>
      <InputLabel>
        <Checkbox label={"A checkbox"} />
        <pre>{JSON.stringify(session, null, 2)}</pre>
      </InputLabel>
    </View2>
  );
}
