import { Editor, FormField, EditorProps, requiredString, validAddress } from "@components";
import { editors } from "@gocode/models";
import { SaveName } from "@gocode/app/App";

export const NameEditor = ({ source, onCancel }: EditorProps<editors.Name>) => {
  const fields: Array<FormField<editors.Name>> = [
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

  const saveData = async (values: editors.Name) => {
    return SaveName(values);
  };

  const legend = "Edit Name";
  return <Editor source={source} fields={fields} saveData={saveData} legend={legend} onCancel={onCancel} />;
};
