import { ReactNode } from "react";
import { Group, Stack, Tabs } from "@mantine/core";
import { SearchBar, ViewTitle } from "@components";
import { SetLastTab } from "@gocode/app/App";
import { useViewState } from "@state";
import classes from "./View.module.css";

export type ViewForm = Record<string, ReactNode>;

export type ViewProps = {
  forms: ViewForm;
  searchable?: boolean;
};

export const View = ({ forms, searchable = false }: ViewProps) => {
  const { route, tabs, activeTab, setActiveTab } = useViewState();

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
      <Tabs
        value={activeTab}
        onChange={(newTab) => {
          if (newTab !== null) {
            setActiveTab(newTab);
            SetLastTab(route, newTab);
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

// TODO: Move to utils
function toProperCase(str: string): string {
  return str.replace(/\w\S*/g, (txt) => {
    return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();
  });
}
