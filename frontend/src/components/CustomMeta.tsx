import React, { ReactNode } from "react";

export interface CustomMeta {
  className?: string;
  editor?: (value: () => any) => ReactNode;
}
