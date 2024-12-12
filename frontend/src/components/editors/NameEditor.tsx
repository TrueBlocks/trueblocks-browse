import { Editor, EditorProps, FormField, requiredString, validAddress } from "@components";
import { SaveName } from "@gocode/app/App";
import { types } from "@gocode/models";

export const NameEditor = ({ source, onCancel }: EditorProps<types.Name>) => {
  const fields: Array<FormField<types.Name>> = [
    {
      name: "address",
      label: "Address",
      validate: validAddress,
    },
    {
      name: "name",
      label: "Name",
      validate: requiredString,
    },
    { name: "tags", label: "Tags" },
    { name: "source", label: "Source" },
  ];

  const saveData = (values: types.Name) => {
    return SaveName(values).then(() => {});
  };

  const legend = "Edit Name";
  return <Editor source={source} fields={fields} saveData={saveData} legend={legend} onCancel={onCancel} />;
};
