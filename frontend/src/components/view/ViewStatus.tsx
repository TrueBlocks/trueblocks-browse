import { useState, useEffect, useRef } from "react";
import { Text } from "@mantine/core";
import { messages } from "@gocode/models";
import { EventsOn, EventsOff } from "@runtime";
import { notifyError } from "../notifications";
import classes from "./View.module.css";

export const ViewStatus = () => {
  const timeoutRef = useRef<NodeJS.Timeout | null>(null);
  const [statusMessage, setStatusMessage] = useState<string>("");
  const [color, setColor] = useState<string>("green");

  useEffect(() => {
    const handleProgress = (msg: messages.MessageMsg) => {
      setStatusMessage(`Progress (${msg.address}): ${msg.num1}/${msg.num2}`);
      setColor("green");
    };

    const handleCompleted = (msg: messages.MessageMsg) => {
      setStatusMessage(`Completed (${msg.address}): ${msg.num1}/${msg.num2}`);
      setColor("green");
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
      timeoutRef.current = setTimeout(() => {
        setStatusMessage("");
      }, 2000);
    };

    const handleCancelled = (msg: messages.MessageMsg) => {
      setStatusMessage(`Canceled (${msg.address})`);
      setColor("green");
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
      timeoutRef.current = setTimeout(() => {
        setStatusMessage("");
      }, 2000);
    };

    const handleWarning = (msg: messages.MessageMsg) => {
      setStatusMessage(`Warning: ${msg.string1} ${msg.address}`);
      setColor("yellow");
    };

    const handleError = (msg: messages.MessageMsg) => {
      if (!msg.string1.includes("Invalid address")) {
        notifyError(msg.string1 + "\n" + msg.string2 || "An error occurred!");
      }
      setStatusMessage(`Error: ${msg.string1} ${msg.address}`);
      setColor("red");
    };

    const handleInfo = (msg: messages.MessageMsg) => {
      setStatusMessage(`Info ${msg.string2} ${msg.string1}`);
      setColor("blue");
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
      timeoutRef.current = setTimeout(() => {
        setStatusMessage("");
      }, 2000);
    };

    const { Message } = messages;
    EventsOn(Message.CANCELLED, handleCancelled);
    EventsOn(Message.COMPLETED, handleCompleted);
    EventsOn(Message.ERROR, handleError);
    EventsOn(Message.INFO, handleInfo);
    EventsOn(Message.PROGRESS, handleProgress);
    EventsOn(Message.WARNING, handleWarning);

    return () => {
      EventsOff(Message.CANCELLED);
      EventsOff(Message.COMPLETED);
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
