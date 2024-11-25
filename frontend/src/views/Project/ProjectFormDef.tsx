// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { Table } from "@tanstack/react-table";
import { DataTable, FieldGroup, AddButton } from "@components";
import { types } from "@gocode/models";
import { useAppState } from "@state";
// EXISTING_CODE

export const ProjectFormDef = (table: Table<types.HistoryContainer>): FieldGroup<types.ProjectContainer>[] => {
  // EXISTING_CODE
  const { info } = useAppState();
  // EXISTING_CODE
  return [
    // EXISTING_CODE
    {
      label: "Data 1",
      colSpan: 4,
      fields: [
        // { label: "fileName", type: "text", accessor: "lastFile" },
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
      label: info.filename,
      fields: [],
      collapsable: false,
      components: [<DataTable<types.HistoryContainer> key={"dataTable"} table={table} loading={false} />],
    },
    // EXISTING_CODE
  ];
};

// EXISTING_CODE
// EXISTING_CODE
