import { Button } from "@mantine/core";
import { StepWizard } from "@gocode/app/App";
import { types } from "@gocode/models";
import { useAppState } from "@state";

export const WizardView = () => {
  const { isConfigured, wizardState, setWizardState } = useAppState();
  const stepWizard = (step: types.WizStep) => {
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
};

type StepProps = {
  stepWizard: (step: types.WizStep) => void;
  back?: boolean;
};

export const ResetWizard = ({ stepWizard }: StepProps) => {
  return (
    <Button size={"xs"} onClick={() => stepWizard(types.WizStep.RESET)}>
      Reset
    </Button>
  );
};

export const BumpWizard = ({ stepWizard, back = false }: StepProps) => {
  return (
    <Button size={"xs"} onClick={() => stepWizard(back ? types.WizStep.PREVIOUS : types.WizStep.NEXT)}>
      {back ? "Back" : "Next"}
    </Button>
  );
};

export const FinishWizard = ({ stepWizard }: StepProps) => {
  return (
    <Button size={"xs"} onClick={() => stepWizard(types.WizStep.FINISH)}>
      Finish
    </Button>
  );
};
