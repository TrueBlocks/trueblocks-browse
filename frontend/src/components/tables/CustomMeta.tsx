import React, { ReactNode } from "react";
import { ColumnDef } from "@tanstack/react-table";

export interface CustomMeta {
  className?: string;
  editor?: (value: () => any) => ReactNode;
}

export type CustomColumnDef<TData, TValue> = ColumnDef<TData, TValue> & {
  meta?: CustomMeta;
};
