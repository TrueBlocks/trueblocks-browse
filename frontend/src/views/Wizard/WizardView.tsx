import React, { useEffect, useState } from "react";
import { Button } from "@mantine/core";
import { useAppState } from "@state";
import { wizard } from "@gocode/models";

export function WizardView() {
  const { isConfigured, wizardState, stepWizard } = useAppState();

  return (
    <div>
      <div>{`wizardState: ${wizardState}`}</div>
      <div>{`isConfigured: ${isConfigured}`}</div>
      <Button size={"xs"} onClick={() => stepWizard(wizard.Step.RESET)}>
        Reset
      </Button>
      <Button size={"xs"} onClick={() => stepWizard(wizard.Step.PREVIOUS)}>
        Prev
      </Button>
      <Button size={"xs"} onClick={() => stepWizard(wizard.Step.NEXT)}>
        Next
      </Button>
      <Button size={"xs"} onClick={() => stepWizard(wizard.Step.FINISH)}>
        Finish
      </Button>
    </div>
  );
}
