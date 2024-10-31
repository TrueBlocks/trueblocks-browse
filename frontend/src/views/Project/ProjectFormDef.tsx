import { DataTable, FieldGroup, AddButton } from "@components";
import { types } from "@gocode/models";

export const ProjectFormDef = (table: any): FieldGroup<types.ProjectContainer>[] => {
  return [
    {
      label: "Data 1",
      colSpan: 4,
      fields: [
        // { label: "fileName", type: "text", accessor: "filename" },
        { label: "nHistories", type: "int", accessor: "nItems" },
        { label: "historySize", type: "bytes", accessor: "historySize" },
        // { label: "dirty", type: "boolean", accessor: "dirty" },
      ],
    },
    {
      label: "Data 2",
      colSpan: 4,
      fields: [
        { label: "nNames", type: "int", accessor: "nNames" },
        { label: "nAbis", type: "int", accessor: "nAbis" },
        { label: "nCaches", type: "int", accessor: "nCaches" },
      ],
    },
    {
      label: "Data 2",
      colSpan: 4,
      fields: [
        { label: "nMonitors", type: "int", accessor: "nMonitors" },
        { label: "nIndexes", type: "int", accessor: "nIndexes" },
        { label: "nManifests", type: "int", accessor: "nManifests" },
      ],
    },
    {
      label: "Buttons",
      buttons: [<AddButton key={"add"} value={"https://trueblocks.io"} />],
    },
    {
      label: "Histories",
      fields: [],
      collapsable: false,
      components: [<DataTable<types.HistoryContainer> key={"dataTable"} table={table} loading={false} />],
    },
  ];
};
