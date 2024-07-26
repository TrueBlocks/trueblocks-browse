import React, { useEffect, useMemo, useState, useRef } from "react";

export function usePopup(popup: ReturnType<typeof useRef<HTMLDivElement | null>>) {
  const id = useMemo(() => crypto.randomUUID(), []);
  const [target, setTarget] = useState<HTMLElement>()
  const floatingStyles = useMemo<React.CSSProperties>(() => {
    if (!target) {
      return { display: "none" }
    }
    const { top, left, width, height } = target.getClientRects()[0];
    const scroll = window.scrollY;

    return {
      position: "absolute",
      left: "0px",
      top: "0px",
      transform: `translate(${left}px, ${top + scroll}px)`,
      willChange: "transform",
      zIndex: 1,
      width: `${width}px`,
      minHeight: `${height}px`
    };
  }, [target]);
  const triggerProps = useMemo(() => ({
    "data-popup-trigger": id
  }), []);

  useEffect(() => {
    const listener = (event: MouseEvent) => {
      if (event.target == undefined) return;
      if (event.target == target) return;
      if ((event.target as Element).matches(`[data-popup-trigger="${id}"]`)) return;
      if (event.target == popup.current) return;
      setTarget(undefined);
    };
    window.addEventListener('click', listener)
    return () => window.removeEventListener('click', listener);
  }, [])

  return {floatingStyles, setTarget, triggerProps};
}