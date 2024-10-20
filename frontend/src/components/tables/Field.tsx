import { ReactNode } from "react";
import { Flex, Text } from "@mantine/core";
import { CellType, Formatter } from "@components";
import classes from "./Field.module.css";

type BaseField = {
  label?: string;
};

type Field<T> = BaseField & {
  type: CellType;
  accessor: keyof T;
};

export type FieldGroup<T> = {
  label: string;
  colSpan?: number;
  fields?: Field<T>[];
  components?: ReactNode[];
  buttons?: ReactNode[];
  collapsable?: boolean;
};

export const isCollapsable = <T,>(group: FieldGroup<T>): boolean => {
  return group.collapsable ?? true;
};

export const isButton = <T,>(group: FieldGroup<T>): boolean => {
  return Boolean(group.buttons);
};

type FieldRendererProps<T> = {
  field: Field<T>;
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
