import { app } from "../../wailsjs/go/models";

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const fetchNoop = async (currentItem: number, itemsPerPage: number) => {};

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const modifyNoop = async (arg1: app.ModifyData) => {};

export const useNoops = () => {
  return { fetchNoop, modifyNoop };
};
