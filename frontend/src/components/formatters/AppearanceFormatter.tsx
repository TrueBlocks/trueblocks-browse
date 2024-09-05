import { useState, useCallback } from "react";
import { Formatter, FormatterProps, Popup, AppearancePopup } from "@components";
import { ClipboardSetText } from "@runtime";

export const AppearanceFormatter = ({ value, value2, className, size = "md" }: Omit<FormatterProps, "type">) => {
  const [isPopupOpen, setPopupOpen] = useState(false);

  const copyHash = useCallback(() => {
    ClipboardSetText(value2).then(() => {
      setPopupOpen(false);
    });
  }, [value2]);

  const appPopup = (
    <AppearancePopup
      hash={value2}
      onSubmit={() => setPopupOpen(false)}
      onClose={() => setPopupOpen(false)}
      onCopy={() => copyHash}
    />
  );

  const editor = isPopupOpen ? appPopup : null;
  return (
    <Popup editor={editor}>
      <div onClick={() => setPopupOpen(true)}>
        <Formatter className={className} size={size} type="text" value={value} />
      </div>
    </Popup>
  );
};
