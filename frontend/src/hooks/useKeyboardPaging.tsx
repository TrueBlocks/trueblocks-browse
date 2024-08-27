import React, { useEffect, useState, DependencyList } from "react";
import { Pager, EmptyPager } from "@components";
import { useHotkeys } from "react-hotkeys-hook";
import { Route } from "@/Routes";
import { CancleContexts } from "@gocode/app/App";

export type Page = {
  selected: number;
  perPage: number;
  pageNumber: number;
};

export function useKeyboardPaging(
  route: Route,
  nItems: number,
  deps: DependencyList = [],
  perPage: number = 20,
  onEnter: (page: Page) => void
): Pager {
  const [pageNumber, setPageNumber] = useState<number>(1);
  const [selected, setSelected] = useState<number>(0);

  const lastPage = Math.floor(nItems / perPage);

  // Navigate one row at a time
  useHotkeys("up", (e) => {
    e.preventDefault();
    // console.log("Up Pressed. curItem:", selected, " perPage:", perPage);
    const curPage = Math.floor(selected / perPage);
    const newRec = Math.max(selected - 1, 0);
    const newPage = Math.floor(newRec / perPage);
    if (newPage !== curPage) {
      setPage(newPage + 1); // page is one based
    }
    setSelected(newRec);
  });
  useHotkeys("down", (e) => {
    e.preventDefault();
    // console.log("Down Pressed. curItem:", selected, " perPage:", perPage);
    const curPage = Math.floor(selected / perPage);
    const newRec = Math.min(selected + 1, nItems - 1);
    const newPage = Math.floor(newRec / perPage);
    if (newPage !== curPage) {
      setPage(newPage + 1); // page is one based
    }
    setSelected(newRec);
  });

  // Navigate one page at a time
  useHotkeys("left", (e) => {
    e.preventDefault();
    // console.log("Left Pressed. curItem:", selected, " perPage:", perPage);
    const curPage = Math.floor(selected / perPage);
    const newRec = Math.max(selected - perPage, 0);
    const newPage = Math.floor(newRec / perPage);
    if (newPage !== curPage) {
      setPage(newPage + 1); // page is one based
    }
    setSelected(newPage * perPage);
  });
  useHotkeys("right", (e) => {
    e.preventDefault();
    // console.log("Right Pressed. curItem:", selected, " perPage:", perPage);
    const curPage = Math.floor(selected / perPage);
    const newRec = Math.min(selected + perPage, nItems - 1);
    const newPage = Math.floor(newRec / perPage);
    if (newPage !== curPage) {
      setPage(newPage + 1); // page is one based
    }
    setSelected(newPage * perPage);
  });

  // Navigate ten pages at a time
  useHotkeys("pageup", (e) => {
    e.preventDefault();
    // console.log("PageUp Pressed. curItem:", selected, " perPage:", perPage);
    const curPage = Math.floor(selected / perPage);
    const newRec = Math.max(selected - perPage * 10, 0);
    const newPage = Math.floor(newRec / perPage);
    if (newPage !== curPage) {
      setPage(newPage + 1); // page is one based
    }
    setSelected(newPage * perPage);
  });
  useHotkeys("pagedown", (e) => {
    e.preventDefault();
    // console.log("PageDown Pressed. curItem:", selected, " perPage:", perPage);
    const curPage = Math.floor(selected / perPage);
    const newRec = Math.min(selected + perPage * 10, nItems - 1);
    const newPage = Math.floor(newRec / perPage);
    if (newPage !== curPage) {
      setPage(newPage + 1); // page is one based
    }
    setSelected(newPage * perPage);
  });

  // Navigate to first and last records
  useHotkeys("home", (e) => {
    e.preventDefault();
    // console.log("Home Pressed. curItem:", selected, " perPage:", perPage);
    setPage(1);
    setSelected(0);
  });

  useHotkeys("end", (e) => {
    e.preventDefault();
    // console.log("End Pressed. curItem:", selected, " perPage:", perPage);
    setPage(lastPage + 1);
    setSelected(lastPage * perPage);
  });

  useHotkeys("esc", (e) => {
    e.preventDefault();
    CancleContexts();
  });

  useHotkeys(
    "enter",
    (e) => {
      e.preventDefault();
      onEnter({
        selected,
        perPage,
        pageNumber,
      });
    },
    [onEnter, selected, pageNumber, perPage]
  );

  const setPage = (newPage: number) => {
    setSelected((newPage - 1) * perPage);
    setPageNumber(newPage);
  };

  if (nItems < 0) {
    return EmptyPager;
  } else {
    return {
      name: route,
      selected,
      perPage,
      nItems,
      pageNumber,
      lastPage,
      setPage,
      offset: () => (pageNumber - 1) * perPage,
    };
  }
}
