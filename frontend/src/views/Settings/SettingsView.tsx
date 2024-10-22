import { useEffect, useState } from "react";
import { Tabs } from "@mantine/core";
import { FieldGroup, FormTable, PublishButton, SpecButton, View } from "@components";
import { config as browseConfig, configtypes, messages } from "@gocode/models";
import { EventsOn, EventsOff } from "@runtime";
import { ViewStateProvider, useAppState } from "@state";
import { useNoops } from "../../hooks";
import classes from "./SettingsView.module.css";

export const SettingsView = () => {
  const { modifyNoop } = useNoops();
  const { settings, fetchSettings } = useAppState();
  const [activeTab, setActiveTab] = useState("session");

  useEffect(() => {
    const tabs = ["session", "config"];
    const handleSwitchTab = (msg: messages.SwitchTabMsg) => {
      const { dest } = msg;
      switch (dest) {
        case "prev":
          setActiveTab((prevTab) => {
            const currentIndex = tabs.indexOf(prevTab);
            return currentIndex > 0 ? tabs[currentIndex - 1] : tabs[tabs.length - 1];
          });
          break;
        case "next":
          setActiveTab((prevTab) => {
            const currentIndex = tabs.indexOf(prevTab);
            return currentIndex < tabs.length - 1 ? tabs[currentIndex + 1] : tabs[0];
          });
          break;
        default:
          break;
      }
    };

    const { Message } = messages;
    EventsOn(Message.SWITCHTAB, handleSwitchTab);
    return () => {
      EventsOff(Message.SWITCHTAB);
    };
  }, []);

  if (!settings) {
    return <div>Loading...</div>;
  }

  const session = settings.session ?? browseConfig.Session.createFrom({});
  const config = settings.config ?? configtypes.Config.createFrom({});

  const route = "settings";
  return (
    <ViewStateProvider route={route} fetchFn={fetchSettings} modifyFn={modifyNoop}>
      <View>
        <Tabs
          value={activeTab}
          onChange={(value) => {
            if (value !== null) {
              setActiveTab(value);
            }
          }}
        >
          <Tabs.List>
            <Tabs.Tab className={classes.tab} value="session">
              Session
            </Tabs.Tab>
            <Tabs.Tab className={classes.tab} value="config">
              Config
            </Tabs.Tab>
          </Tabs.List>

          <Tabs.Panel value="session">
            <FormTable data={session} groups={createSessionForm()} />
          </Tabs.Panel>

          <Tabs.Panel value="config">
            <FormTable data={config} groups={createConfigForm()} />
          </Tabs.Panel>
        </Tabs>
      </View>
    </ViewStateProvider>
  );
};

const createSessionForm = (): FieldGroup<browseConfig.Session>[] => {
  return [
    {
      label: "Session Data 1",
      colSpan: 6,
      fields: [
        { label: "chain", type: "text", accessor: "chain" },
        { label: "lastFile", type: "text", accessor: "lastFile" },
        { label: "lastRoute", type: "text", accessor: "lastRoute" },
      ],
    },
    {
      label: "Session Data 2",
      colSpan: 6,
      collapsable: false,
      fields: [
        // { label: "lastSub", type: "text", accessor: "lastSub" },
        // { label: "window", type: "text", accessor: "window" },
        // { label: "daemons", type: "text", accessor: "lastRoute" },
        // { label: "wizard", type: "text", accessor: "lastRoute" },
        // { label: "toggles", type: "text", accessor: "lastRoute" },
      ],
    },
  ];
};

const createConfigForm = (): FieldGroup<configtypes.Config>[] => {
  return [
    {
      label: "Buttons",
      buttons: [
        <PublishButton key="publish" value="https://trueblocks.io" />,
        <SpecButton
          key="spec"
          value="https://trueblocks.io/papers/2023/specification-for-the-unchained-index-v2.0.0-release.pdf"
        />,
      ],
    },
    // {
    //   label: "Version Group",
    //   colSpan: 12, // Full width since it's the only field
    //   components: [<div key={"1"}>Hello world</div>],
    // },
  ];
};
