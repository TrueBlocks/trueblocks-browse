import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, DataTable, GroupDefinition } from "@components";
import { ModifyNoop } from "@gocode/app/App";
import { types } from "@gocode/models";
import { useAppState, ViewStateProvider } from "@state";
import { tableColumns } from "./ManifestTable";

export function ManifestView() {
  const { manifests, fetchManifests } = useAppState();

  const table = useReactTable({
    data: manifests.chunks || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <ViewStateProvider route={"manifest"} nItems={manifests.nItems} fetchFn={fetchManifests} modifyFn={ModifyNoop}>
      <View>
        <FormTable data={manifests} definition={createManifestForm(table)} />
      </View>
    </ViewStateProvider>
  );
}

type theInstance = InstanceType<typeof types.ManifestContainer>;
function createManifestForm(table: any): GroupDefinition<theInstance>[] {
  return [
    {
      title: "Manifest Data",
      colSpan: 6,
      fields: [
        { label: "version", type: "text", accessor: "version" },
        { label: "chain", type: "text", accessor: "chain" },
        { label: "specification", type: "hash", accessor: "specification" },
        { label: "lastUpdate", type: "date", accessor: "lastUpdate" },
      ],
    },
    {
      title: "Statistics",
      colSpan: 6,
      fields: [
        { label: "nBlooms", type: "int", accessor: "nBlooms" },
        { label: "bloomsSize", type: "bytes", accessor: "bloomsSize" },
        { label: "nIndexes", type: "int", accessor: "nIndexes" },
        { label: "indexSize", type: "bytes", accessor: "indexSize" },
      ],
    },
    {
      title: "Chunks",
      fields: [],
      components: [
        {
          component: <DataTable<types.ChunkRecord> table={table} loading={false} />,
        },
      ],
    },
  ];
}
