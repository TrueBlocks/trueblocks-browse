import { Text } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { FormTable, View, ViewForm } from "@components";
import { useNoops, useRenderCounter } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { IndexesTableDef, IndexedFormDef } from ".";

export const IndexesView = () => {
  const { modifyNoop } = useNoops();
  const { indexes, fetchIndexes } = useAppState();
  const renderCount = useRenderCounter();

  const table = useReactTable({
    data: indexes.items || [],
    columns: IndexesTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "indexes";
  const tabs = ["indexes"];
  const forms: ViewForm = {
    indexes: <FormTable data={indexes} groups={IndexedFormDef(table)} />,
  };
  return (
    <ViewStateProvider route={route} nItems={indexes.nItems} fetchFn={fetchIndexes} modifyFn={modifyNoop}>
      <Text>Render count: {renderCount}</Text>
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};
