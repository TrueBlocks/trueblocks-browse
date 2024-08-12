import React from "react";

import { Pagination } from "@mantine/core";
import classes from "./Paginator.module.css";

export type Pager = {
  curItem: number;
  perPage: number;
  count: number;
  pageNumber: number;
  totalPages: number;
};

export const Paginator = ({ pager }: { pager: Pager }) => {
  return (
    <div className={classes.paginator}>
      <Pagination
        classNames={{ control: classes.red }}
        // siblings={1}
        size="sm"
        value={pager.pageNumber}
        total={pager.totalPages}
        // withEdges
        styles={{ control: { border: "1px solid grey" } }}
      />
    </div>
  );
};
