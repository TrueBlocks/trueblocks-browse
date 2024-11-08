import { useState, useEffect } from "react";
import { Button, Text } from "@mantine/core";
import { StepWizard, GetDeferredErrors } from "@gocode/app/App";
import { app, types } from "@gocode/models";
import { useAppState } from "@state";
import classes from "./WizardView.module.css";

export const WizardView = () => {
  const { isConfigured, wizardState, setWizardState } = useAppState();
  const [cn, setCn] = useState(classes.wizOkay);
  const [errors, setErrors] = useState<app.WizError[]>([]);

  const stepWizard = (step: types.WizStep) => {
    StepWizard(step).then((state) => {
      setWizardState(state);
    });
  };

  useEffect(() => {
    setCn(wizardState === types.WizState.ERROR ? classes.wizError : classes.wizOkay);
  }, [wizardState]);

  useEffect(() => {
    GetDeferredErrors().then((errorList) => {
      setErrors(errorList);
    });
  }, [wizardState]);

  return (
    <div>
      <Text className={cn}>{`wizardState: ${wizardState}`}</Text>
      <Text className={cn}>{`isConfigured: ${isConfigured}`}</Text>
      {errors?.length > 0 && (
        <div>
          {errors.map((wizErr, index) => (
            <div key={index}>{`n: ${wizErr.count} err: ${wizErr.error}`}</div>
          ))}
        </div>
      )}
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
