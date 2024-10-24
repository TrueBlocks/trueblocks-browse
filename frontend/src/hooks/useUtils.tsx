const ShortenAddr = (val: string) => (val.length > 14 ? `${val.slice(0, 8)}...${val.slice(-6)}` : val);

export const useUtils = () => ({ ShortenAddr });
