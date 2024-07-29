import React, { ReactNode, useEffect } from "react";
import { Badge } from "@mantine/core";
import { types, names } from "@gocode/models";

export const NameTags = ({ name }: { name: types.NameEx }) => {
  const [tags, setTags] = React.useState<ReactNode>([]);

  useEffect(() => {
    var types: ReactNode[] = [];
    if (name.type & names.Parts.REGULAR) {
      types.push(<Badge color="blue">R</Badge>);
    }
    if (name.type & names.Parts.CUSTOM) {
      types.push(<Badge color="yellow">C</Badge>);
    }
    if (name.type & names.Parts.PREFUND) {
      types.push(<Badge color="green">P</Badge>);
    }
    if (name.type & names.Parts.BADDRESS) {
      types.push(<Badge color="pink">B</Badge>);
    }
    setTags(<div>{types.map((tag) => tag)}</div>);
  }, [name]);

  return <div>{tags}</div>;
};
