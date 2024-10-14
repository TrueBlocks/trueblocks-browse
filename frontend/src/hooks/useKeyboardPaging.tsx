import { useEffect, useState } from "react";
import { useHotkeys } from "react-hotkeys-hook";
import { Pager, EmptyPager } from "@components";
import { CancelAllContexts, Reload } from "@gocode/app/App";
import { Route } from "@layout";
import { useAppState } from "@state";

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

  const setPage = (newPage: number) => {
    setSelected((newPage - 1) * perPage);
    setPageNumber(newPage);
  };

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
    CancelAllContexts();
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
      getOffset,
    };
  }
}
