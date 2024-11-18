import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { DebugState, FormTable, View, ViewForm } from "@components";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { IndexesTableDef, IndexesFormDef } from ".";

export const IndexesView = () => {
  const { modifyNoop } = useNoops();
  const { indexes, fetchIndexes } = useAppState();

  const table = useReactTable({
    data: indexes.items || [],
    columns: IndexesTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "indexes";
  const tabs = ["indexes"];
  const forms: ViewForm = {
    indexes: <FormTable data={indexes} groups={IndexesFormDef(table)} />,
  };
  return (
    <ViewStateProvider route={route} nItems={indexes.nItems} fetchFn={fetchIndexes} modifyFn={modifyNoop}>
      <DebugState n={indexes.lastUpdate} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};
