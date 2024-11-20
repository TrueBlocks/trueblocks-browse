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

  // EXISTING_CODE
  const stepWizard = (step: types.WizStep) => {
    StepWizard(step).then(() => {
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
  const tabs = ["wizard"];
  const forms: ViewForm = {
    wizard: <FormTable data={wizard} groups={WizardFormDef(table, wizard.nItems, stepWizard)} />,
  };

  // if (!(wizard?.items?.length > 0)) {
  //   return <>Loading...</>;
  // }

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={wizard.nItems}
      fetchFn={fetchWizard}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState n={wizard.lastUpdate} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
