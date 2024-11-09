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
  const [prevDisabled, setPrevDisabled] = useState(false);
  const [nextDisabled, setNextDisabled] = useState(false);
  const [finishDisabled, setFinishDisabled] = useState(false);

  const stepWizard = (step: types.WizStep) => {
    StepWizard(step).then((state) => {
      setWizardState(state);
      setPrevDisabled(state === types.WizState.WELCOME || state === types.WizState.ERROR);
      setNextDisabled(state === types.WizState.FINISHED);
    });
  };

  useEffect(() => {
    setCn(wizardState === types.WizState.ERROR ? classes.wizError : classes.wizOkay);
  }, [wizardState]);

  useEffect(() => {
    GetDeferredErrors().then((errorList) => {
      setErrors(errorList);
      setFinishDisabled(errorList?.length > 0);
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
      <WizResetButton disabled={false} step={stepWizard} state={wizardState} />
      <WizBumpButton disabled={prevDisabled} step={stepWizard} state={wizardState} back />
      <WizBumpButton disabled={nextDisabled} step={stepWizard} state={wizardState} />
      <WizFinishButton disabled={finishDisabled} step={stepWizard} state={wizardState} />
    </div>
  );
};

type StepProps = {
  step: (step: types.WizStep) => void;
  state: types.WizState;
  disabled: boolean;
  back?: boolean;
};

export const WizResetButton = ({ step }: StepProps) => {
  return (
    <Button size={"xs"} onClick={() => step(types.WizStep.RESET)}>
      Reset
    </Button>
  );
};

export const WizBumpButton = ({ step, disabled, back = false }: StepProps) => {
  return (
    <Button disabled={disabled} size={"xs"} onClick={() => step(back ? types.WizStep.PREVIOUS : types.WizStep.NEXT)}>
      {back ? "Back" : "Next"}
    </Button>
  );
};

export const WizFinishButton = ({ step, disabled }: StepProps) => {
  return (
    <Button disabled={disabled} size={"xs"} onClick={() => step(types.WizStep.FINISH)}>
      Finish
    </Button>
  );
};
