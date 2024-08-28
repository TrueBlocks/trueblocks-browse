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
  const [lastPage, setLastPage] = useState<number>(1);
  const [selected, setSelected] = useState<number>(0);

  useEffect(() => {
    setLastPage(Math.ceil(nItems / perPage));
  }, [nItems, perPage]);

  const setRecord = (newRecord: number) => {
    const curPage = Math.floor(selected / perPage);
    const newPage = Math.floor(newRecord / perPage);
    if (newPage !== curPage) {
      setPage(newPage + 1);
    }
    setSelected(newRecord);
  };

  const setPage = (newPage: number) => {
    setSelected((newPage - 1) * perPage);
    setPageNumber(newPage);
  };

  // keyboard shortcuts
  useHotkeys("up", (e) => {
    e.preventDefault();
    setRecord(Math.max(selected - 1, 0));
  });
  useHotkeys("left", (e) => {
    e.preventDefault();
    setRecord(Math.max(selected - perPage, 0));
  });
  useHotkeys("pageup", (e) => {
    e.preventDefault();
    setRecord(Math.max(selected - perPage * 10, 0));
  });
  useHotkeys("home", (e) => {
    e.preventDefault();
    setRecord(0);
  });

  useHotkeys("down", (e) => {
    e.preventDefault();
    setRecord(Math.min(selected + 1, nItems - 1));
  });
  useHotkeys("right", (e) => {
    e.preventDefault();
    setRecord(Math.min(selected + perPage, nItems - 1));
  });
  useHotkeys("pagedown", (e) => {
    e.preventDefault();
    setRecord(Math.min(selected + perPage * 10, nItems - 1));
  });
  useHotkeys("end", (e) => {
    e.preventDefault();
    setRecord(nItems - 1);
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
      setRecord,
      setPage,
      getOffset: () => (pageNumber - 1) * perPage,
    };
  }
}
