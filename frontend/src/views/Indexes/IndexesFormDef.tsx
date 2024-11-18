// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { Table } from "@tanstack/react-table";
import { DataTable, FieldGroup, SpecButton, PinButton } from "@components";
import { types } from "@gocode/models";
// EXISTING_CODE

export const IndexesFormDef = (table: Table<types.ChunkStats>): FieldGroup<types.IndexContainer>[] => {
  return [
    // EXISTING_CODE
    {
      label: "Index Data",
      colSpan: 6,
      fields: [
        { label: "bloomSz", type: "bytes", accessor: "bloomSz" },
        { label: "chunkSz", type: "bytes", accessor: "chunkSz" },
        { label: "nAddrs", type: "int", accessor: "nAddrs" },
        { label: "nApps", type: "int", accessor: "nApps" },
        { label: "nBlocks", type: "int", accessor: "nBlocks" },
        { label: "nBlooms", type: "int", accessor: "nBlooms" },
      ],
    },
    {
      label: "Statistics",
      colSpan: 6,
      fields: [
        { label: "nItems", type: "int", accessor: "nItems" },
        { label: "addrsPerBlock", type: "float", accessor: "addrsPerBlock" },
        { label: "appsPerAddr", type: "float", accessor: "appsPerAddr" },
        { label: "appsPerBlock", type: "float", accessor: "appsPerBlock" },
      ],
    },
    {
      label: "Buttons",
      buttons: [
        <PinButton key={"pin"} value="https://trueblocks.io" />,
        <SpecButton
          key={"spec"}
          value="https://trueblocks.io/papers/2023/specification-for-the-unchained-index-v2.0.0-release.pdf"
        />,
      ],
    },
    {
      label: "Chunks",
      collapsable: false,
      components: [<DataTable<types.ChunkStats> key={"dataTable"} table={table} loading={false} />],
    },
    // EXISTING_CODE
  ];
};

// EXISTING_CODE
// EXISTING_CODE

//-------------------------------------------------------------------
// Template variables:
// class:         Index
// routeLabel:    Indexes
// itemName:      ChunkStats
// isHistory:     false
// isSession:     false
// isConfig:      false
