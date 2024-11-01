import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm } from "@components";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { ManifestsTableDef, ManifestsFormDef } from ".";

export const ManifestsView = () => {
  const { modifyNoop } = useNoops();
  const { manifests, fetchManifests } = useAppState();

  const table = useReactTable({
    data: manifests.items || [],
    columns: ManifestsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "manifests";
  const tabs = ["manifests"];
  const forms: ViewForm = {
    manifests: <FormTable data={manifests} groups={ManifestsFormDef(table)} />,
  };
  return (
    <ViewStateProvider route={route} nItems={manifests.nItems} fetchFn={fetchManifests} modifyFn={modifyNoop}>
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};
