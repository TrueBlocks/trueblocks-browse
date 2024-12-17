import { useState, useRef, useEffect } from "react";
import { TextInput, Group } from "@mantine/core";
import { useHotkeys } from "react-hotkeys-hook";
import { GetFilter, SetFilter } from "@gocode/app/App";
import { useViewState } from "@state";

export const SearchBar = () => {
  const { fetchFn, pager } = useViewState();
  const [filter, setFilter] = useState<string>("");
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    GetFilter().then((filterData) => {
      setFilter(filterData.criteria);
    });
  }, []);

  const handleSearch = () => {
    const criteria = filter.trim();
    setFilter(criteria);
    SetFilter(criteria).then(() => {
      fetchFn(pager.getOffset(), pager.perPage);
    });
    inputRef.current?.blur();
  };

  useHotkeys("mod+shift+f", (e: KeyboardEvent) => {
    e.preventDefault();
    inputRef.current?.focus();
  });

  return (
    <Group style={{ justifyContent: "flex-end", gap: "8px" }}>
      <TextInput
        ref={inputRef}
        value={filter}
        onChange={(e) => setFilter(e.currentTarget.value)}
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
