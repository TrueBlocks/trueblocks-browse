// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { Table } from "@tanstack/react-table";
import { DataTable, FieldGroup, CleanButton, PublishButton, AddButton, NameEditor } from "@components";
import { types } from "@gocode/models";
// EXISTING_CODE

export const NamesFormDef = (editable: types.Name, table: Table<types.Name>): FieldGroup<types.NameContainer>[] => {
  // EXISTING_CODE
  // EXISTING_CODE
  return [
    // EXISTING_CODE
    {
      label: "",
      collapsable: false,
      disabled: editable.name === "",
      components: [<NameEditor key={"nameEditor"} source={editable} onCancel={() => {}} />],
    },
    {
      label: "Name Data",
      colSpan: 6,
      disabled: editable.name !== "",
      fields: [
        { label: "nNames", type: "int", accessor: "nItems" },
        { label: "nContracts", type: "int", accessor: "nContracts" },
        { label: "nErc20s", type: "int", accessor: "nErc20s" },
        { label: "nErc721s", type: "int", accessor: "nErc721s" },
        { label: "nDeleted", type: "int", accessor: "nDeleted" },
      ],
    },
    {
      label: "Database Parts",
      colSpan: 6,
      fields: [
        { label: "sizeOnDisc", type: "bytes", accessor: "sizeOnDisc" },
        { label: "nCustom", type: "int", accessor: "nCustom" },
        { label: "nRegular", type: "int", accessor: "nRegular" },
        { label: "nPrefund", type: "int", accessor: "nPrefund" },
        { label: "nSystem", type: "int", accessor: "nSystem" },
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
      label: "Names",
      collapsable: false,
      components: [<DataTable<types.Name> key={"dataTable"} table={table} loading={false} />],
    },
    // EXISTING_CODE
  ];
};

// EXISTING_CODE
// EXISTING_CODE
