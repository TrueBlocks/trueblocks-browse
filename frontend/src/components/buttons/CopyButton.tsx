import { showNotification } from "@mantine/notifications";
import { IconCopy, IconCheck } from "@tabler/icons-react";
import { BaseButton, ButtonProps, ButtonMouseEvent } from "@components";

// CopyButton copies the address of the row to the clipboard.
export const CopyButton = ({ value, onClick, ...props }: ButtonProps) => {
  const handleClick = (e: ButtonMouseEvent) => {
    if (onClick) {
      onClick(e);
    }

    const shortened = (val: string) => (val.length > 10 ? `${val.slice(0, 6)}...${val.slice(-4)}` : val);
    showNotification({
      title: "Copied",
      message: `Copied ${shortened(value as string)} to clipboard`,
      icon: <IconCheck size={16} />,
      color: "green",
      autoClose: 2000,
    });

    console.log("Copied to clipboard");
  };

  return <BaseButton {...props} tip="Copy to clipboard" onClick={handleClick} leftSection={<IconCopy />} />;
};
