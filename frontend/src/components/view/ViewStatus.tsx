import React, { ReactNode, useState, useEffect } from "react";
import classes from "./ViewStatus.module.css";
import { messages } from "@gocode/models";
import { EventsOn, EventsOff } from "@runtime";
import { Text } from "@mantine/core";

export function ViewStatus() {
  const [statusMessage, setStatusMessage] = useState<string>("");
  const [color, setColor] = useState<string>(classes.green);

  useEffect(() => {
    const handleDone = () => {
      setStatusMessage("done");
      setColor(classes.green);
    };

    const handleProgress = (p: messages.ProgressMsg) => {
      setStatusMessage(`Progress (${p.address}): ${p.have}/${p.want}`);
      setColor(classes.green);
    };

    const handleWarning = (warnStr: string) => {
      setStatusMessage(`Warning: ${warnStr}`);
      setColor(classes.yellow);
    };

    const handleError = (errorStr: string) => {
      setStatusMessage(`Error: ${errorStr}`);
      setColor(classes.red);
    };

    EventsOn("COMPLETED", handleDone);
    EventsOn("PROGRESS", handleProgress);
    EventsOn("WARNING", handleWarning);
    EventsOn("ERROR", handleError);

    return () => {
      EventsOff("COMPLETED");
      EventsOff("PROGRESS");
      EventsOff("WARNING");
      EventsOff("ERROR");
    };
  }, []);

  return (
    <Text size="lg">
      <div className={color}>{statusMessage}</div>
    </Text>
  );
}
