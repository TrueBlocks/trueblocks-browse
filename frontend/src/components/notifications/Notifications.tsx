import { showNotification } from "@mantine/notifications";
import { IconCheck, IconX } from "@tabler/icons-react";

export const notifySuccess = (message: string) => {
  showNotification({
    title: "Success",
    message,
    icon: <IconCheck size={16} />,
    color: "green",
  });
};

export const notifyError = (message: string) => {
  showNotification({
    title: "Error",
    message,
    icon: <IconX size={16} />,
    color: "red",
  });
};

export const notifyCopy = (message: string) => {
  showNotification({
    title: "Copied",
    message: `Copied ${message} to clipboard`,
    icon: <IconCheck size={16} />,
    color: "green",
  });
};
