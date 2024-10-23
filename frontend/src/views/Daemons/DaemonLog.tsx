import { Card, Text } from "@mantine/core";
import { messages } from "@gocode/models";

export const DaemonLog = ({ logMessages }: { logMessages: messages.DaemonMsg[] }) => {
  return (
    <Card style={{ width: "100%", maxHeight: "16rem", overflowY: "auto" }}>
      {logMessages.map((log, index) => (
        <Text c={log.msg2} key={index}>
          {log.msg1}
        </Text>
      ))}
    </Card>
  );
};
