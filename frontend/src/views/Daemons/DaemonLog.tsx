import React from "react";
import { Card, Text, Title } from "@mantine/core";
import { messages } from "@gocode/models";

export const DaemonLog = ({ logMessages }: { logMessages: messages.DaemonMsg[] }) => {
  return (
    <Card style={{ maxHeight: "16rem", overflowY: "auto" }}>
      {logMessages.map((log, index) => (
        <Text c={log.color} key={index}>
          {log.message}
        </Text>
      ))}
    </Card>
  );
};
