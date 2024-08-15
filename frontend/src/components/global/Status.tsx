import React, { useState, useEffect } from "react";
import { Text } from "@mantine/core";
import { messages } from "@gocode/models";
import { GetStatus } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";

export const Status = () => {
  const [updated, setUpdated] = useState<string>("");
  const [node, setNode] = useState<string | undefined>("");
  const [file, setFile] = useState<string | undefined>("");
  const [chain, setChain] = useState<string | undefined>("");

  useEffect(() => {
    const fetch = async () => {
      GetStatus(0, 15).then((status) => {
        setUpdated(status.latestUpdate);
        setNode(status.clientVersion);
        // const now = new Date();
        // const currentTimeString = now.toLocaleTimeString();
        setFile("not loaded");
        setChain(status.chain ? status.chain : "mainnet");
      });
    };
    fetch();

    const handleRefresh = () => {
      // const now = new Date();
      // const currentTimeString = now.toLocaleTimeString();
      // console.log(currentTimeString);
      // console.log("heard that in Status");
      fetch();
    };

    var { Message } = messages;
    EventsOn(Message.DAEMON, handleRefresh);
    return () => {
      EventsOff(Message.DAEMON);
    };
  }, []);

  return <Text size={"sm"}>{`${node} / ${chain} / ${file} / ${updated}`}</Text>;
};
