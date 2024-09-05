import { Button } from "@mantine/core";
import { wizard } from "@gocode/models";
import { useAppState } from "@state";

export function WizardView() {
  const { isConfigured, wizardState } = useAppState();

  return (
    <div>
      <div>{`wizardState: ${wizardState}`}</div>
      <div>{`isConfigured: ${isConfigured}`}</div>
      <ResetWizard />
      <StepWizard back />
      <StepWizard />
      <FinishWizard />
    </div>
  );
}

export function ResetWizard() {
  const { stepWizard } = useAppState();
  return (
    <Button size={"xs"} onClick={() => stepWizard(wizard.Step.RESET)}>
      Reset
    </Button>
  );
}

export function StepWizard({ back = false }: { back?: boolean }) {
  const { stepWizard } = useAppState();
  return (
    <Button size={"xs"} onClick={() => stepWizard(back ? wizard.Step.PREVIOUS : wizard.Step.NEXT)}>
      {back ? "Back" : "Next"}
    </Button>
  );
}

export function FinishWizard() {
  const { stepWizard } = useAppState();
  return (
    <Button size={"xs"} onClick={() => stepWizard(wizard.Step.FINISH)}>
      Finish
    </Button>
  );
}
