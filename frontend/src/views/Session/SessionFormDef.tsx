import { FieldGroup, EditButton } from "@components";
import { types } from "@gocode/models";

export const SessionFormDef = (session: types.SessionContainer): FieldGroup<types.Session>[] => {
  return [
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
    /*
      LastSub: string;
 */
  ];
};

type SessionProps = {
  session: types.SessionContainer;
};

const SessionWindow = ({ session }: SessionProps) => {
  return <div>{`${JSON.stringify(session.window, null, 2)}`}</div>;
};

const SessionToggles = ({ session }: SessionProps) => {
  return <div>{`${JSON.stringify(session.toggles, null, 2)}`}</div>;
};
