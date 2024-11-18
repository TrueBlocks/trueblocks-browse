// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { Button } from "@mantine/core";
import { Table } from "@tanstack/react-table";
import { DataTable, FieldGroup } from "@components";
import { types } from "@gocode/models";
import { useAppState } from "@state";
// EXISTING_CODE

export const WizardFormDef = (
  table: Table<types.WizError>,
  nItems: number,
  stepWizard: (step: types.WizStep) => void
): FieldGroup<types.WizardContainer>[] => {
  const { wizard } = useAppState();
  return [
    // EXISTING_CODE
    {
      label: "Wizard State",
      colSpan: 12,
      collapsable: false,
      fields: [
        { label: "nItems", type: "int", accessor: "nItems" },
        { label: "chain", type: "text", accessor: "chain" },
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
    // EXISTING_CODE
  ];
};

// EXISTING_CODE
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

// EXISTING_CODE

//-------------------------------------------------------------------
// Template variables:
// class:         Wizard
// routeLabel:    Wizard
// itemName:      WizError
// isHistory:     false
// isSession:     false
// isConfig:      false
