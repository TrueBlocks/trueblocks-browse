import { DataTable, FieldGroup, CleanButton, PublishButton, AddButton } from "@components";
import { types } from "@gocode/models";

export const NamesFormDef = (table: any): FieldGroup<types.NameContainer>[] => {
  return [
    {
      label: "Name Data",
      colSpan: 6,
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
  ];
};
