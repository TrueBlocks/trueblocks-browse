import { Button } from "@mantine/core";
import { Table } from "@tanstack/react-table";
// import { DataTable, FieldGroup, PublishButton, CleanButton, AddButton } from "@components";
import { DataTable, FieldGroup } from "@components";
import { types } from "@gocode/models";
import { StepWizard } from "../../../wailsjs/go/app/App";
import { useAppState } from "../../state";

export const WizardFormTable = (table: Table<types.WizError>, nItems: number): FieldGroup<types.WizardContainer>[] => {
  const { wizard, fetchWizard } = useAppState();

  const stepWizard = (step: types.WizStep) => {
    StepWizard(step).then(() => {
      fetchWizard(0, 100);
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
        <WizHomeButton key="home" state={wizard.state} onClick={stepWizard} />,
        <WizPrevButton key="prev" state={wizard.state} onClick={stepWizard} />,
        <WizNextButton key="next" state={wizard.state} onClick={stepWizard} />,
        <WizFiniButton key="fini" state={wizard.state} disabled={nItems > 0} onClick={stepWizard} />,
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
  state: types.WizState;
  disabled?: boolean;
  onClick: (step: types.WizStep) => void;
};

export const WizHomeButton = ({ state, onClick, disabled = false }: StepProps) => {
  disabled = state === types.WizState.WELCOME || disabled;
  return (
    <Button disabled={disabled} size={"xs"} onClick={() => onClick(types.WizStep.FIRST)}>
      First
    </Button>
  );
};

export const WizPrevButton = ({ state, onClick, disabled = false }: StepProps) => {
  disabled = state === types.WizState.WELCOME || disabled;
  return (
    <Button disabled={disabled} size={"xs"} onClick={() => onClick(types.WizStep.PREVIOUS)}>
      Back
    </Button>
  );
};

export const WizNextButton = ({ state, onClick, disabled = false }: StepProps) => {
  disabled = state === types.WizState.INDEX || disabled;
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
