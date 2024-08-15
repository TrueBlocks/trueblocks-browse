import React, { ReactNode, useEffect, useState } from "react";
import { NavLink } from "@mantine/core";
import { Link, useRoute } from "wouter";
import { GetLast, GetLastSub, SetLast } from "@gocode/app/App";
// TODO: This alias is wrong, can it not be @Routes See also @/App.module.css
import { routeItems } from "@/Routes";
import { useLocation } from "wouter";
import { messages } from "@gocode/models";
import { EventsOn, EventsOff } from "@runtime";

export function Menu() {
  const [activeRoute, setActiveRoute] = useState("/");
  const [_, setLocation] = useLocation();

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
      // console.log(`Navigating to ${msg.route}`);
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
    // console.log("Menu::handleRouteChange", route);
    setActiveRoute(route);
    if (route.startsWith("/history")) {
      // console.log("Menu::startsWith /history", route);
      GetLastSub("/history").then((subRoute) => {
        // console.log("Menu::GetLastSub", subRoute);
        route = route.replace("/:address", subRoute);
        // console.log("Menu::newRoute", route);
        setLocation(route);
      });
      SetLast("route", route);
      setActiveRoute("/history/:address");
    } else {
      SetLast("route", route);
      setActiveRoute(route);
    }
  };

  return (
    <div style={{ flexGrow: 1 }}>
      {routeItems
        .sort((a, b) => a.order - b.order)
        .map((item) => (
          <StyledNavLink
            key={item.route}
            label={item.label}
            icon={item.icon}
            href={item.route}
            onClick={() => handleRouteChange(item.route)}
            activeRoute={activeRoute}
          />
        ))}
    </div>
  );
}

type StyledNavLinkProps = {
  label: string;
  href: string;
  icon?: ReactNode;
  children?: ReactNode;
  onClick?: () => void;
  activeRoute: string;
};

function StyledNavLink(params: StyledNavLinkProps) {
  const [isActive] = useRoute(params.href);
  const isActiveRoute = params.activeRoute === params.href;
  return (
    <Link style={{ color: "white" }} href={params.href}>
      <NavLink
        label={params.label}
        active={isActive || isActiveRoute}
        leftSection={params.icon}
        onClick={params.onClick}
      >
        {params.children}
      </NavLink>
    </Link>
  );
}
