import { useEffect, useState } from "react";
import { useLocation } from "wouter";
import { StyledNavLink } from "@components";
import { GetRoute, GetLastAddress, SetRoute } from "@gocode/app/App";
import { messages, types } from "@gocode/models";
import { routeItems, RouteItem } from "@layout";
import { EventsOn, EventsOff } from "@runtime";
import { useAppState } from "@state";

export const Menu = () => {
  const [activeRoute, setActiveRoute] = useState("/");
  const [, setLocation] = useLocation();
  const [filteredMenu, setFilteredMenu] = useState<RouteItem[]>([]);
  const { wizard } = useAppState();

  useEffect(() => {
    GetRoute().then((route) => {
      if (route.startsWith("/history")) {
        setActiveRoute("/history/:address");
      } else {
        setActiveRoute(route);
      }
    });
  }, []);

  useEffect(() => {
    const handleNavigation = (msg: messages.MessageMsg) => {
      setLocation(msg.string1);
      setActiveRoute(msg.string1);
    };

    const { Message } = messages;
    EventsOn(Message.NAVIGATE, handleNavigation);
    return () => {
      EventsOff(Message.NAVIGATE);
    };
  }, [setLocation]);

  const handleRouteChange = (route: string) => {
    setActiveRoute(route);
    if (route.startsWith("/history")) {
      GetLastAddress().then((address) => {
        const addr = address as unknown as string;
        route = route.replace(":address", addr);
        setLocation(route);
        SetRoute("/history", addr, ""); // TODO: Put active Tab in App state so we can use it here
      });
      setActiveRoute("/history/:address");
    } else {
      SetRoute(route, "", ""); // TODO: Put active Tab in App state so we can use it here
      setActiveRoute(route);
    }
  };

  useEffect(() => {
    setFilteredMenu(
      routeItems
        .filter((item: RouteItem) =>
          wizard.state === types.WizState.FINISHED ? item.route !== "/wizard" : item.route === "/wizard"
        )
        .sort((a, b) => a.order - b.order)
    );
  }, [wizard.state]);

  return (
    <div style={{ flexGrow: 1 }}>
      {filteredMenu.map((item) => {
        return (
          <StyledNavLink
            key={item.route}
            label={item.label}
            icon={item.icon}
            href={item.route}
            onClick={() => handleRouteChange(item.route)}
            activeRoute={activeRoute}
          />
        );
      })}
    </div>
  );
};
