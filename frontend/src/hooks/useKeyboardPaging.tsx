import React, { useEffect, useState, DependencyList } from "react";
import { Pager } from "@components";
import { useHotkeys } from "react-hotkeys-hook";

export function useKeyboardPaging(nItems: number, deps: DependencyList = [], perPage: number = 20): Pager {
  const [curItem, setCurItem] = useState<number>(0);

  useHotkeys("left", (event) => {
    setCurItem((cur) => Math.max(cur - 1, 0));
    event.preventDefault();
  });
  useHotkeys("pageup", (event) => {
    setCurItem((cur) => Math.max(cur - perPage * 10, 0));
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
  useHotkeys("pagedown", (event) => {
    var max = Math.max(nItems - perPage * 10, 0);
    setCurItem((cur) => Math.min(max, cur + perPage * 10));
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

  const setPage = (newPage: number) => {
    setCurItem((newPage - 1) * perPage);
  };

  const pageNumber = curItem < perPage ? 1 : Math.ceil(curItem / perPage) + 1;
  const totalPages = Math.ceil(nItems / perPage);
  return {
    curItem: curItem,
    perPage: perPage,
    count: nItems,
    pageNumber: pageNumber,
    totalPages: totalPages,
    setpage: setPage,
  };
}
