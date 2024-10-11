import { app, base } from "../../wailsjs/go/models";
import { useViewState } from "../state";

export const useModifyFn = (address: base.Address) => {
  const { fetchFn, modifyFn, pager } = useViewState();

  const deleteItem = () =>
    modifyFn(app.ModifyData.createFrom({ operation: "delete", address })).then(() =>
      fetchFn(pager.getOffset(), pager.perPage)
    );

  const undeleteItem = () =>
    modifyFn(app.ModifyData.createFrom({ operation: "undelete", address })).then(() =>
      fetchFn(pager.getOffset(), pager.perPage)
    );

  const removeItem = () =>
    modifyFn(app.ModifyData.createFrom({ operation: "remove", address })).then(() =>
      fetchFn(pager.getOffset(), pager.perPage)
    );

  return { deleteItem, undeleteItem, removeItem };
};
