import React from "react";

import { Pagination } from "@mantine/core";
import classes from "./Paginator.module.css";

export type Pager = {
  curItem: number;
  perPage: number;
  count: number;
};

export const Paginator = ({ pager }: { pager: Pager }) => {
  const pageNumber = pager.curItem < pager.perPage ? 1 : Math.ceil(pager.curItem / pager.perPage) + 1;
  const totalPages = Math.ceil(pager.count / pager.perPage);

  return (
    <div className={classes.paginator}>
      <Pagination
        classNames={{ control: classes.red }}
        // siblings={1}
        size="sm"
        value={pageNumber}
        total={totalPages}
        // withEdges
        styles={{ control: { border: "1px solid grey" } }}
      />
    </div>
  );
};
