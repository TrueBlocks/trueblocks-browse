import { useState, useEffect, useRef } from "react";
import { Text } from "@mantine/core";
import { messages } from "@gocode/models";
import { EventsOn, EventsOff } from "@runtime";
import { notifyError, notifyInfo, notifyWarning } from "../notifications";
import classes from "./View.module.css";

export const ViewStatus = () => {
  const timeoutRef = useRef<NodeJS.Timeout | null>(null);
  const [statusMessage, setStatusMessage] = useState<string>("");
  const [color, setColor] = useState<string>("green");

  useEffect(() => {
    const handleStarted = (msg: messages.MessageMsg) => {
      notifyInfo(`Started (${msg.address})`);
    };

    const handleProgress = (msg: messages.MessageMsg) => {
      setStatusMessage(`Progress (${msg.address}): ${msg.num1}/${msg.num2}`);
      setColor("green");
    };

    const handleCompleted = (msg: messages.MessageMsg) => {
      notifyInfo(`Completed (${msg.address}): ${msg.num1}/${msg.num2}`);
      setStatusMessage("");
    };

    const handleCanceled = (msg: messages.MessageMsg) => {
      setStatusMessage(`Canceled (${msg.address})`);
      setColor("green");
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
      timeoutRef.current = setTimeout(() => {
        setStatusMessage("");
      }, 2000);
    };

    const handleInfo = (msg: messages.MessageMsg) => {
      notifyInfo(`${msg.string1} ${msg.string2}`);
    };

    const handleWarning = (msg: messages.MessageMsg) => {
      notifyWarning(`${msg.string1} ${msg.string2}`);
    };

    const handleError = (msg: messages.MessageMsg) => {
      notifyError(`${msg.string1} ${msg.string2}`);
    };

    const { Message } = messages;
    EventsOn(Message.STARTED, handleStarted);
    EventsOn(Message.PROGRESS, handleProgress);
    EventsOn(Message.COMPLETED, handleCompleted);
    EventsOn(Message.CANCELED, handleCanceled);
    EventsOn(Message.INFO, handleInfo);
    EventsOn(Message.WARNING, handleWarning);
    EventsOn(Message.ERROR, handleError);

    return () => {
      EventsOff(Message.STARTED);
      EventsOff(Message.PROGRESS);
      EventsOff(Message.COMPLETED);
      EventsOff(Message.CANCELED);
      EventsOff(Message.INFO);
      EventsOff(Message.WARNING);
      EventsOff(Message.ERROR);
    };
  }, []);

  const blankSpace = "\u00A0";
  return (
    <div className={classes.viewStatus}>
      <Text size="lg" c={color}>{statusMessage || blankSpace}</Text>
    </div>
  );
};
