import { IconCopy } from "@tabler/icons-react";
import { BaseButton, ButtonProps, ButtonMouseEvent, notifyCopy } from "@components";
import { useUtils } from "@hooks";
import { ClipboardSetText } from "@runtime";

// CopyButton copies the address of the row to the clipboard.
export const CopyButton = ({ value, onClose, ...props }: ButtonProps) => {
  const { ShortenAddr } = useUtils();
  const handleClick = (e: ButtonMouseEvent) => {
    ClipboardSetText(value as string).then(() => {});
    notifyCopy(ShortenAddr(value as string));
  };

  return <BaseButton {...props} tip="Copy to clipboard" onClick={handleClick} onClose={onClose} icon={<IconCopy />} />;
};
