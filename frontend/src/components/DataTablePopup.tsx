import { ActionIcon, Button, Group, TextInput } from "@mantine/core";
import { IconEdit } from "@tabler/icons-react";
import React, { useEffect, useMemo, useState, useRef, useContext } from "react";
import { useHotkeys } from "react-hotkeys-hook";

export const PopupContext = React.createContext<{opened?: boolean, setTarget?: React.Dispatch<React.SetStateAction<HTMLElement | undefined>>}>({});

export function usePopup<T extends HTMLElement>(popup: ReturnType<typeof useRef<T | null>>) {
  const id = useMemo(() => crypto.randomUUID(), []);
  const [target, setTarget] = useState<HTMLElement>()
  const opened = useMemo(() => !!target, [target])
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
  useHotkeys("Escape", (event) => {
    event.preventDefault();
    event.stopImmediatePropagation();
    setTarget(undefined);
  })

  useEffect(() => {
    const listener = (event: MouseEvent) => {
      if (event.target == undefined) return;
      if (event.target == target) return;

      // const closestPopup = (event.target as Element).closest('[data-popup]');
      // console.log(">>>", !!closestPopup, event.target)
      // if (closestPopup) {
      //   return
      // }
      console.log(">>>", event.target);
      if ((event.target as Element).matches('[data-popup] *')) return;

      if ((event.target as Element).matches(`[data-popup-trigger="${id}"]`)) return;

      setTarget(undefined);
    };
    window.addEventListener('click', listener)
    return () => window.removeEventListener('click', listener);
  }, [])

  return {floatingStyles, setTarget, triggerProps, opened};
}

export function DataTableStringEditor({ value }: { value?: () => any }) {
  const [inputValue, setInputValue] = useState(String(value?.() || ""));
  const [edit, setEdit] = useState(false);
  const { opened, setTarget } = useContext(PopupContext);

  if (!edit) {
    return (
      <Group>
        <div>{inputValue}</div>
        <ActionIcon onClick={() => setEdit(true)}>
          <IconEdit />
        </ActionIcon>
      </Group>
    );
  }

  return (
    <form onSubmit={(e) => e.preventDefault()}>
      <Group>
        <TextInput
          value={value?.()}
          onChange={(event) => setInputValue(event.currentTarget.value)}
          // error={inputValue.length > 0}

        />
        <Button type="submit">Save</Button>
        <Button type="button" variant="outline" onClick={() => { setEdit(false); setTarget?.(undefined); }}>Cancel</Button>
      </Group>
    </form>
  );
}