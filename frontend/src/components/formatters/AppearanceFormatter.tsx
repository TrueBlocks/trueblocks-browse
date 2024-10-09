import { useState, useCallback } from "react";
import { Formatter, FormatterProps, Popup, AppearancePopup } from "@components";
import { ClipboardSetText } from "@runtime";

export const AppearanceFormatter = ({ value, value2, className }: Omit<Omit<FormatterProps, "type">, "size">) => {
  const [isPopupOpen, setPopupOpen] = useState(false);

  const copyHash = useCallback(() => {
    ClipboardSetText(String(value2)).then(() => {
      setPopupOpen(false);
    });
  }, [value2]);

  const appPopup = (
    <AppearancePopup
      hash={String(value2)}
      onSubmit={() => setPopupOpen(false)}
      onClose={() => setPopupOpen(false)}
      onCopy={copyHash}
    />
  );

  const editor = isPopupOpen ? appPopup : null;
  return (
    <Popup editor={editor}>
      <div onClick={() => setPopupOpen(true)}>
        <Formatter className={className} type="text" value={value} />
      </div>
    </Popup>
  );
};
