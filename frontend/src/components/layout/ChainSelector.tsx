import { useEffect, useState, useRef, useCallback } from "react";
import { Select } from "@mantine/core";
import { GetChains, SetChain } from "@gocode/app/App";
import { useAppState } from "@state";

export const ChainSelector = () => {
  const { info, config } = useAppState();
  const [selected, setSelected] = useState<string>(info.chain);
  const [chainList, setChainList] = useState<string[]>(["mainnet"]);
  const focusRef = useRef<HTMLDivElement>(null);
  const selectRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    focusRef.current?.focus();
  }, []);

  const selectChain = useCallback((newChain: string) => {
    setSelected(newChain);
    SetChain(newChain).then(() => {});
  }, []);

  const handleChange = (value: string | null) => {
    if (value) {
      selectChain(value);
      selectRef.current?.blur();
    }
  };

  useEffect(() => {
    GetChains().then((chains) => {
      setChainList(chains.length > 0 ? chains : ["mainnet"]);
    });
  }, [config]);

  useEffect(() => {
    if (!chainList.includes(selected)) {
      selectChain("mainnet");
    }
  }, [selected, chainList, selectChain]);

  return (
    <div>
      <div ref={focusRef} tabIndex={-1} />
      <Select
        id="chain-selector"
        value={selected}
        onChange={handleChange}
        data={chainList.map((ch) => ({ value: ch, label: ch }))}
        placeholder="Choose a chain"
        ref={selectRef}
        autoFocus={false}
      />
    </div>
  );
};
