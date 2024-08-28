import React from "react";
import { Pagination } from "@mantine/core";
import classes from "./Paginator.module.css";
import { Route } from "@/Routes";

export type Pager = {
  name: Route;
  nItems: number;
  selected: number;
  pageNumber: number;
  perPage: number;
  lastPage: number;
  setRecord: (newRecord: number) => void;
  setPage: (newPage: number) => void;
  getOffset: () => number;
};

export const Paginator = ({ pager }: { pager: Pager | null }) => {
  if (!pager) {
    return null;
  }

  return (
    <div className={classes.paginator}>
      <Pagination
        // siblings={1}
        size="sm"
        value={pager.pageNumber}
        total={pager.lastPage}
        // withEdges
        classNames={{
          root: classes.root,
          control: classes.control,
        }}
        onChange={(value) => {
          pager.setRecord((value - 1) * pager.perPage);
        }}
      />
    </div>
  );
};

export const EmptyPager: Pager = {
  name: "",
  nItems: 0,
  selected: 0,
  perPage: 0,
  pageNumber: 0,
  lastPage: 0,
  setPage: (newPage: number) => {},
  setRecord: (newRecord: number) => {},
  getOffset: () => 0,
};
