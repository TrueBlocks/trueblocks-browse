import React, { useState, useEffect } from "react";
import { Text } from "@mantine/core";
import { messages } from "@gocode/models";
import { StatusPage } from "@gocode/app/App";
import { EventsOn, EventsOff } from "@runtime";

export const AppStatus = () => {
  const [updated, setUpdated] = useState<string>("");
  const [node, setNode] = useState<string | undefined>("");
  const [file, setFile] = useState<string | undefined>("");
  const [chain, setChain] = useState<string | undefined>("");

  useEffect(() => {
    const fetch = async () => {
      StatusPage(0, 15).then((status) => {
        setUpdated(status.latestUpdate);
        setNode(status.clientVersion);
        setFile("not loaded");
        setChain(status.chain ? status.chain : "mainnet");
      });
    };
    fetch();

    const handleRefresh = () => {
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
