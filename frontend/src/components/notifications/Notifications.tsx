import { showNotification } from "@mantine/notifications";
import { IconCheck, IconX } from "@tabler/icons-react";

export const notifySuccess = (message: string) => {
  showNotification({
    position: "top-center",
    title: "Success",
    size: "md",
    message,
    icon: <IconCheck size={16} />,
    color: "green",
  });
};

export const notifyError = (message: string) => {
  showNotification({
    position: "top-center",
    title: "Error",
    size: "md",
    message,
    icon: <IconX size={16} />,
    color: "red",
  });
};

// TODO: BOGUS ICON IS WRONG
export const notifyWarning = (message: string) => {
  showNotification({
    position: "top-center",
    title: "Warning",
    size: "md",
    message,
    icon: <IconX size={16} />,
  });
};

// TODO: BOGUS ICON IS WRONG
export const notifyInfo = (message: string) => {
  showNotification({
    position: "top-center",
    title: "Info",
    size: "md",
    message,
    icon: <IconX size={16} />,
  });
};

export const notifyCopy = (message: string) => {
  showNotification({
    position: "top-center",
    title: "Copied",
    size: "md",
    message: `Copied ${message} to clipboard`,
    icon: <IconCheck size={16} />,
    color: "green",
  });
};
