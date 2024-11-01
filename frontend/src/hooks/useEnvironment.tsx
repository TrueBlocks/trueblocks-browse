import { useState, useEffect } from "react";
import { GetEnv } from "@gocode/app/App";

export const useEnvironment = (key: string): string => {
  const [envMap, setEnvMap] = useState<Map<string, string>>(new Map());

  useEffect(() => {
    const fetchEnvValue = () => {
      if (!envMap.has(key)) {
        GetEnv(key).then((val: string) => {
          setEnvMap((prevMap) => new Map(prevMap).set(key, val));
        });
      }
    };
    fetchEnvValue();
  }, [key, envMap]);

  return envMap.get(key) || "";
};
