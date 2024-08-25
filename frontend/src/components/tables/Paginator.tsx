import React from "react";
import { Pagination } from "@mantine/core";
import classes from "./Paginator.module.css";
import { Route } from "@/Routes";

export type Pager = {
  name: Route;
  count: number;
  curItem: number;
  pageNumber: number;
  perPage: number;
  totalPages: number;
  setPage: (newPage: number) => void;
  offset: () => number;
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
        total={pager.totalPages}
        // withEdges
        classNames={{
          root: classes.root,
          control: classes.control,
        }}
        onChange={(value) => {
          pager.setPage(value);
        }}
      />
    </div>
  );
};

export const EmptyPager: Pager = {
  name: "",
  count: 0,
  curItem: 0,
  perPage: 0,
  pageNumber: 0,
  totalPages: 0,
  setPage: (newPage: number) => {},
  offset: () => 0,
};
