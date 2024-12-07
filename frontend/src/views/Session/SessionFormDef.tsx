// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { Table } from "@tanstack/react-table";
import { FieldGroup, EditButton, DataTable } from "@components";
import { types } from "@gocode/models";
import { GetChain, GetFile, GetFolder, GetRoute, GetAddress } from "@gocode/types/SessionContainer";
// EXISTING_CODE

export const SessionFormDef = (table: Table<types.Nothing>): FieldGroup<types.SessionContainer>[] => {
  // EXISTING_CODE
  // EXISTING_CODE
  return [
    // EXISTING_CODE
    {
      label: "Last",
      colSpan: 6,
      collapsable: false,
      fields: [
        { label: "lastChain", type: "text", getter: () => GetChain() },
        { label: "lastFile", type: "text", getter: () => GetFile() },
        { label: "lastFolder", type: "text", getter: () => GetFolder() },
        { label: "lastRoute", type: "text", getter: () => GetRoute() },
        { label: "lastAddress", type: "text", getter: () => GetAddress() },
      ],
    },
    {
      label: "Buttons",
      buttons: [<EditButton key={"add"} value={"https://trueblocks.io"} />],
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
      label: "Nothing",
      collapsable: false,
      components: [<DataTable<types.Nothing> key={"dataTable"} table={table} loading={false} />],
    },
    // EXISTING_CODE
  ];
};

// EXISTING_CODE
const SessionWindow = () => {
  // const { session } = useAppState();
  return <div>CANNOT</div>; // {`${JSON.stringify(session.window, null, 2)}`}</div>;
};

const SessionFlags = () => {
  // const { session } = useAppState();
  return <div>CANNOT</div>; //  <div>{`${JSON.stringify(session.flags, null, 2)}`}</div>;
};
// EXISTING_CODE
