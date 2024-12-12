// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { Table } from "@tanstack/react-table";
import { FieldGroup, EditButton } from "@components";
import { types } from "@gocode/models";
import { useAppState } from "@state";
// EXISTING_CODE

export const SessionFormDef = (table: Table<types.Nothing>): FieldGroup<types.SessionContainer>[] => {
  // EXISTING_CODE
  // EXISTING_CODE
  return [
    // EXISTING_CODE
    {
      label: "Last",
      fields: [
        { label: "lastChain", type: "text", accessor: "lastChain" },
        { label: "lastFile", type: "text", accessor: "lastFile" },
        { label: "lastFolder", type: "text", accessor: "lastFolder" },
        { label: "lastRoute", type: "text", accessor: "lastRoute" },
        { label: "lastAddress", type: "text", accessor: "lastAddress" },
      ],
    },
    {
      label: "Window",
      collapsable: false,
      components: [<SessionWindow key={"window"} />],
    },
    {
      label: "Flags",
      collapsable: false,
      components: [<SessionFlags key={"flags"} />],
    },
    {
      label: "Buttons",
      buttons: [<EditButton key={"add"} value={"https://trueblocks.io"} />],
    },
    // {
    //   label: "Nothing",
    //   collapsable: false,
    //   components: [<DataTable<types.Nothing> key={"dataTable"} table={table} loading={false} />],
    // },
    // EXISTING_CODE
  ];
};

// EXISTING_CODE
const SessionWindow = () => {
  const { session } = useAppState();
  return <div>{`${JSON.stringify(session.window, null, 2)}`}</div>;
};

const SessionFlags = () => {
  const { session } = useAppState();
  return <div>{`${JSON.stringify(session.flags, null, 2)}`}</div>;
};
// EXISTING_CODE
