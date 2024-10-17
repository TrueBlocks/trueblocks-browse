import { Stack } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, DataTable, GroupDefinition, SpecButton, PublishButton } from "@components";
import { ModifyNoop } from "@gocode/app/App";
import { types } from "@gocode/models";
import { useAppState, ViewStateProvider } from "@state";
import { tableColumns } from "./ManifestsTable";

export const ManifestsView = () => {
  const { manifests, fetchManifests } = useAppState();

  const table = useReactTable({
    data: manifests.chunks || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <ViewStateProvider route={"manifests"} nItems={manifests.nItems} fetchFn={fetchManifests} modifyFn={ModifyNoop}>
      <View>
        <FormTable data={manifests} definition={createManifestForm(table)} />
      </View>
    </ViewStateProvider>
  );
};

type theInstance = InstanceType<typeof types.ManifestContainer>;
const createManifestForm = (table: any): GroupDefinition<theInstance>[] => {
  return [
    {
      legend: "Manifest Data",
      colSpan: 5,
      fields: [
        { label: "version", type: "text", accessor: "version" },
        { label: "chain", type: "text", accessor: "chain" },
        { label: "specification", type: "hash", accessor: "specification" },
        { label: "lastUpdate", type: "date", accessor: "lastUpdate" },
      ],
    },
    {
      legend: "Statistics",
      colSpan: 5,
      fields: [
        { label: "nBlooms", type: "int", accessor: "nBlooms" },
        { label: "bloomsSize", type: "bytes", accessor: "bloomsSize" },
        { label: "nIndexes", type: "int", accessor: "nIndexes" },
        { label: "indexSize", type: "bytes", accessor: "indexSize" },
      ],
    },
    {
      legend: "Buttons",
      colSpan: 2,
      components: [
        {
          component: (
            <Stack>
              <PublishButton value="https://trueblocks.io">Publish</PublishButton>
              <SpecButton value="https://trueblocks.io/papers/2023/specification-for-the-unchained-index-v2.0.0-release.pdf">
                Spec
              </SpecButton>
            </Stack>
          ),
        },
      ],
    },
    {
      legend: "Chunks",
      components: [
        {
          component: <DataTable<types.ChunkRecord> table={table} loading={false} />,
        },
      ],
    },
  ];
};
