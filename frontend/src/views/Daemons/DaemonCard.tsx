import { useState, useEffect } from "react";
import { Card, Text, Group, Badge, Title } from "@mantine/core";
import { GetDaemonState } from "@gocode/app/App";
import { daemons } from "@gocode/models";

export const DaemonCard = ({ daemon, toggle }: { daemon: daemons.Daemon; toggle: (name: string) => void }) => {
  const [stateStr, setStateStr] = useState<string>("");
  const { name, sleep, started, color, ticks } = daemon;

  useEffect(() => {
    GetDaemonState(name).then((s) => {
      setStateStr(s);
    });
  }, [name, daemon]);

  const handleToggle = () => {
    toggle(name);
  };

  return (
    <Card>
      <Group style={{ justifyContent: "space-between" }}>
        <Title order={4} c={color}>
          {name}
        </Title>
        <div onClick={handleToggle} style={{ cursor: "pointer" }}>
          <Badge bg={stateStr === daemons.State.RUNNING ? "green" : "red"}>{stateStr}</Badge>
        </div>
      </Group>
      <Text size="sm">Sleep Duration: {sleep}</Text>
      <Text size="sm">Started At: {new Date(started).toLocaleString()}</Text>
      <Text size="sm">Ticks: {ticks}</Text>
    </Card>
  );
};
