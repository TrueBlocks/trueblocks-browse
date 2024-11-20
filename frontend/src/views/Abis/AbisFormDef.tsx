// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { Table } from "@tanstack/react-table";
import { DataTable, FieldGroup, PublishButton, CleanButton, AddButton } from "@components";
import { types } from "@gocode/models";
// EXISTING_CODE

export const AbisFormDef = (table: Table<types.Abi>): FieldGroup<types.AbiContainer>[] => {
  return [
    // EXISTING_CODE
    {
      label: "Abi Data",
      colSpan: 6,
      fields: [
        { label: "nItems", type: "int", accessor: "nItems" },
        { label: "nFunctions", type: "int", accessor: "nFunctions" },
        { label: "nEvents", type: "int", accessor: "nEvents" },
        { label: "fileSize", type: "bytes", accessor: "fileSize" },
      ],
    },
    {
      label: "Bounds",
      colSpan: 6,
      fields: [
        { label: "largestFile", type: "text", accessor: "largestFile" },
        { label: "mostFunctions", type: "text", accessor: "mostFunctions" },
        { label: "mostEvents", type: "text", accessor: "mostEvents" },
      ],
    },
    {
      label: "Buttons",
      buttons: [
        <AddButton key={"add"} value={"https://trueblocks.io"} />,
        <CleanButton key={"clean"} value={"https://trueblocks.io"} />,
        <PublishButton key={"publish"} value={"https://trueblocks.io"} />,
      ],
    },
    {
      label: "Files",
      collapsable: false,
      components: [<DataTable<types.Abi> key={"dataTable"} table={table} loading={false} />],
    },
    // EXISTING_CODE
  ];
};

// EXISTING_CODE
// EXISTING_CODE
