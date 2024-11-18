// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { FieldGroup, EditButton } from "@components";
import { types } from "@gocode/models";
// EXISTING_CODE

export const SessionFormDef = (session: types.SessionContainer): FieldGroup<types.Session>[] => {
  return [
    // EXISTING_CODE
    {
      label: "Last",
      colSpan: 6,
      collapsable: false,
      fields: [
        { label: "lastChain", type: "text", accessor: "lastChain" },
        { label: "lastFile", type: "text", accessor: "lastFile" },
        { label: "lastFolder", type: "text", accessor: "lastFolder" },
        { label: "lastRoute", type: "text", accessor: "lastRoute" },
        // { label: "lastSub", type: "text", accessor: "lastSub" },
      ],
    },
    {
      label: "Buttons",
      buttons: [<EditButton key={"add"} value={"https://trueblocks.io"} />],
    },
    {
      label: "Window",
      collapsable: false,
      components: [<SessionWindow key={"window"} session={session} />],
    },
    {
      label: "Toggles",
      collapsable: false,
      components: [<SessionToggles key={"window"} session={session} />],
    },
    // EXISTING_CODE
  ];
};

// EXISTING_CODE
type SessionProps = {
  session: types.SessionContainer;
};

const SessionWindow = ({ session }: SessionProps) => {
  return <div>{`${JSON.stringify(session.window, null, 2)}`}</div>;
};

const SessionToggles = ({ session }: SessionProps) => {
  return <div>{`${JSON.stringify(session.toggles, null, 2)}`}</div>;
};
// EXISTING_CODE

//-------------------------------------------------------------------
// Template variables:
// class:         Session
// routeLabel:    Session
// itemName:
// isHistory:     false
// isSession:     true
// isConfig:      false
