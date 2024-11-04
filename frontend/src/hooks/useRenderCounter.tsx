import { useState, useEffect } from "react";

export const useRenderCounter = () => {
  const [renderCount, setRenderCount] = useState(0);

  useEffect(() => {
    setRenderCount((count) => count + 1);
  }, []);

  return renderCount;
};
