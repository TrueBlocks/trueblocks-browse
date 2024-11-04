import { Text } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm } from "@components";
import { useNoops, useRenderCounter } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { ManifestsTableDef, ManifestsFormDef } from ".";

export const ManifestsView = () => {
  const { modifyNoop } = useNoops();
  const { manifests, fetchManifests } = useAppState();
  const renderCount = useRenderCounter();

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
      <Text>Render count: {renderCount}</Text>
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};
