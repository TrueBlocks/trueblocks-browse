import React, { ReactNode } from "react";
import { NavLink } from "@mantine/core";
import { IconHome, IconSettings } from "@tabler/icons-react";
import { Link, useRoute } from "wouter";

function GlobalMenu() {
  return (
    <>
      <StyledNavLink label="Home" icon={<IconHome />} href="/" />
      <StyledNavLink
        label="Settings"
        icon={<IconSettings />}
        href="/settings"
      />
    </>
  );
}

export default GlobalMenu;

type StyledNavLinkParams = {
  label: string;
  href: string;
  icon?: ReactNode;
  children?: ReactNode;
};
function StyledNavLink(params: StyledNavLinkParams) {
  const [isActive] = useRoute(params.href);

  return (
    <Link href={params.href}>
      <NavLink label={params.label} active={isActive} leftSection={params.icon}>
        {params.children}
      </NavLink>
    </Link>
  );
}
