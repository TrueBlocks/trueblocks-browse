import { useState, useCallback } from "react";
import { Formatter, FormatterProps, Popup, AppearancePopup } from "@components";
import { ClipboardSetText } from "@runtime";

export const AppearanceFormatter = ({ value, value2, className }: Omit<Omit<FormatterProps, "type">, "size">) => {
  const [isPopupOpen, setPopupOpen] = useState(false);

  const onCopy = useCallback(() => {
    ClipboardSetText(String(value2)).then(() => {
      setPopupOpen(false);
    });
  }, [value2]);

  const onClose = useCallback(() => setPopupOpen(false), []);

  const appPopup = <AppearancePopup hash={String(value2)} onSubmit={onClose} onClose={onClose} onCopy={onCopy} />;
  return (
    <Popup editor={isPopupOpen ? appPopup : null}>
      <div onClick={() => setPopupOpen(true)}>
        <Formatter className={className} type="text" value={value} />
      </div>
    </Popup>
  );
};
