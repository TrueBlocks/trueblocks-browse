import { useEffect, useState } from "react";
import { StyledNavLink } from "@components";
import { types } from "@gocode/models";
import { routeItems, RouteItem } from "@layout";
import { useAppState } from "@state";

export const Menu = () => {
  const [filteredMenu, setFilteredMenu] = useState<RouteItem[]>([]);
  const { wizard, route, routeChanged } = useAppState();

  useEffect(() => {
    setFilteredMenu(
      routeItems
        .filter((item: RouteItem) =>
          wizard.state === types.WizState.FINISHED ? item.route !== "wizard" : item.route === "wizard"
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
            href={"/" + item.route}
            onClick={() => routeChanged(item.route)}
            activeRoute={"/" + route}
          />
        );
      })}
    </div>
  );
};
