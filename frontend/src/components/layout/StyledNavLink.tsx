import { ReactNode } from "react";
import { NavLink } from "@mantine/core";
import { Link, useRoute } from "wouter";

type StyledNavLinkProps = {
  label: string;
  href: string;
  icon?: ReactNode;
  children?: ReactNode;
  onClick?: () => void;
  activeRoute: string;
};

export function StyledNavLink(params: StyledNavLinkProps) {
  const [isActive] = useRoute(params.href);
  const isActiveRoute = params.activeRoute === params.href;

  return (
    <NavLink
      component={Link}
      to={params.href}
      label={params.label}
      active={isActive || isActiveRoute}
      leftSection={params.icon}
      onClick={params.onClick}
      style={{ color: "white" }}
    >
      {params.children}
    </NavLink>
  );
}
