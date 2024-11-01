import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm } from "@components";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { StatusFormDef, StatusTableDef } from ".";

export const StatusView = () => {
  const { modifyNoop } = useNoops();
  const { status, fetchStatus } = useAppState();

  const table = useReactTable({
    data: status.items || [],
    columns: StatusTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "status";
  const tabs = ["status"];
  const forms: ViewForm = {
    status: <FormTable data={status} groups={StatusFormDef(table)} />,
  };
  return (
    <ViewStateProvider route={route} nItems={status.nItems} fetchFn={fetchStatus} modifyFn={modifyNoop}>
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};
