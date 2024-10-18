import { ReactNode } from "react";
import { Flex, Text } from "@mantine/core";
import { CellType, Formatter } from "@components";
import classes from "./Field.module.css";

type field<T> = {
  label?: string;
  type: CellType;
  accessor: keyof T;
};

type component = {
  label?: string;
  component: ReactNode;
};

export type FieldGroup<T> = {
  legend: string;
  colSpan?: number;
  fields?: field<T>[];
  components?: component[];
  collapsable?: boolean;
};

type FieldRendererProps<T> = {
  field: field<T>;
  data: Partial<T>;
};

export const FieldRenderer = <T,>({ field, data }: FieldRendererProps<T>) => {
  const label = field.label ? <Text className={classes.fieldPrompt}>{field.label}</Text> : <></>;
  const value = <Formatter type={field.type} value={data?.[field.accessor]} />;

  return (
    <Flex gap="md" align="center">
      {label}
      {value}
    </Flex>
  );
};
