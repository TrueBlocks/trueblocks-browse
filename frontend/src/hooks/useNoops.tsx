import { app } from "@gocode/models";
import { Page } from "./useKeyboardPaging";

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const fetchNoop = (currentItem: number, itemsPerPage: number) => {};

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const modifyNoop = async (arg1: app.ModifyData) => {};

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const enterNoop = (_unused: Page) => {};

export const useNoops = () => {
  return { enterNoop, fetchNoop, modifyNoop };
};
