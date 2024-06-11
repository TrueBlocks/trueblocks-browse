import React, { ReactNode } from "react";
import { NavLink } from "@mantine/core";
import { IconHome, IconSettings, IconTag } from "@tabler/icons-react";
import { Link, useRoute } from "wouter";

// StyledNavLink is a helper component that renders navigation link
// and sets its `active` prop (makes active link more visible)
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

function GlobalMenu() {
  return (
    <>
      <StyledNavLink label="Home" icon={<IconHome />} href="/" />
      <StyledNavLink label="Names" icon={<IconTag />} href="/names" />
      <StyledNavLink
        label="Settings"
        icon={<IconSettings />}
        href="/settings"
      />
    </>
  );
}

export default GlobalMenu;
