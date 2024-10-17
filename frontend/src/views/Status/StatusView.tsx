import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, DataTable, GroupDefinition } from "@components";
import { ModifyNoop } from "@gocode/app/App";
import { types } from "@gocode/models";
import { useAppState, ViewStateProvider } from "@state";
import { tableColumns } from "./StatusTable";

export const StatusView = () => {
  const { status, fetchStatus } = useAppState();

  const table = useReactTable({
    data: status.caches || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <ViewStateProvider route={"status"} nItems={status.nItems} fetchFn={fetchStatus} modifyFn={ModifyNoop}>
      <View>
        <FormTable data={status} definition={createStatusForm(table)} />
      </View>
    </ViewStateProvider>
  );
};

type theInstance = InstanceType<typeof types.StatusContainer>;
const createStatusForm = (table: any): GroupDefinition<theInstance>[] => {
  return [
    {
      legend: "System Data",
      colSpan: 7,
      fields: [
        { label: "trueblocks", type: "text", accessor: "version" },
        { label: "client", type: "text", accessor: "clientVersion" },
        { label: "isArchive", type: "boolean", accessor: "isArchive" },
        { label: "isTracing", type: "boolean", accessor: "isTracing" },
      ],
    },
    {
      legend: "API Keys",
      colSpan: 5,
      fields: [
        { label: "hasEsKey", type: "boolean", accessor: "hasEsKey" },
        { label: "hasPinKey", type: "boolean", accessor: "hasPinKey" },
        { label: "rpcProvider", type: "url", accessor: "rpcProvider" },
      ],
    },
    {
      legend: "Configuration Paths",
      colSpan: 7,
      fields: [
        { label: "rootConfig", type: "path", accessor: "rootConfig" },
        { label: "chainConfig", type: "path", accessor: "chainConfig" },
        { label: "indexPath", type: "path", accessor: "indexPath" },
        { label: "cachePath", type: "path", accessor: "cachePath" },
      ],
    },
    {
      legend: "Statistics",
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
      legend: "Caches",
      fields: [],
      components: [
        {
          component: <DataTable<types.CacheItem> table={table} loading={false} />,
        },
      ],
    },
  ];
};
