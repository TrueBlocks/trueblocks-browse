import { useState, useEffect } from "react";
import { Button, Text } from "@mantine/core";
import { StepWizard, GetDeferredErrors } from "@gocode/app/App";
import { types } from "@gocode/models";
import { useAppState } from "@state";

export const WizardView = () => {
  const { isConfigured, wizState, setWizState } = useAppState();
  const [errors, setErrors] = useState<types.WizError[]>([]);
  const [prevDisabled, setPrevDisabled] = useState(false);
  const [nextDisabled, setNextDisabled] = useState(false);
  const [finishDisabled, setFinishDisabled] = useState(false);

  const stepWizard = (step: types.WizStep) => {
    StepWizard(step).then((state) => {
      setWizState(state);
      setPrevDisabled(state === types.WizState.WELCOME);
      setNextDisabled(state === types.WizState.FINISHED);
    });
  };

  useEffect(() => {
    GetDeferredErrors().then((errorList) => {
      setErrors(errorList);
      setFinishDisabled(errorList?.length > 0);
    });
  }, [wizState]);

  return (
    <div>
      <Text>{`wizState: ${wizState}`}</Text>
      <Text>{`isConfigured: ${isConfigured}`}</Text>
      {errors?.length > 0 && (
        <div>
          {errors.map((wizErr, index) => (
            <div key={index}>{`n: ${wizErr.count} err: ${wizErr.error}`}</div>
          ))}
        </div>
      )}
      <WizFirstButton disabled={false} step={stepWizard} />{" "}
      <WizBumpButton disabled={prevDisabled} step={stepWizard} back />{" "}
      <WizBumpButton disabled={nextDisabled} step={stepWizard} />{" "}
      <WizFinishButton disabled={finishDisabled} step={stepWizard} />
    </div>
  );
};

type StepProps = {
  step: (step: types.WizStep) => void;
  disabled: boolean;
  back?: boolean;
};

export const WizFirstButton = ({ step }: StepProps) => {
  return (
    <Button size={"xs"} onClick={() => step(types.WizStep.FIRST)}>
      First
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
