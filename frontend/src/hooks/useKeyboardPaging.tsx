import { useEffect, useState, DependencyList } from "react";
import { useHotkeys } from "react-hotkeys-hook";

export function useKeyboardPaging<T>(items: T[], nItems: number, deps: DependencyList = [], perPage: number = 20) {
  const [curItem, setCurItem] = useState<number>(0);

  useHotkeys("left", (event) => {
    setCurItem((cur) => Math.max(cur - 1, 0));
    event.preventDefault();
  });
  useHotkeys("up", (event) => {
    setCurItem((cur) => Math.max(cur - perPage, 0));
    event.preventDefault();
  });
  useHotkeys("home", (event) => {
    setCurItem(0);
    event.preventDefault();
  });

  useHotkeys("right", (event) => {
    var max = Math.max(nItems - perPage, 0);
    setCurItem((cur) => Math.min(max, cur + 1));
    event.preventDefault();
  });
  useHotkeys("down", (event) => {
    var max = Math.max(nItems - perPage, 0);
    setCurItem((cur) => Math.min(max, cur + perPage));
    event.preventDefault();
  });
  useHotkeys("end", (event) => {
    var max = Math.max(nItems - perPage, 0);
    setCurItem(max);
    event.preventDefault();
  });

  useEffect(() => {
    setCurItem(0);
  }, deps);

  return { curItem, perPage };
}
