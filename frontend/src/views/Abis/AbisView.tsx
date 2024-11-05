import { Text } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm } from "@components";
import { ModifyAbi } from "@gocode/app/App";
import { useRenderCounter } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { AbisFormTable, AbisTableDef } from ".";

export const AbisView = () => {
  const { abis, fetchAbis } = useAppState();
  const renderCount = useRenderCounter();

  const table = useReactTable({
    data: abis.items || [],
    columns: AbisTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "abis";
  const tabs = ["abis"];
  const forms: ViewForm = {
    abis: <FormTable data={abis} groups={AbisFormTable(table)} />,
  };
  return (
    <ViewStateProvider route={route} nItems={abis.nItems} fetchFn={fetchAbis} modifyFn={ModifyAbi}>
      <Text>Render count: {renderCount}</Text>
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};
