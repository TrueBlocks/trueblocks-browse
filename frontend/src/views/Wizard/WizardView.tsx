import { useState, useEffect } from "react";
import { Button, Text } from "@mantine/core";
import { StepWizard, GetDeferredErrors } from "@gocode/app/App";
import { types } from "@gocode/models";
import { useAppState } from "@state";
import classes from "./WizardView.module.css";

export const WizardView = () => {
  const { isConfigured, wizardState, setWizardState } = useAppState();
  const [cn, setCn] = useState(classes.wizOkay);
  const [errors, setErrors] = useState<types.WizardError[]>([]);

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
      <WizResetButton step={stepWizard} state={wizardState} />
      <WizBumpButton step={stepWizard} state={wizardState} back />
      <WizBumpButton step={stepWizard} state={wizardState} />
      <WizFinishButton step={stepWizard} state={wizardState} />
    </div>
  );
};

type StepProps = {
  step: (step: types.WizStep) => void;
  state: types.WizState;
  back?: boolean;
};

export const WizResetButton = ({ step }: StepProps) => {
  return (
    <Button size={"xs"} onClick={() => step(types.WizStep.RESET)}>
      Reset
    </Button>
  );
};

export const WizBumpButton = ({ step, state, back = false }: StepProps) => {
  const bDis = back && (state === types.WizState.WELCOME || state === types.WizState.ERROR);
  const fDis = !back && state === types.WizState.OKAY;
  return (
    <Button
      disabled={bDis || fDis}
      size={"xs"}
      onClick={() => step(back ? types.WizStep.PREVIOUS : types.WizStep.NEXT)}
    >
      {back ? "Back" : "Next"}
    </Button>
  );
};

export const WizFinishButton = ({ step }: StepProps) => {
  return (
    <Button size={"xs"} onClick={() => step(types.WizStep.FINISH)}>
      Finish
    </Button>
  );
};
