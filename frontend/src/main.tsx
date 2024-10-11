import { StrictMode } from "react";
import { MantineProvider } from "@mantine/core";
import { Notifications } from "@mantine/notifications";
import { createRoot } from "react-dom/client";
import "./style.css";
import { App } from "./App";
import "@mantine/core/styles.css";
import "@mantine/notifications/styles.css";

const container = document.getElementById("root");
const root = createRoot(container!);
root.render(
  <StrictMode>
    <MantineProvider defaultColorScheme="dark">
      <Notifications position="bottom-center" />
      <App />
    </MantineProvider>
  </StrictMode>
);
