import React from "react";
import { Checkbox, InputLabel } from "@mantine/core";
import { View2 } from "@components";

export function SettingsView() {
  return (
    <View2>
      <InputLabel>
        <Checkbox label={"A checkbox"} />
      </InputLabel>
    </View2>
  );
}
