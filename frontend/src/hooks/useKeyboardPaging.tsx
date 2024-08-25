import React, { useEffect, useState, DependencyList, KeyboardEvent } from "react";
import { useLocation } from "wouter";
import { Pager, EmptyPager } from "@components";
import { useHotkeys, HotkeyCallback } from "react-hotkeys-hook";
import { Route } from "@/Routes";

export function useKeyboardPaging(
  route: Route,
  nItems: number,
  deps: DependencyList = [],
  perPage: number = 20
): Pager {
  const [curItem, setCurItem] = useState<number>(0);
  const [location] = useLocation();

  const handleKey = (e: any, callback: React.SetStateAction<number>) => {
    if (!location.includes(route)) {
      return;
    }
    e.preventDefault();
    setCurItem(callback);
  };

  useHotkeys("left", (e) => {
    handleKey(e, (cur) => Math.max(cur - 1, 0));
  });
  useHotkeys("pageup", (e) => {
    handleKey(e, (cur) => Math.max(cur - perPage * 10, 0));
  });
  useHotkeys("up", (e) => {
    handleKey(e, (cur) => Math.max(cur - perPage, 0));
  });
  useHotkeys("home", (e) => {
    handleKey(e, Math.max(0, 0));
  });

  useHotkeys("right", (e) => {
    handleKey(e, (cur) => Math.min(Math.max(nItems - perPage, 0), cur + 1));
  });
  useHotkeys("pagedown", (e) => {
    handleKey(e, (cur) => Math.min(Math.max(nItems - perPage * 10, 0), cur + perPage * 10));
  });
  useHotkeys("down", (e) => {
    handleKey(e, (cur) => Math.min(Math.max(nItems - perPage, 0), cur + perPage));
  });
  useHotkeys("end", (e) => {
    handleKey(e, Math.max(nItems - perPage, 0));
  });

  useEffect(() => {
    setCurItem(0);
  }, deps);

  const setPage = (newPage: number) => {
    setCurItem((newPage - 1) * perPage);
  };

  const pageNumber = curItem < perPage ? 1 : Math.ceil(curItem / perPage) + 1;
  const totalPages = Math.ceil(nItems / perPage);
  if (nItems < 0) {
    return EmptyPager;
  } else {
    return {
      name: route,
      curItem: curItem,
      perPage: perPage,
      count: nItems,
      pageNumber: pageNumber,
      totalPages: totalPages,
      setPage: setPage,
      offset: () => (pageNumber - 1) * perPage,
    };
  }
}
