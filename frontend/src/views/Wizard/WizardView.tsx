// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm, DebugState } from "@components";
import { StepWizard } from "@gocode/app/App";
import { types } from "@gocode/models";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { WizardFormDef, WizardTableDef } from ".";
// EXISTING_CODE

export const WizardView = () => {
  const { wizard, fetchWizard } = useAppState();
  const { modifyNoop } = useNoops();
  const handleEnter = () => {
    stepWizard(types.WizStep.NEXT);
  };
  const handleModify = modifyNoop;

  // eslint-disable-next-line prefer-const
  let customTabs: string[] = [];
  // eslint-disable-next-line prefer-const
  let customForms: Record<string, JSX.Element> = {};
  // EXISTING_CODE
  const stepWizard = (step: string) => {
    StepWizard(step as types.WizStep).then(() => {
      fetchWizard(0, 100);
    });
  };
  // EXISTING_CODE

  const table = useReactTable({
    data: wizard?.items || [],
    columns: WizardTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "wizard";
  const tabs = ["wizard", ...(customTabs || [])];
  const forms: ViewForm = {
    wizard: <FormTable data={wizard} groups={WizardFormDef(table)} />,
    ...customForms,
  };

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={wizard.nItems}
      fetchFn={fetchWizard}
      onEnter={handleEnter}
      modifyFn={handleModify}
      clickFn={stepWizard}
      tabs={tabs}
    >
      <DebugState u={[wizard.updater]} />
      <View forms={forms} />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
