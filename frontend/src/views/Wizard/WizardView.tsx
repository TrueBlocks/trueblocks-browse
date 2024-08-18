import React, { useEffect, useState } from "react";
import { Button } from "@mantine/core";
// import { useAppState } from "@state";
import { wizard } from "@gocode/models";

export function WizardView() {
  return <>Wizard</>;
  // const { isConfigured, wizardState, stepWizard } = useAppState();

  // return (
  //   <div>
  //     <div>{`wizardState: ${wizardState}`}</div>
  //     <div>{`isConfigured: ${isConfigured}`}</div>
  //     <ResetWizard />
  //     <StepWizard back />
  //     <StepWizard />
  //     <FinishWizard />
  //   </div>
  // );
}

export function ResetWizard() {
  return <>Wizard</>;
  // const { stepWizard } = useAppState();
  // return (
  //   <Button size={"xs"} onClick={() => stepWizard(wizard.Step.RESET)}>
  //     Reset
  //   </Button>
  // );
}

export function StepWizard({ back = false }: { back?: boolean }) {
  return <>Wizard</>;
  // const { stepWizard } = useAppState();
  // return (
  //   <Button size={"xs"} onClick={() => stepWizard(back ? wizard.Step.PREVIOUS : wizard.Step.NEXT)}>
  //     {back ? "Back" : "Next"}
  //   </Button>
  // );
}

export function FinishWizard() {
  return <>Wizard</>;
  // const { stepWizard } = useAppState();
  // return (
  //   <Button size={"xs"} onClick={() => stepWizard(wizard.Step.FINISH)}>
  //     Finish
  //   </Button>
  // );
}
