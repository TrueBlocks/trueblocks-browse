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
      setStatusMessage(`${msg.msg} ${msg.filename}`);
      setColor("green");
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
      timeoutRef.current = setTimeout(() => {
        setStatusMessage("");
      }, 2000);
    };

    const handleProgress = (msg: messages.ProgressMsg) => {
      setStatusMessage(`Progress (${msg.address}): ${msg.have}/${msg.want}`);
      setColor("green");
    };

    const handleCompleted = (msg: messages.ProgressMsg) => {
      setStatusMessage(`Completed (${msg.address}): ${msg.have}/${msg.want}`);
      setColor("green");
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
      timeoutRef.current = setTimeout(() => {
        setStatusMessage("");
      }, 2000);
    };

    const handleCancel = (msg: messages.ProgressMsg) => {
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
      setStatusMessage(`Warning: ${msg.errStr} ${msg.address}`);
      setColor("yellow");
    };

    const handleError = (msg: messages.ErrorMsg) => {
      setStatusMessage(`Error: ${msg.errStr} ${msg.address}`);
      setColor("red");
    };

    const handleInfo = (msg: messages.InfoMsg) => {
      setStatusMessage(`Info [${new Date().toLocaleString()}]: ${msg.message}`);
      setColor("blue");
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
      timeoutRef.current = setTimeout(() => {
        setStatusMessage("");
      }, 2000);
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

  const blankSpace = "\u00A0";
  return (
    <div className={classes.viewStatus}>
      <Text size="lg" c={color}>
        {statusMessage || blankSpace}
      </Text>
    </div>
  );
};
