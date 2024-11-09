import { Button } from "@mantine/core";
import { Table } from "@tanstack/react-table";
// import { DataTable, FieldGroup, PublishButton, CleanButton, AddButton } from "@components";
import { DataTable, FieldGroup } from "@components";
import { types } from "@gocode/models";
import { StepWizard } from "../../../wailsjs/go/app/App";
import { useAppState } from "../../state";

export const WizardFormTable = (table: Table<types.WizError>, nItems: number): FieldGroup<types.WizardContainer>[] => {
  const { wizState, setWizState } = useAppState();

  const stepWizard = (step: types.WizStep) => {
    StepWizard(step).then((state) => {
      setWizState(state);
    });
  };

  return [
    {
      label: "Wizard State",
      colSpan: 12,
      collapsable: false,
      fields: [
        { label: "nItems", type: "int", accessor: "nItems" },
        { label: "Chain", type: "text", accessor: "chain" },
        { label: "lastUpdate", type: "date", accessor: "lastUpdate" },
        { label: "state", type: "text", accessor: "state" },
      ],
    },
    {
      label: "Buttons",
      buttons: [
        <WizHomeButton key="home" wizState={wizState} onClick={stepWizard} />,
        <WizPrevButton key="prev" wizState={wizState} onClick={stepWizard} />,
        <WizNextButton key="next" wizState={wizState} onClick={stepWizard} />,
        <WizFiniButton key="fini" wizState={wizState} disabled={nItems > 0} onClick={stepWizard} />,
      ],
    },
    {
      label: "Errors",
      collapsable: false,
      components: [<DataTable<types.WizError> key={"dataTable"} table={table} loading={false} />],
    },
  ];
};

type StepProps = {
  wizState: types.WizState;
  disabled?: boolean;
  onClick: (step: types.WizStep) => void;
};

export const WizHomeButton = ({ wizState, onClick, disabled = false }: StepProps) => {
  disabled = wizState === types.WizState.WELCOME || disabled;
  return (
    <Button disabled={disabled} size={"xs"} onClick={() => onClick(types.WizStep.FIRST)}>
      First
    </Button>
  );
};

export const WizPrevButton = ({ wizState, onClick, disabled = false }: StepProps) => {
  disabled = wizState === types.WizState.WELCOME || disabled;
  return (
    <Button disabled={disabled} size={"xs"} onClick={() => onClick(types.WizStep.PREVIOUS)}>
      Back
    </Button>
  );
};

export const WizNextButton = ({ wizState, onClick, disabled = false }: StepProps) => {
  disabled = wizState === types.WizState.FINISHED || disabled;
  return (
    <Button disabled={disabled} size={"xs"} onClick={() => onClick(types.WizStep.NEXT)}>
      Next
    </Button>
  );
};

export const WizFiniButton = ({ onClick, disabled = false }: StepProps) => {
  return (
    <Button disabled={disabled} size={"xs"} onClick={() => onClick(types.WizStep.FINISH)}>
      Finish
    </Button>
  );
};
