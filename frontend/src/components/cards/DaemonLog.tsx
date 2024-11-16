import { Card, Text } from "@mantine/core";
import { messages } from "@gocode/models";

export const DaemonLog = ({ logMessages }: { logMessages: messages.MessageMsg[] }) => {
  return (
    <Card style={{ width: "100%", maxHeight: "16rem", overflowY: "auto" }}>
      {logMessages.map((log, index) => (
        <Text c={log.string2} key={index}>
          {log.string1}
        </Text>
      ))}
    </Card>
  );
};
