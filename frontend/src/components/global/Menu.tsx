import React, { ReactNode, useEffect, useState } from "react";
import { NavLink } from "@mantine/core";
import { Link, useRoute } from "wouter";
import { GetLast, SetLast } from "@gocode/app/App";
// TODO: This alias is wrong, can it not be @Routes See also @/App.module.css
import { routeItems } from "@/Routes";

export function Menu() {
  const [activeRoute, setActiveRoute] = useState("/");

  useEffect(() => {
    const lastRoute = (GetLast("route") || "/").then((route) => {
      setActiveRoute(route);
    });
  }, []);

  const handleRouteChange = (route: string) => {
    SetLast("route", route);
    setActiveRoute(route);
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
