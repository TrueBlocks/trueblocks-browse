// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, TabItem, ViewForm, DebugState } from "@components";
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

  const stepWizard = (step: string) => {
    StepWizard(step as types.WizStep).then(() => {
      fetchWizard(0, 100);
    });
  };

  const table = useReactTable({
    data: wizard?.items || [],
    columns: WizardTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const tabItems: ViewForm = {
    wizard: <TabItem data={wizard} groups={WizardFormDef(table)} />,
  };

  return (
    <ViewStateProvider
      // do not remove - delint
      nItems={wizard.nItems}
      fetchFn={fetchWizard}
      onEnter={handleEnter}
      modifyFn={handleModify}
      clickFn={stepWizard}
    >
      <DebugState u={[wizard.updater]} />
      <View tabItems={tabItems} />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
