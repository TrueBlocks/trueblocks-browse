import { DataTable, FieldGroup, CleanButton, AddButton } from "@components";
import { types } from "@gocode/models";

export const MonitorFormDef = (table: any): FieldGroup<types.MonitorContainer>[] => {
  return [
    {
      label: "Monitor Data",
      colSpan: 6,
      fields: [
        { label: "nMonitors", type: "int", accessor: "nItems" },
        { label: "nRecords", type: "int", accessor: "nRecords" },
        { label: "nNamed", type: "int", accessor: "nNamed" },
        { label: "fileSize", type: "bytes", accessor: "fileSize" },
      ],
    },
    {
      label: "Other",
      colSpan: 6,
      fields: [
        { label: "nEmpty", type: "int", accessor: "nEmpty" },
        { label: "nStaged", type: "int", accessor: "nStaged" },
        { label: "nDeleted", type: "int", accessor: "nDeleted" },
      ],
    },
    {
      label: "Buttons",
      buttons: [
        <AddButton key={"add"} value={"https://trueblocks.io"} />,
        <CleanButton key={"clean"} value={"https://trueblocks.io"} />,
      ],
    },
    {
      label: "Available Monitors",
      collapsable: false,
      components: [<DataTable<types.Monitor> key={"dataTable"} table={table} loading={false} />],
    },
  ];
};
