import React from "react";
import { Checkbox, InputLabel } from "@mantine/core";
import View from "@/components/view/View";

function SettingsView() {
  return (
    <View title="Settings">
      <InputLabel>
        <Checkbox /> A checkbox
      </InputLabel>
    </View>
  );
}

export default SettingsView;
