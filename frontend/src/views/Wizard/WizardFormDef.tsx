// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { Button } from "@mantine/core";
import { Table } from "@tanstack/react-table";
import { DataTable, FieldGroup } from "@components";
import { types } from "@gocode/models";
import { useAppState, useViewState } from "@state";
// EXISTING_CODE

export const WizardFormDef = (table: Table<types.WizError>): FieldGroup<types.WizardContainer>[] => {
  // EXISTING_CODE
  const { wizard } = useAppState();
  // EXISTING_CODE
  return [
    // EXISTING_CODE
    {
      label: "Wizard State",
      colSpan: 12,
      collapsable: false,
      fields: [
        { label: "nItems", type: "int", accessor: "nItems" },
        { label: "chain", type: "text", accessor: "chain" },
        { label: "state", type: "text", accessor: "state" },
      ],
    },
    {
      label: "Buttons",
      buttons: [
        <WizHomeButton key="home" state={wizard.state} />,
        <WizPrevButton key="prev" state={wizard.state} />,
        <WizNextButton key="next" state={wizard.state} />,
        <WizFiniButton key="fini" state={wizard.state} disabled={wizard.nItems > 0} />,
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
};

export const WizHomeButton = ({ state, disabled = false }: StepProps) => {
  const { clickFn } = useViewState();
  disabled = state === types.WizState.WELCOME || disabled;
  return (
    <Button disabled={disabled} size={"xs"} onClick={() => clickFn && clickFn(types.WizStep.FIRST)}>
      First
    </Button>
  );
};

export const WizPrevButton = ({ state, disabled = false }: StepProps) => {
  const { clickFn } = useViewState();
  disabled = state === types.WizState.WELCOME || disabled;
  return (
    <Button disabled={disabled} size={"xs"} onClick={() => clickFn && clickFn(types.WizStep.PREVIOUS)}>
      Back
    </Button>
  );
};

export const WizNextButton = ({ state, disabled = false }: StepProps) => {
  const { clickFn } = useViewState();
  disabled = state === types.WizState.INDEX || disabled;
  return (
    <Button disabled={disabled} size={"xs"} onClick={() => clickFn && clickFn(types.WizStep.NEXT)}>
      Next
    </Button>
  );
};

export const WizFiniButton = ({ disabled = false }: StepProps) => {
  const { clickFn } = useViewState();
  return (
    <Button disabled={disabled} size={"xs"} onClick={() => clickFn && clickFn(types.WizStep.FINISH)}>
      Finish
    </Button>
  );
};

// EXISTING_CODE
