import { useState, useEffect, useCallback } from "react";
import { Title } from "@mantine/core";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import { useLocation } from "wouter";
import { CloseButton } from "@components";
import { messages } from "@gocode/models";
import { useViewName } from "@hooks";
import { EventsEmit } from "@runtime";
import { useAppState } from "../../state";
import classes from "./Help.module.css";

// Glob import for markdown files as raw content
const helpFiles = import.meta.glob("/src/assets/help/*.md", { query: "?raw", import: "default" }) as Record<
  string,
  () => Promise<string>
>;

export function Help(): JSX.Element {
  const { wizardState } = useAppState();
  const [location] = useLocation();
  const [markdown, setMarkdown] = useState<string>("Loading...");
  const [error, setError] = useState<boolean>(false);
  const viewName = useViewName();

  const onClose = useCallback(() => {
    EventsEmit(messages.Message.TOGGLELAYOUT, { string1: "help" });
  }, []);

  useEffect(() => {
    const baseRoute = location.split("/")[1];
    const helpFileName: string =
      baseRoute === "wizard" ? `wizard-${String(wizardState)}.md` : `${baseRoute === "" ? "project" : baseRoute}.md`;
    const filePath = Object.keys(helpFiles).find((key) => key.endsWith(`/help/${helpFileName}`));

    const loadMarkdown = async (): Promise<void> => {
      if (!filePath) {
        setError(true);
        setMarkdown("Sorry, the help file could not be found.");
        return;
      }
      try {
        const content = await helpFiles[filePath]();
        setMarkdown(content);
      } catch (error) {
        setError(true);
        setMarkdown("Sorry, the help file could not be loaded: " + error);
      }
    };

    loadMarkdown();
  }, [location, wizardState]);

  return (
    <div className={classes.helpPanel}>
      <CloseButton onClose={onClose} />
      <Title order={4} className={classes.header}>
        {viewName}
      </Title>
      <ReactMarkdown remarkPlugins={[remarkGfm]}>{error ? markdown : markdown}</ReactMarkdown>
    </div>
  );
}
