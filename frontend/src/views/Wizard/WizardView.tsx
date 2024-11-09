import { useState, useEffect } from "react";
import { Text } from "@mantine/core";
import { StepWizard, GetWizErrs } from "@gocode/app/App";
import { types } from "@gocode/models";
import { useAppState } from "@state";
import { WizFiniButton, WizHomeButton, WizNextButton, WizPrevButton } from "./WizardFormDef";

export const WizardView = () => {
  const { isConfigured, wizState, setWizState } = useAppState();
  const [errors, setErrors] = useState<types.WizError[]>([]);

  const stepWizard = (step: types.WizStep) => {
    StepWizard(step).then((state) => {
      setWizState(state);
    });
  };

  useEffect(() => {
    GetWizErrs().then((errorList) => {
      setErrors(errorList);
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
      <WizHomeButton key="home" wizState={wizState} onClick={stepWizard} />{" "}
      <WizPrevButton key="prev" wizState={wizState} onClick={stepWizard} />{" "}
      <WizNextButton key="next" wizState={wizState} onClick={stepWizard} />{" "}
      <WizFiniButton key="fini" wizState={wizState} disabled={errors?.length > 0} onClick={stepWizard} />,
    </div>
  );
};
