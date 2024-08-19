import React from "react";
import { Pagination } from "@mantine/core";
import classes from "./Paginator.module.css";

export type Pager = {
  curItem: number;
  perPage: number;
  count: number;
  pageNumber: number;
  totalPages: number;
  setpage: (newPage: number) => void;
};

export const Paginator = ({ pager }: { pager: Pager }) => {
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
          pager.setpage(value);
        }}
      />
    </div>
  );
};
