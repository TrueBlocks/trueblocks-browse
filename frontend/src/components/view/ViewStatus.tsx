import { useState, useEffect } from "react";
import { Text } from "@mantine/core";
import { messages } from "@gocode/models";
import { EventsOn, EventsOff } from "@runtime";
import classes from "./ViewStatus.module.css";

export function ViewStatus() {
  const [statusMessage, setStatusMessage] = useState<string>("");
  const [color, setColor] = useState<string>(classes.green);

  useEffect(() => {
    const handleDocument = (msg: messages.DocumentMsg) => {
      setStatusMessage(`${msg.msg} ${msg.filename}`);
      setColor(classes.green);
    };

    const handleProgress = (msg: messages.ProgressMsg) => {
      setStatusMessage(`Progress (${msg.address}): ${msg.have}/${msg.want}`);
      setColor(classes.green);
    };

    const handleCompleted = (msg: messages.ProgressMsg) => {
      setStatusMessage(`Completed (${msg.address}): ${msg.have}/${msg.want}`);
      setColor(classes.green);
      setTimeout(() => {
        setStatusMessage("");
      }, 1000); // 1000ms = 1 second
    };

    const handleCancel = (msg: messages.ProgressMsg) => {
      setStatusMessage(`Canceled (${msg.address})`);
      setColor(classes.green);
      setTimeout(() => {
        setStatusMessage("");
      }, 1000); // 1000ms = 1 second
    };

    const handleWarning = (msg: messages.ErrorMsg) => {
      setStatusMessage(`Warning: ${msg.errStr} ${msg.address}`);
      setColor(classes.yellow);
    };

    const handleError = (msg: messages.ErrorMsg) => {
      setStatusMessage(`Error: ${msg.errStr} ${msg.address}`);
      setColor(classes.red);
    };

    const handleInfo = (msg: messages.InfoMsg) => {
      setStatusMessage(`Info: ${msg.message}`);
      setColor(classes.blue);
    };

    const { Message } = messages;
    EventsOn(Message.DOCUMENT, handleDocument);
    EventsOn(Message.PROGRESS, handleProgress);
    EventsOn(Message.COMPLETED, handleCompleted);
    EventsOn(Message.CANCELLED, handleCancel);
    EventsOn(Message.WARNING, handleWarning);
    EventsOn(Message.ERROR, handleError);
    EventsOn(Message.INFO, handleInfo);

    return () => {
      EventsOff(Message.DOCUMENT);
      EventsOff(Message.PROGRESS);
      EventsOff(Message.COMPLETED);
      EventsOff(Message.CANCELLED);
      EventsOff(Message.WARNING);
      EventsOff(Message.ERROR);
      EventsOff(Message.INFO);
    };
  }, []);

  return (
    <Text size="lg">
      <div className={color}>{statusMessage}</div>
    </Text>
  );
}
