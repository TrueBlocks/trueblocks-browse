import React, { ReactNode } from "react";
import { NavLink } from "@mantine/core";
import { IconHome, IconSettings, IconTag, IconList } from "@tabler/icons-react";
import { Link, useRoute } from "wouter";

function Menu() {
  return (
    <div style={{ flexGrow: 1 }}>
      <StyledNavLink label="Home" icon={<IconHome />} href="/" />
      <StyledNavLink label="Settings" icon={<IconSettings />} href="/settings" />
    </div>
  );
}

export default Menu;

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
