// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm, DebugState } from "@components";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { ManifestsTableDef, ManifestsFormDef } from ".";
// EXISTING_CODE

export const ManifestsView = () => {
  const { manifests, fetchManifests } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  // EXISTING_CODE
  // EXISTING_CODE

  const table = useReactTable({
    data: manifests?.items || [],
    columns: ManifestsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "manifests";
  const tabs = ["manifests"];
  const forms: ViewForm = {
    manifests: <FormTable data={manifests} groups={ManifestsFormDef(table)} />,
  };

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={manifests.nItems}
      fetchFn={fetchManifests}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState u={[manifests.updater]} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
