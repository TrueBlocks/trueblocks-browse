import React, { ReactNode, useState, useEffect } from "react";
import classes from "./ViewStatus.module.css";
import { messages } from "@gocode/models";
import { EventsOn, EventsOff } from "@runtime";
import { Text } from "@mantine/core";

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

    const handleWarning = (msg: messages.ErrorMsg) => {
      setStatusMessage(`Warning: ${msg.errStr} ${msg.address}`);
      setColor(classes.yellow);
    };

    const handleError = (msg: messages.ErrorMsg) => {
      setStatusMessage(`Error: ${msg.errStr} ${msg.address}`);
      setColor(classes.red);
    };

    var { Message } = messages;
    EventsOn(Message.DOCUMENT, handleDocument);
    EventsOn(Message.PROGRESS, handleProgress);
    EventsOn(Message.COMPLETED, handleCompleted);
    EventsOn(Message.WARN, handleWarning);
    EventsOn(Message.ERROR, handleError);

    return () => {
      EventsOff(Message.DOCUMENT);
      EventsOff(Message.PROGRESS);
      EventsOff(Message.COMPLETED);
      EventsOff(Message.WARN);
      EventsOff(Message.ERROR);
    };
  }, []);

  return (
    <Text size="lg">
      <div className={color}>{statusMessage}</div>
    </Text>
  );
}
