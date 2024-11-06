import { useEffect, useState, useCallback } from "react";
import { Select } from "@mantine/core";
import { GetChains, Logger, SetChain } from "@gocode/app/App";
import { useAppState } from "../../state";

export const ChainSelector = () => {
  const { config } = useAppState();
  const [chain, setChain] = useState<string>("mainnet");
  const [chainList, setChainList] = useState<string[]>([]);

  const selectChain = useCallback((newChain: string) => {
    Logger([`selectChain: ${newChain}`]).then(() => {});
    setChain(newChain);
    SetChain(newChain).then(() => {});
  }, []);

  useEffect(() => {
    GetChains().then((chains) => {
      setChainList(chains.length > 0 ? chains : ["mainnet"]);
    });
  }, [config]);

  useEffect(() => {
    if (!chainList.includes(chain)) {
      Logger([`Selected chain "${chain}" not found. Switching to "${chainList[0]}"`]).then(() => {});
      selectChain(chainList[0]);
    }
  }, [chain, chainList, selectChain]);

  const handleChange = (value: string | null) => {
    if (value) {
      Logger([`handleChange-Selected chain: ${value}`]).then(() => {});
      selectChain(value);
    }
  };

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
