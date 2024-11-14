import { useState, useRef, useEffect } from "react";
import { TextInput, Group } from "@mantine/core";
import { useHotkeys } from "react-hotkeys-hook";
import { useViewState } from "@state";

export const SearchBar = () => {
  const { filter, updateFilter } = useViewState();
  const [localFilter, setLocalFilter] = useState(filter);
  const inputRef = useRef<HTMLInputElement>(null);

  const updateLocalFilter = (value: string) => {
    setLocalFilter(value);
  };

  const handleSearch = () => {
    const trimmedFilter = localFilter.trim();
    updateFilter(trimmedFilter);
    setLocalFilter(trimmedFilter);
    inputRef.current?.blur();
  };

  useEffect(() => {
    setLocalFilter(filter);
  }, [filter]);

  useHotkeys("mod+shift+f", (e: KeyboardEvent) => {
    e.preventDefault();
    inputRef.current?.focus();
  });

  return (
    <Group style={{ justifyContent: "flex-end", gap: "8px" }}>
      <TextInput
        ref={inputRef}
        value={localFilter}
        onChange={(e) => updateLocalFilter(e.currentTarget.value)}
        placeholder="Search..."
        onKeyDown={(e) => {
          if (e.key === "Enter") {
            handleSearch();
          }
        }}
        style={{ flexGrow: 1 }}
        size="xs"
      />
    </Group>
  );
};
