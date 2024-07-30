import React, { useState, useEffect } from "react";
import { Card, Text, Group, Badge, Title } from "@mantine/core";
import { StateToString } from "@gocode/app/App";
import { servers } from "@gocode/models";

export const ServerCard = ({ server, toggle }: { server: servers.Server; toggle: (name: servers.Type) => void }) => {
  const [stateStr, setStateStr] = useState<string>("");
  const { name, sleep, started, color, runs, state } = server;

  useEffect(() => {
    StateToString(name).then((s) => {
      setStateStr(s);
    });
  }, [state]);

  const handleToggle = () => {
    toggle(name);
  };

  return (
    <Card shadow="xl">
      <Group style={{ justifyContent: "space-between", marginBottom: 5 }}>
        <Title order={4} c={color}>
          {name}
        </Title>
        <div onClick={handleToggle} style={{ cursor: "pointer" }}>
          <Badge bg={state === servers.State.RUNNING ? "green" : "red"}>{stateStr}</Badge>
        </div>
      </Group>
      <Text size="sm" style={{ lineHeight: 1.5 }}>
        Sleep Duration: {sleep}
      </Text>
      <Text size="sm" style={{ lineHeight: 1.5 }}>
        Started At: {new Date(started).toLocaleString()}
      </Text>
      <Text size="sm" style={{ lineHeight: 1.5 }}>
        Runs: {runs}
      </Text>
    </Card>
  );
};
