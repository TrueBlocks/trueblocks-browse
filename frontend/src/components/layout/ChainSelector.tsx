import { useEffect, useState, useCallback } from "react";
import { Select } from "@mantine/core";
import { GetChains, SetChain } from "@gocode/app/App";
import { useAppState } from "@state";

export const ChainSelector = () => {
  const { info, config } = useAppState();
  const [selected, setSelected] = useState<string>(info.chain);
  const [chainList, setChainList] = useState<string[]>(["mainnet"]);

  const selectChain = useCallback((newChain: string) => {
    setSelected(newChain);
    SetChain(newChain).then(() => {});
  }, []);

  const handleChange = (value: string | null) => {
    if (value) {
      selectChain(value);
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
      <Select
        id="chain-selector"
        value={selected}
        onChange={handleChange}
        data={chainList.map((ch) => ({ value: ch, label: ch }))}
        placeholder="Choose a chain"
      />
    </div>
  );
};
