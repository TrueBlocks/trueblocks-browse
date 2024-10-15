import { useEffect, useState } from "react";
import { editors } from "@gocode/models";
import { LoadName, SaveName } from "../../../wailsjs/go/app/App";
import { Editor, FormField } from "./Editor";

export const NameEditor = () => {
  const [initialValues, setInitialValues] = useState<editors.Name | null>(null);

  const fields: Array<FormField<editors.Name>> = [
    { name: "name", label: "Name", validate: (value: string) => (value.length === 0 ? "Name is required" : null) },
    {
      name: "address",
      label: "Address",
      validate: (value: string) =>
        /^0x[a-fA-F0-9]{40}$/.test(value) || /^[a-zA-Z0-9.-]+\.eth$/.test(value)
          ? null
          : "Invalid Ethereum address or ENS name",
    },
    { name: "tags", label: "Tags" },
    { name: "source", label: "Source" },
  ];

  useEffect(() => {
    const fetchInitialValues = async () => {
      const initial = editors.Name.createFrom({
        name: "",
        address: "0x0",
        tags: "",
        source: "",
        decimals: 18,
        symbol: "",
      });
      setInitialValues(initial);
    };

    fetchInitialValues();
  }, []);

  const validateSchema = (values: editors.Name) =>
    fields.reduce(
      (errors, field) => {
        const error = field.validate?.(values[field.name]);
        if (error) errors[field.name] = error;
        return errors;
      },
      {} as Record<string, string | null>
    );

  const loadData = async () => {
    const data = await LoadName("0x12");
    return editors.Name.createFrom({
      ...data,
    });
  };

  const saveData = async (values: editors.Name) => {
    const convertedData = {
      ...values,
    };
    return SaveName(convertedData);
  };

  // Render nothing until initialValues are ready
  if (!initialValues) {
    return <div>Loading...</div>; // You can replace this with a loader component
  }

  return (
    <Editor
      initialValues={initialValues}
      fields={fields}
      validateSchema={validateSchema}
      loadData={loadData}
      saveData={saveData}
      legend="Edit Name"
    />
  );
};
