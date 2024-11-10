import { useEffect, useState, useCallback } from "react";
import { Select } from "@mantine/core";
import { GetChains, SetChain } from "@gocode/app/App";
import { useAppState } from "@state";

export const ChainSelector = () => {
  const { config } = useAppState();
  const [chain, setChain] = useState<string>("mainnet");
  const [chainList, setChainList] = useState<string[]>(["mainnet"]);

  const selectChain = useCallback((newChain: string) => {
    setChain(newChain);
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
    if (!chainList.includes(chain)) {
      selectChain("mainnet");
    }
  }, [chain, chainList, selectChain]);

  return (
    <div>
      <Select
        id="chain-selector"
        value={chain}
        onChange={handleChange}
        data={chainList.map((chain) => ({ value: chain, label: chain }))}
        placeholder="Choose a chain"
      />
    </div>
  );
};
