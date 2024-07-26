import React from "react";

// Find: NewViews
import { IconHome, IconArticle, IconTag, IconServer, IconSettings, } from "@tabler/icons-react";
import { HomeView, HistoryView, NamesView, ServersView, SettingsView } from "@views";

// Note:
//  Change with care. The order of the items in this list matters (the last one is the default).
//  The order field is used to sort the menu items.
export const routeItems = [
  {
    order: 4,
    route: "/history/:address",
    label: "History",
    icon: <IconArticle />,
    component: HistoryView,
  },
  {
    order: 5,
    route: "/names",
    label: "Names",
    icon: <IconTag />,
    component: NamesView,
  },
  {
    order: 6,
    route: "/servers",
    label: "Servers",
    icon: <IconServer />,
    component: ServersView,
  },
  {
    order: 7,
    route: "/settings",
    label: "Settings",
    icon: <IconSettings />,
    component: SettingsView,
  },
  {
    order: 1,
    route: "/",
    label: "Home",
    icon: <IconHome />,
    component: HomeView,
  },
];
