import { useEffect, useState } from "react";
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

  const handleChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const newChain = event.target.value;
    changeChain(newChain);
  };

  return (
    <div>
      <label htmlFor="chain-selector">Select Chain:</label>
      <select id="chain-selector" value={chain} onChange={handleChange}>
        {chainList.map((chain) => (
          <option key={chain} value={chain}>
            {chain}
          </option>
        ))}
      </select>
    </div>
  );
};
