import React, { ReactNode, useEffect } from "react";
import { Badge } from "@mantine/core";
import { app } from "@gocode/models";

export const NameTags = ({ name }: { name: app.NameEx }) => {
  const [tags, setTags] = React.useState<ReactNode>([]);

  useEffect(() => {
    var types: ReactNode[] = [];
    if (name.type & 2) {
      types.push(<Badge color="blue">Reg</Badge>);
    }
    if (name.type & 4) {
      types.push(<Badge color="yellow">Cus</Badge>);
    }
    if (name.type & 8) {
      types.push(<Badge color="green">Pre</Badge>);
    }
    if (name.type & 16) {
      types.push(<Badge color="pink">Bad</Badge>);
    }
    setTags(<div>{types.map((tag) => tag)}</div>);
  }, [name]);

  return <div>{tags}</div>;
};
