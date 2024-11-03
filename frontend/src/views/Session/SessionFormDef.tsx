import { FieldGroup, EditButton } from "@components";
import { types } from "@gocode/models";
import { useAppState } from "@state";

export const SessionFormDef = (): FieldGroup<types.Session>[] => {
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
      components: [<SessionWindow key={"window"} />],
    },
    {
      label: "Wizard",
      collapsable: false,
      components: [<SessionWizard key={"window"} />],
    },
    {
      label: "Toggles",
      collapsable: false,
      components: [<SessionToggles key={"window"} />],
    },
    /*
      LastSub: string;
 */
  ];
};

const SessionWindow = () => {
  const { session } = useAppState();
  return <div>{`${JSON.stringify(session.window, null, 2)}`}</div>;
};

const SessionWizard = () => {
  const { session } = useAppState();
  return <div>{`${JSON.stringify(session.wizard, null, 2)}`}</div>;
};

const SessionToggles = () => {
  const { session } = useAppState();
  return <div>{`${JSON.stringify(session.toggles, null, 2)}`}</div>;
};
