import React, { ReactNode, useEffect, useState } from "react";
import { NavLink } from "@mantine/core";
import { IconHome, IconTag, IconSettings } from "@tabler/icons-react";
import { Link, useRoute } from "wouter";
import { GetLast, SetLast } from "@gocode/app/App";

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

  var menuItems = [
    { route: "/", label: "Home", icon: <IconHome /> },
    { route: "/names", label: "Names", icon: <IconTag /> },
    { route: "/settings", label: "Settings", icon: <IconSettings /> },
  ];

  return (
    <div style={{ flexGrow: 1 }}>
      {menuItems.map((item) => (
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

type StyledNavLinkParams = {
  label: string;
  href: string;
  icon?: ReactNode;
  children?: ReactNode;
  onClick?: () => void;
  activeRoute: string;
};

function StyledNavLink(params: StyledNavLinkParams) {
  const [isActive] = useRoute(params.href);
  const isActiveRoute = params.activeRoute === params.href;
  return (
    <Link href={params.href}>
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
