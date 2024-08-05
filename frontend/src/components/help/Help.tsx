import React, { useState, useEffect } from "react";
import { Text } from "@mantine/core";
import { useLocation } from "wouter";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";

// Glob import for markdown files as raw content
const helpFiles = import.meta.glob("/src/assets/help/*.md", { as: "raw" }) as Record<string, () => Promise<string>>;

export function Help(): JSX.Element {
  const [location] = useLocation();
  const [markdown, setMarkdown] = useState<string>("Loading...");
  const [error, setError] = useState<boolean>(false);

  useEffect(() => {
    const baseRoute = location.split("/")[1];
    const helpFileName: string = `${baseRoute === "" ? "home" : baseRoute}.md`;
    const filePath = Object.keys(helpFiles).find((key) => key.endsWith(`/help/${helpFileName}`));

    const loadMarkdown = async (): Promise<void> => {
      if (filePath) {
        try {
          const content = await helpFiles[filePath](); // Await the promise to get the raw content
          setMarkdown(content);
        } catch (error) {
          setError(true);
          setMarkdown("Sorry, the help file could not be loaded.");
        }
      } else {
        setError(true);
        setMarkdown("Sorry, the help file could not be found.");
      }
    };

    loadMarkdown();
  }, [location]);

  return (
    <div>
      <ReactMarkdown remarkPlugins={[remarkGfm]}>{markdown}</ReactMarkdown>
    </div>
  );
}
