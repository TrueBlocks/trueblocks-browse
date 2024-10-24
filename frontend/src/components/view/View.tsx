import { ReactNode, useState, useEffect } from "react";
import { Stack, Tabs } from "@mantine/core";
import { ViewTitle } from "@components";
import { messages } from "@gocode/models";
import { EventsOn, EventsOff } from "@runtime";
import classes from "./View.module.css";

export type ViewProps = {
  tabs: string[];
  forms: Record<string, ReactNode>;
};

export const View = ({ tabs, forms }: ViewProps) => {
  const [activeTab, setActiveTab] = useState<string>(tabs[0]);

  useEffect(() => {
    const handleSwitchTab = (msg: messages.MessageMsg) => {
      const { string1 } = msg;
      switch (string1) {
        case "next":
          setActiveTab((prevTab) => {
            const currentIndex = tabs.indexOf(prevTab);
            return currentIndex > 0 ? tabs[currentIndex - 1] : tabs[tabs.length - 1];
          });
          break;
        case "prev":
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
  }, [tabs]);

  if (!tabs || tabs.length < 1) {
    return <>Loading...</>;
  }

  // const keys = Object.keys(forms);
  // const types = Object.values(forms);
  // const tt = types.map((t) => typeof t);

  return (
    <Stack className={classes.viewContainer}>
      <ViewTitle />
      <Tabs
        value={activeTab}
        onChange={(t) => {
          if (t !== null) {
            setActiveTab(t);
          }
        }}
      >
        <Tabs.List>
          {tabs?.map((tab) => (
            <Tabs.Tab key={tab} className={classes.tab} value={tab}>
              {toProperCase(tab)}
            </Tabs.Tab>
          ))}
        </Tabs.List>
        {tabs?.map((tab) => (
          <Tabs.Panel key={tab} value={tab}>
            {forms ? forms[tab] : null}
          </Tabs.Panel>
        ))}
      </Tabs>
    </Stack>
  );
};

function toProperCase(str: string): string {
  return str.replace(/\w\S*/g, (txt) => {
    return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();
  });
}
