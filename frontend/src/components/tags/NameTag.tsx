import { ReactNode, useEffect, useState } from "react";
import { Badge } from "@mantine/core";
import { types } from "@gocode/models";

export const NameTags = ({ name }: { name: types.Name }) => {
  const [tags, setTags] = useState<ReactNode>([]);

  useEffect(() => {
    // TODO: Can't figure out how to get this in the frontend
    // REGULAR = 2,
    // CUSTOM = 4,
    // PREFUND = 8,
    // BADDRESS = 16,
    const types: ReactNode[] = [];
    if (name.parts && name.parts & 2 /* REGULAR */) {
      types.push(
        <Badge key="regular" size="xs" color="blue">
          R
        </Badge>
      );
    }
    if (name.parts && name.parts & 4 /* CUSTOM */) {
      types.push(
        <Badge key="custom" size="xs" color="yellow">
          C
        </Badge>
      );
    }
    if (name.parts && name.parts & 8 /* PREFUND */) {
      types.push(
        <Badge key="prefund" size="xs" color="green">
          P
        </Badge>
      );
    }
    if (name.parts && name.parts & 16 /* BADDRESS */) {
      types.push(
        <Badge key="baddress" size="xs" color="pink">
          B
        </Badge>
      );
    }
    setTags(<div>{types}</div>);
  }, [name]);

  return <div>{tags}</div>;
};
