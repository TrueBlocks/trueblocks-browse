import { useEffect, useState } from "react";
import { Button, TextInput, Grid, Group, Loader } from "@mantine/core";
import { useForm } from "@mantine/form";
import { notifySuccess, notifyError, FieldsetWrapper } from "@components";

export interface FormField<T> {
  name: keyof T;
  label: string;
  placeholder?: string;
  validate?: (value: any) => string | null;
  inputProps?: Record<string, any>;
}

export interface EditorProps<T> {
  source: T;
  onCancel: () => void;
}

export interface editProps<T> extends EditorProps<T> {
  fields: FormField<T>[];
  saveData: (values: T) => Promise<void>;
  legend: string;
}

export const Editor = <T extends object>({ source, fields, saveData, onCancel, legend }: editProps<T>) => {
  const [saving, setSaving] = useState(false);
  const form = useForm<T>({
    initialValues: source,
    validate: fields
      .filter((field) => field.validate)
      .reduce(
        (acc, field) => ({
          ...acc,
          [field.name]: field.validate!,
        }),
        {}
      ),
    validateInputOnChange: true,
  });

  useEffect(() => {
    const handleSelectAll = (event: KeyboardEvent) => {
      const activeElement = document.activeElement;
      if ((event.metaKey || event.ctrlKey) && event.key === "a" && activeElement instanceof HTMLInputElement) {
        event.preventDefault();
        activeElement.select();
      }
    };

    window.addEventListener("keydown", handleSelectAll);
    return () => {
      window.removeEventListener("keydown", handleSelectAll);
    };
  }, []);

  const handleSubmit = (values: T) => {
    setSaving(true);
    try {
      saveData(values).then(() => {
        notifySuccess("Data saved successfully!");
      });
    } catch (error) {
      notifyError((error as Error).message || "Failed to save data!");
    } finally {
      setSaving(false);
    }
  };

  return (
    <FieldsetWrapper legend={legend}>
      <form onSubmit={form.onSubmit(handleSubmit)}>
        <Grid>
          {fields.map((field) => {
            return (
              <Grid.Col span={6} key={String(field.name)}>
                <TextInput
                  label={field.label}
                  placeholder={field.placeholder || `Enter ${String(field.name)}`}
                  {...form.getInputProps(field.name)}
                  {...field.inputProps}
                />
              </Grid.Col>
            );
          })}
        </Grid>
        <Group justify="flex-end" mt="md">
          <Button variant="default" onClick={onCancel} disabled={saving}>
            Cancel
          </Button>
          <Button type="submit" disabled={saving}>
            {saving ? <Loader size="sm" /> : "Submit"}
          </Button>
        </Group>
      </form>
    </FieldsetWrapper>
  );
};
