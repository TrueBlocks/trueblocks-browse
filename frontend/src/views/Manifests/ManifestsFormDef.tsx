// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { Table } from "@tanstack/react-table";
import { DataTable, FieldGroup, SpecButton, PublishButton } from "@components";
import { types } from "@gocode/models";
// EXISTING_CODE

export const ManifestsFormDef = (table: Table<types.ChunkRecord>): FieldGroup<types.ManifestContainer>[] => {
  return [
    // EXISTING_CODE
    {
      label: "Manifest Data",
      colSpan: 6,
      fields: [
        { label: "version", type: "text", accessor: "version" },
        { label: "chain", type: "text", accessor: "chain" },
        { label: "specification", type: "hash", accessor: "specification" },
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
    // EXISTING_CODE
  ];
};

// EXISTING_CODE
// EXISTING_CODE
