import { Button } from "@mantine/core";
import { StepWizard } from "@gocode/app/App";
import { wizard } from "@gocode/models";
import { useAppState } from "@state";

export function WizardView() {
  const { isConfigured, wizardState, setWizardState } = useAppState();
  const stepWizard = (step: wizard.Step) => {
    StepWizard(step).then((state) => {
      setWizardState(state);
    });
  };

  return (
    <div>
      <div>{`wizardState: ${wizardState}`}</div>
      <div>{`isConfigured: ${isConfigured}`}</div>
      <ResetWizard stepWizard={stepWizard} />
      <BumpWizard stepWizard={stepWizard} back />
      <BumpWizard stepWizard={stepWizard} />
      <FinishWizard stepWizard={stepWizard} />
    </div>
  );
}

type StepProps = {
  stepWizard: (step: wizard.Step) => void;
  back?: boolean;
};

export function ResetWizard({ stepWizard }: StepProps) {
  return (
    <Button size={"xs"} onClick={() => stepWizard(wizard.Step.RESET)}>
      Reset
    </Button>
  );
}

export function BumpWizard({ stepWizard, back = false }: StepProps) {
  return (
    <Button size={"xs"} onClick={() => stepWizard(back ? wizard.Step.PREVIOUS : wizard.Step.NEXT)}>
      {back ? "Back" : "Next"}
    </Button>
  );
}

export function FinishWizard({ stepWizard }: StepProps) {
  return (
    <Button size={"xs"} onClick={() => stepWizard(wizard.Step.FINISH)}>
      Finish
    </Button>
  );
}
