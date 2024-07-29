import React, { useState, useEffect } from "react";
import { Select } from "@mantine/core";
import { GetExistingAddrs } from "@gocode/app/App";

type SelectItem = {
  value: string;
  label: string;
};

interface EditableSelectProps {
  value: string;
  onChange: (value: string) => void;
  label: string;
  placeholder: string;
}

export const EditableSelect: React.FC<EditableSelectProps> = ({ value, onChange, label, placeholder }) => {
  const [options, setOptions] = useState<SelectItem[]>([]);
  const [inputValue, setInputValue] = useState<string>(value);

  useEffect(() => {
    async function fetchAddresses() {
      GetExistingAddrs().then((existingAddrs) => {
        const formattedOptions = existingAddrs.map((addr) => ({ value: addr, label: addr }));
        setOptions(formattedOptions);
        if (!existingAddrs.includes(value)) {
          setInputValue(existingAddrs.length > 0 ? existingAddrs[0] : "");
        } else {
          setInputValue(value); // Ensure the inputValue is set to the prop value if it exists in the options
        }
      });
    }
    fetchAddresses();
  }, [value]);

  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === "Enter" && inputValue.trim() !== "") {
      const newOption: SelectItem = { value: inputValue, label: inputValue };
      setOptions((current) => {
        const existingIndex = current.findIndex((option) => option.value.toLowerCase() === inputValue.toLowerCase());
        if (existingIndex === -1) {
          return [...current, newOption];
        } else {
          return current.map((option, index) => (index === existingIndex ? newOption : option));
        }
      });
      onChange(inputValue);
      event.preventDefault();
      setTimeout(() => {
        const selectElement = document.querySelector(".mantine-Select-input") as HTMLElement;
        if (selectElement) {
          selectElement.blur();
        }
      }, 0);
    }
  };

  const handleChange = (value: string | null) => {
    if (value !== null) {
      setInputValue(value);
      onChange(value);
    }
  };

  return (
    <Select
      label={label}
      placeholder={placeholder}
      data={options}
      searchable
      value={inputValue}
      onSearchChange={setInputValue}
      onChange={handleChange}
      onKeyDown={handleKeyDown}
      style={{ width: "600px" }}
    />
  );
};
