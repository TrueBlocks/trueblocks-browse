import React, { ReactNode } from "react";
import { Container, Grid, Flex, Text, Title, Divider, Stack } from "@mantine/core";

type FieldDefinition<T> = {
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
  customComponents?: CustomComponentDefinition[]; // New field for custom components
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
              <Title order={4}>{group.title}</Title>
              <Divider />
              {group.fields?.map((field, fieldIndex) => (
                <Flex key={fieldIndex} gap="md" align="center">
                  <Text style={{ backgroundColor: "lightgrey", minWidth: "150px" }}>{field.label}</Text>{" "}
                  <Text>{String(data[field.accessor])}</Text>
                </Flex>
              ))}
              {group.customComponents?.map((customComponent, componentIndex) => (
                <div key={componentIndex}>{customComponent.component}</div>
              ))}
            </Stack>
          </Grid.Col>
        ))}
      </Grid>
      <Divider my="lg" />
    </Container>
  );
}
