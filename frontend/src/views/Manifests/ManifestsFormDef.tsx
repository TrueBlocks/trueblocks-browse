import { DataTable, FieldGroup, SpecButton, PublishButton } from "@components";
import { types } from "@gocode/models";

export const ManifestsFormDef = (table: any): FieldGroup<types.ManifestContainer>[] => {
  return [
    {
      label: "Manifest Data",
      colSpan: 6,
      fields: [
        { label: "version", type: "text", accessor: "version" },
        { label: "chain", type: "text", accessor: "chain" },
        { label: "specification", type: "hash", accessor: "specification" },
        { label: "lastUpdate", type: "date", accessor: "lastUpdate" },
      ],
    },
    {
      label: "Statistics",
      colSpan: 6,
      fields: [
        { label: "nBlooms", type: "int", accessor: "nBlooms" },
        { label: "bloomsSize", type: "bytes", accessor: "bloomsSize" },
        { label: "nIndexes", type: "int", accessor: "nIndexes" },
        { label: "indexSize", type: "bytes", accessor: "indexSize" },
      ],
    },
    {
      label: "Buttons",
      buttons: [
        <PublishButton key={"publish"} value="https://trueblocks.io" />,
        <SpecButton
          key={"spec"}
          value="https://trueblocks.io/papers/2023/specification-for-the-unchained-index-v2.0.0-release.pdf"
        />,
      ],
    },
    {
      label: "Chunks",
      collapsable: false,
      components: [<DataTable<types.ChunkRecord> key={"dataTable"} table={table} loading={false} />],
    },
  ];
};
