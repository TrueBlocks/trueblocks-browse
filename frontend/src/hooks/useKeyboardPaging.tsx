import { useState, DependencyList, useCallback, useMemo, useEffect } from "react";
import { Pager, EmptyPager } from "@components";
import { useHotkeys } from "react-hotkeys-hook";
import { Route } from "@/Routes";
import { CancleContexts, Reload } from "@gocode/app/App";
import { useAppState } from "@state";
import { base } from "@gocode/models";

export type Page = {
  selected: number;
  getOffset: () => number;
};

export function useKeyboardPaging(
  route: Route,
  nItems: number,
  perPage: number = 20,
  onEnter: (page: Page) => void
): Pager {
  const { address } = useAppState();
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

  const setPage = useCallback((newPage: number) => {
    setSelected((newPage - 1) * perPage);
    setPageNumber(newPage);
  }, [setSelected, setPageNumber]);

  const getOffset = () => (pageNumber - 1) * perPage;

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
  useHotkeys("mod+r", (e) => {
    e.preventDefault();
    Reload(address).then(() => {});
  });

  useHotkeys(
    "enter",
    (e) => {
      e.preventDefault();
      onEnter({
        selected,
        getOffset,
      });
    },
    [onEnter, selected, pageNumber, perPage]
  );

  const memo = useMemo(() => ({
    name: route,
    selected,
    perPage,
    nItems,
    pageNumber,
    lastPage,
    setRecord,
    setPage,
    getOffset,
  }), [selected, pageNumber, lastPage]);

  if (nItems < 0) {
    return EmptyPager;
  } else {
    return memo;
  }
}
