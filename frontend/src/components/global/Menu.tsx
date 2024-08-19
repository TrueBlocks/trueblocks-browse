import React, { useEffect, useState } from "react";
import { GetLast, GetLastSub, SetLast } from "@gocode/app/App";
// TODO: This alias is wrong, can it not be @Routes See also @/App.module.css
import { routeItems, RouteItem } from "@/Routes";
import { useLocation } from "wouter";
import { messages } from "@gocode/models";
import { EventsOn, EventsOff } from "@runtime";
import { useAppState } from "@state";
import { StyledNavLink } from ".";

export function Menu() {
  const [activeRoute, setActiveRoute] = useState("/");
  const [_, setLocation] = useLocation();
  const { isConfigured } = useAppState();

  useEffect(() => {
    (GetLast("route") || "/").then((route) => {
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
      GetLastSub("/history").then((subRoute) => {
        route = route.replace("/:address", subRoute);
        setLocation(route);
      });
      SetLast("route", route);
      setActiveRoute("/history/:address");
    } else {
      SetLast("route", route);
      setActiveRoute(route);
    }
  };

  const routes = routeItems
    .filter((item: RouteItem) => (isConfigured ? item.route !== "/wizard" : item.route === "/wizard"))
    .sort((a, b) => a.order - b.order);

  return (
    <div style={{ flexGrow: 1 }}>
      {routes.map((item) => {
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
}
