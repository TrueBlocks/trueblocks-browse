import { IconCopy } from "@tabler/icons-react";
import { BaseButton, ButtonProps, ButtonMouseEvent, notifyCopy } from "@components";
import { ClipboardSetText } from "../../../wailsjs/runtime/runtime";

// CopyButton copies the address of the row to the clipboard.
export const CopyButton = ({ value, onClose, ...props }: ButtonProps) => {
  const handleClick = (e: ButtonMouseEvent) => {
    ClipboardSetText(value as string).then(() => {});
    const shortened = (val: string) => (val.length > 14 ? `${val.slice(0, 8)}...${val.slice(-6)}` : val);
    notifyCopy(shortened(value as string));
  };

  return <BaseButton {...props} tip="Copy to clipboard" onClick={handleClick} onClose={onClose} icon={<IconCopy />} />;
};
