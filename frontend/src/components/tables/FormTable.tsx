import { ReactNode } from "react";
import { Fieldset, Container, Grid, Flex, Text, Stack } from "@mantine/core";
import { Formatter, knownType } from "@components";
import classes from "./FormTable.module.css";

type FieldDefinition<T> = {
  type: knownType;
  label: string;
  accessor: keyof T;
};

type CustomComponentDefinition = {
  component: ReactNode;
};

export type GroupDefinition<T> = {
  title: string;
  fields?: FieldDefinition<T>[]; // Optional to allow custom components without fields
  colSpan?: number;
  components?: CustomComponentDefinition[]; // New field for custom components
};

type FormTableProps<T> = {
  data: Partial<T>;
  definition: GroupDefinition<T>[];
};

export function FormTable<T>({ data, definition }: FormTableProps<T>) {
  return (
    <Container style={{ maxWidth: "95%", paddingRight: "5%" }}>
      <Grid>
        {definition.map((group, index) => (
          <Grid.Col span={group.colSpan || 12} key={index}>
            <Stack>
              <Fieldset legend={group.title} bg="white" className={classes.fieldSet}>
                {group.fields?.map((field, fieldIndex) => {
                  const value = <Formatter type={field.type} value={data[field.accessor]} />;
                  return (
                    <Flex key={fieldIndex} gap="md" align="center">
                      {field.label.length > 0 ? <Text className={classes.fieldPrompt}>{field.label}</Text> : <></>}
                      <Text>{value}</Text>
                    </Flex>
                  );
                })}
                {group.components?.map((customComponent, componentIndex) => (
                  <div key={componentIndex}>{customComponent.component}</div>
                ))}
              </Fieldset>
            </Stack>
          </Grid.Col>
        ))}
      </Grid>
    </Container>
  );
}
