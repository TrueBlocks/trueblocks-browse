import { useEffect, useRef, useState } from "react";
import { Button, TextInput, Grid, Group, Loader } from "@mantine/core";
import { useForm } from "@mantine/form";
import { showNotification } from "@mantine/notifications";
import { notifySuccess, notifyError, FieldsetWrapper } from "@components";

export interface FormField<T> {
  name: keyof T;
  label: string;
  placeholder?: string;
  validate?: (value: any) => string | null;
  inputProps?: Record<string, any>;
}

export interface EditorProps<T> {
  initialValues: T;
  fields: FormField<T>[];
  validateSchema: (values: T) => Record<string, string | null>;
  loadData: () => Promise<T>;
  saveData: (values: T) => Promise<void>;
  legend: string;
}

export const Editor = <T extends object>({
  initialValues,
  fields,
  validateSchema,
  loadData,
  saveData,
  legend,
}: EditorProps<T>) => {
  const [loading, setLoading] = useState(false);
  const form = useForm<T>({
    initialValues,
    validate: validateSchema,
    validateInputOnChange: true,
  });

  const formRef = useRef(form);
  useEffect(() => {
    const handleSelectAll = (event: KeyboardEvent) => {
      const activeElement = document.activeElement;
      if ((event.metaKey || event.ctrlKey) && event.key === "a" && activeElement instanceof HTMLInputElement) {
        event.preventDefault();
        activeElement.select();
      }
    };

    const fetchData = async () => {
      try {
        const data = await loadData();
        formRef.current.setValues(data);
      } catch (error) {
        showNotification({
          title: "Error",
          message: `Failed to load data ${error}`,
          color: "red",
        });
      }
    };

    fetchData();
    window.addEventListener("keydown", handleSelectAll);

    return () => {
      window.removeEventListener("keydown", handleSelectAll);
    };
  }, [loadData]);

  const handleSubmit = (values: T) => {
    setLoading(true);
    saveData(values)
      .then(() => {
        notifySuccess("Data saved successfully!");
        setLoading(false);
      })
      .catch((error) => {
        notifyError((error as Error).message || "Failed to save data!");
        setLoading(false);
      });
  };

  return (
    <FieldsetWrapper legend={legend}>
      <form onSubmit={form.onSubmit(handleSubmit)}>
        <Grid>
          {fields.map((field) => (
            <Grid.Col span={6} key={String(field.name)}>
              <TextInput
                label={field.label}
                placeholder={field.placeholder || `Enter ${String(field.name)}`}
                {...form.getInputProps(field.name)}
                error={form.errors[field.name]}
                {...field.inputProps}
              />
            </Grid.Col>
          ))}
        </Grid>
        <Group justify="flex-end" mt="md">
          <Button type="submit" disabled={loading}>
            {loading ? <Loader size="sm" /> : "Submit"}
          </Button>
        </Group>
      </form>
    </FieldsetWrapper>
  );
};
