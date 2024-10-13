import { useState, useEffect, useCallback } from "react";
import { Title } from "@mantine/core";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import { useLocation } from "wouter";
import { CloseButton } from "@components";
import { messages } from "@gocode/models";
import { useViewName } from "@hooks";
import { EventsEmit } from "@runtime";
import classes from "./Help.module.css";

// Glob import for markdown files as raw content
const helpFiles = import.meta.glob("/src/assets/help/*.md", { query: "?raw", import: "default" }) as Record<
  string,
  () => Promise<string>
>;

export function Help(): JSX.Element {
  const [location] = useLocation();
  const [markdown, setMarkdown] = useState<string>("Loading...");
  const [error, setError] = useState<boolean>(false);
  const viewName = useViewName();

  const onClose = useCallback(() => {
    EventsEmit(messages.Message.TOGGLE, { Component: "help" });
  }, []);

  useEffect(() => {
    const baseRoute = location.split("/")[1];
    const helpFileName: string = `${baseRoute === "" ? "project" : baseRoute}.md`;
    const filePath = Object.keys(helpFiles).find((key) => key.endsWith(`/help/${helpFileName}`));

    const loadMarkdown = async (): Promise<void> => {
      if (filePath) {
        try {
          const content = await helpFiles[filePath](); // Await the promise to get the raw content
          setMarkdown(content);
        } catch (error) {
          setError(true);
          setMarkdown("Sorry, the help file could not be loaded: " + error);
        }
      } else {
        setError(true);
        setMarkdown("Sorry, the help file could not be found.");
      }
    };

    loadMarkdown();
  }, [location]);

  return (
    <div className={classes.helpPanel}>
      <CloseButton onClose={onClose} />
      <Title order={4} className={classes.header}>
        {viewName}
      </Title>
      <ReactMarkdown remarkPlugins={[remarkGfm]}>{error ? "error" : markdown}</ReactMarkdown>
    </div>
  );
}
