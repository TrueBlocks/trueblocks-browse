import { useEffect, useState } from "react";
import { useLocation } from "wouter";
import { StyledNavLink } from "@components";
import { GetRoute, GetAddress, SetRoute } from "@gocode/app/App";
import { messages } from "@gocode/models";
import { routeItems, RouteItem } from "@layout";
import { EventsOn, EventsOff } from "@runtime";
import { useAppState } from "@state";

export const Menu = () => {
  const [activeRoute, setActiveRoute] = useState("/");
  const [_, setLocation] = useLocation();
  const [filteredMenu, setFilteredMenu] = useState<RouteItem[]>([]);
  const { isConfigured } = useAppState();

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
    const handleNavigation = (msg: messages.NavigateMsg) => {
      setLocation(msg.route);
      setActiveRoute(msg.route);
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
      GetAddress().then((address) => {
        const addr = address as unknown as string;
        route = route.replace(":address", addr);
        setLocation(route);
        SetRoute(route, addr);
      });
      setActiveRoute("/history/:address");
    } else {
      SetRoute(route, "");
      setActiveRoute(route);
    }
  };

  useEffect(() => {
    setFilteredMenu(
      routeItems
        .filter((item: RouteItem) => (isConfigured ? item.route !== "/wizard" : item.route === "/wizard"))
        .sort((a, b) => a.order - b.order)
    );
  }, [isConfigured]);

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