import React, { ReactNode, useEffect } from "react";
import { Badge } from "@mantine/core";
import { types } from "@gocode/models";

export const NameTags = ({ name }: { name: types.Name }) => {
  const [tags, setTags] = React.useState<ReactNode>([]);

  useEffect(() => {
    // TODO: Can't figure out how to get this in the frontend
    // REGULAR = 2,
    // CUSTOM = 4,
    // PREFUND = 8,
    // BADDRESS = 16,
    var types: ReactNode[] = [];
    if (name.parts && name.parts & 2 /* REGULAR */) {
      types.push(<Badge color="blue">R</Badge>);
    }
    if (name.parts && name.parts & 4 /* CUSTOM */) {
      types.push(<Badge color="yellow">C</Badge>);
    }
    if (name.parts && name.parts & 8 /* PREFUND */) {
      types.push(<Badge color="green">P</Badge>);
    }
    if (name.parts && name.parts & 16 /* BADDRESS */) {
      types.push(<Badge color="pink">B</Badge>);
    }
    setTags(<div>{types.map((tag) => tag)}</div>);
  }, [name]);

  return <div>{tags}</div>;
};
