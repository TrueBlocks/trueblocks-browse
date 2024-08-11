import React from "react";

import { Pagination } from "@mantine/core";
import classes from "./Paginator.module.css";

export const Paginator = ({ pageNumber, totalPages }: { pageNumber: number; totalPages: number }) => {
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
