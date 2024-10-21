import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, DataTable, FieldGroup, CleanButton } from "@components";
import { types } from "@gocode/models";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { tableColumns } from "./StatusTable";

export const StatusView = () => {
  const { modifyNoop } = useNoops();
  const { status, fetchStatus } = useAppState();

  const table = useReactTable({
    data: status.caches || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "status";
  return (
    <ViewStateProvider route={route} nItems={status.nItems} fetchFn={fetchStatus} modifyFn={modifyNoop}>
      <View>
        <FormTable data={status} groups={createStatusForm(table)} />
      </View>
    </ViewStateProvider>
  );
};

const createStatusForm = (table: any): FieldGroup<types.StatusContainer>[] => {
  return [
    {
      label: "System Data",
      colSpan: 7,
      fields: [
        { label: "trueblocks", type: "text", accessor: "version" },
        { label: "client", type: "text", accessor: "clientVersion" },
        { label: "isArchive", type: "boolean", accessor: "isArchive" },
        { label: "isTracing", type: "boolean", accessor: "isTracing" },
      ],
    },
    {
      label: "API Keys",
      colSpan: 5,
      fields: [
        { label: "hasEsKey", type: "boolean", accessor: "hasEsKey" },
        { label: "hasPinKey", type: "boolean", accessor: "hasPinKey" },
        { label: "rpcProvider", type: "url", accessor: "rpcProvider" },
      ],
    },
    {
      label: "Configuration Paths",
      colSpan: 7,
      fields: [
        { label: "rootConfig", type: "path", accessor: "rootConfig" },
        { label: "chainConfig", type: "path", accessor: "chainConfig" },
        { label: "indexPath", type: "path", accessor: "indexPath" },
        { label: "cachePath", type: "path", accessor: "cachePath" },
      ],
    },
    {
      label: "Statistics",
      colSpan: 5,
      fields: [
        { label: "lastUpdate", type: "date", accessor: "lastUpdate" },
        { label: "nCaches", type: "int", accessor: "nItems" },
        { label: "nFiles", type: "int", accessor: "nFiles" },
        { label: "nFolders", type: "int", accessor: "nFolders" },
        { label: "sizeInBytes", type: "bytes", accessor: "nBytes" },
      ],
    },
    {
      label: "Buttons",
      buttons: [<CleanButton key={"clean"} value={"https://trueblocks.io"} />],
    },
    {
      label: "Caches",
      fields: [],
      collapsable: false,
      components: [<DataTable<types.CacheItem> key={"dataTable"} table={table} loading={false} />],
    },
  ];
};
