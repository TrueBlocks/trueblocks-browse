import { useState, useCallback } from "react";
import { Formatter, FormatterProps, Popup, AppearancePopup } from "@components";

export const AppearanceFormatter = ({ value, value2, className }: Omit<Omit<FormatterProps, "type">, "size">) => {
  const [isPopupOpen, setPopupOpen] = useState(false);

  const onClose = useCallback(() => setPopupOpen(false), []);
  const appPopup = <AppearancePopup hash={String(value2)} onSubmit={onClose} onClose={onClose} />;
  return (
    <Popup editor={isPopupOpen ? appPopup : null}>
      <div onClick={() => setPopupOpen(true)}>
        <Formatter className={className} type="text" value={value} />
      </div>
    </Popup>
  );
};
