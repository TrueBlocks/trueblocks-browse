import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm, DebugState } from "@components";
import { StepWizard } from "@gocode/app/App";
import { types } from "@gocode/models";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { WizardFormTable, WizardTableDef } from ".";

export const WizardView = () => {
  const { modifyNoop } = useNoops();
  const { wizard, fetchWizard } = useAppState();

  const stepWizard = (step: types.WizStep) => {
    StepWizard(step).then(() => {
      fetchWizard(0, 100);
    });
  };

  const handleEnter = () => {
    stepWizard(types.WizStep.NEXT);
  };

  const table = useReactTable({
    data: wizard.items || [],
    columns: WizardTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "wizard";
  const tabs = ["wizard"];
  const forms: ViewForm = {
    wizard: <FormTable data={wizard} groups={WizardFormTable(table, wizard.nItems, stepWizard)} />,
  };
  return (
    <ViewStateProvider
      route={route}
      nItems={wizard.nItems}
      fetchFn={fetchWizard}
      modifyFn={modifyNoop}
      onEnter={handleEnter}
    >
      <DebugState n={wizard.lastUpdate} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};
