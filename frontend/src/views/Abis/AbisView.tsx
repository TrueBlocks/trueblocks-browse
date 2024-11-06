import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm, DebugState } from "@components";
import { ModifyAbi } from "@gocode/app/App";
import { useAppState, ViewStateProvider } from "@state";
import { AbisFormTable, AbisTableDef } from ".";

export const AbisView = () => {
  const { abis, fetchAbis } = useAppState();

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
      <DebugState n={abis.lastUpdate} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};
