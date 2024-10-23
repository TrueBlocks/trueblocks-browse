import { useState, useEffect, useRef } from "react";
import { Text } from "@mantine/core";
import { messages } from "@gocode/models";
import { EventsOn, EventsOff } from "@runtime";
import classes from "./View.module.css";

export const ViewStatus = () => {
  const timeoutRef = useRef<NodeJS.Timeout | null>(null);
  const [statusMessage, setStatusMessage] = useState<string>("");
  const [color, setColor] = useState<string>("green");

  useEffect(() => {
    const handleDocument = (msg: messages.DocumentMsg) => {
      setStatusMessage(`${msg.msg2} ${msg.msg1}`);
      setColor("green");
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
      timeoutRef.current = setTimeout(() => {
        setStatusMessage("");
      }, 2000);
    };

    const handleProgress = (msg: messages.ProgressMsg) => {
      setStatusMessage(`Progress (${msg.address}): ${msg.int1}/${msg.int2}`);
      setColor("green");
    };

    const handleCompleted = (msg: messages.ProgressMsg) => {
      setStatusMessage(`Completed (${msg.address}): ${msg.int1}/${msg.int2}`);
      setColor("green");
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
      timeoutRef.current = setTimeout(() => {
        setStatusMessage("");
      }, 2000);
    };

    const handleCancel = (msg: messages.MessageMsg) => {
      setStatusMessage(`Canceled (${msg.address})`);
      setColor("green");
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
      timeoutRef.current = setTimeout(() => {
        setStatusMessage("");
      }, 2000);
    };

    const handleWarning = (msg: messages.ErrorMsg) => {
      setStatusMessage(`Warning: ${msg.msg1} ${msg.address}`);
      setColor("yellow");
    };

    const handleError = (msg: messages.ErrorMsg) => {
      setStatusMessage(`Error: ${msg.msg1} ${msg.address}`);
      setColor("red");
    };

    const handleInfo = (msg: messages.InfoMsg) => {
      setStatusMessage(`Info [${new Date().toLocaleString()}]: ${msg.msg1}`);
      setColor("blue");
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
      timeoutRef.current = setTimeout(() => {
        setStatusMessage("");
      }, 2000);
    };

    const { Message } = messages;
    EventsOn(Message.CANCELLED, handleCancel);
    EventsOn(Message.COMPLETED, handleCompleted);
    EventsOn(Message.DOCUMENT, handleDocument);
    EventsOn(Message.ERROR, handleError);
    EventsOn(Message.INFO, handleInfo);
    EventsOn(Message.PROGRESS, handleProgress);
    EventsOn(Message.WARNING, handleWarning);

    return () => {
      EventsOff(Message.CANCELLED);
      EventsOff(Message.COMPLETED);
      EventsOff(Message.DOCUMENT);
      EventsOff(Message.ERROR);
      EventsOff(Message.INFO);
      EventsOff(Message.PROGRESS);
      EventsOff(Message.WARNING);
    };
  }, []);

  const blankSpace = "\u00A0";
  return (
    <div className={classes.viewStatus}>
      <Text size="lg" c={color}>
        {statusMessage || blankSpace}
      </Text>
    </div>
  );
};
