import { useEffect, useState } from "react";
import { Select } from "@mantine/core";
import { GetChainList } from "@gocode/app/App";
import { useAppState } from "@state";

export const ChainSelector = () => {
  const { chain, changeChain } = useAppState();
  const [chainList, setChainList] = useState<string[]>([]);

  useEffect(() => {
    GetChainList().then((chains) => {
      setChainList(chains);
    });
  }, []);

  useEffect(() => {
    GetChainList().then((chains) => {
      setChainList(chains);
      if (!chains.includes(chain)) {
        changeChain(chains[0]);
      }
    });
  }, [chain, changeChain]);

  const handleChange = (value: string | null) => {
    if (value) {
      changeChain(value);
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
