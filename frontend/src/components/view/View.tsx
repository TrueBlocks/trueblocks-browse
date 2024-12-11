import { ReactNode, useEffect, useState } from "react";
import { Group, Stack, Tabs } from "@mantine/core";
import { SearchBar, ViewTitle } from "@components";
import { GetTabs } from "@gocode/app/App";
import { useAppState } from "@state";
import classes from "./View.module.css";

export type ViewForm = Record<string, ReactNode>;

export type ViewProps = {
  tabItems: ViewForm;
  searchable?: boolean;
};

export const View = ({ tabItems, searchable = false }: ViewProps) => {
  const { activeTab, tabChanged } = useAppState();
  const [tabs, setTabs] = useState<string[] | null>(null);

  useEffect(() => {
    GetTabs().then((tts) => {
      setTabs(tts);
    });
  }, []);

  return (
    <Stack className={classes.viewContainer}>
      <Group style={{ justifyContent: "space-between", alignItems: "center" }}>
        <ViewTitle />
        {searchable && (
          <div style={{ width: "30%" }}>
            <SearchBar />
          </div>
        )}
      </Group>
      <Tabs value={activeTab} onChange={(newVal) => tabChanged(newVal || "")}>
        <Tabs.List>
          {tabs?.map((tab: string) => (
            <Tabs.Tab key={tab} className={classes.tab} value={tab}>
              {toProperCase(tab)}
            </Tabs.Tab>
          ))}
        </Tabs.List>
        {tabs?.map((tab: string) => (
          <Tabs.Panel key={tab} value={tab}>
            {tabItems[tab]}
          </Tabs.Panel>
        ))}
      </Tabs>
    </Stack>
  );
};

// TODO: Move to utils
function toProperCase(str: string): string {
  return str.replace(/\w\S*/g, (txt) => {
    return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();
  });
}
